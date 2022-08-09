package gen

import (
	"errors"
	"github.com/flosch/pongo2"
	"github.com/smartwalle/pongo2render"
	"go-knowledge/generator/internal/model"
	"go-knowledge/generator/pkg/arg"
	"go/format"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

var (
	//存放tpl文件，key是相对路径
	tpls map[string]string
	//存放tpl文件夹，key是相对路径
	tplDirs map[string]bool
)

func init() {
	pongo2.RegisterFilter("lowerfirst", lowerfirst)
}

// lowerfirst 首字母小写，注意不要和go关键字冲突
func lowerfirst(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	if in.Len() <= 0 {
		return pongo2.AsValue(""), nil
	}

	t := in.String()
	r, size := utf8.DecodeRuneInString(t)
	return pongo2.AsValue(strings.ToLower(string(r)) + t[size:]), nil
}

func loadTmpl() {
	tmplRepoDir := arg.TmplDir
	err := filepath.Walk(tmplRepoDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		if info.IsDir() && path != tmplRepoDir {
			relPath, _ := filepath.Rel(tmplRepoDir, path)
			tplDirs[relPath] = true
		}

		if err != nil {
			return err
		}

		b, e := ioutil.ReadFile(path)
		if e != nil {
			return nil
		}
		relPath, e := filepath.Rel(tmplRepoDir, path)
		if e != nil {
			return nil
		}

		tpls[relPath] = string(b)
		return nil
	})

	if err != nil {
		log.Println(err)
	}
}

// Render 渲染模板
func Render(schemalTpls map[string]model.Table) {
	loadTmpl()
	camelTableNames := make(map[string]struct{})
	for tableName := range schemalTpls {
		camelTableNames[snakeToCamel(tableName)] = struct{}{}
	}

	ctx := pongo2.Context{
		"camelTableName": camelTableNames,
	}
	render(ctx, schemalTpls)
}

//替换path中的特殊变量
func getPath(path string, tableName string) string {
	path = strings.ReplaceAll(path, "TABLE_NAME", tableName)
	path = strings.ReplaceAll(path, "go.tmpl", ".gen.go")
	return path
}

func loadImports() map[string]struct{} {
	imports := make(map[string]struct{})
	for relPath := range tplDirs {
		imports[arg.Module+"/app/"+relPath] = struct{}{}
	}
	imports["go.uber.org/zap"] = struct{}{}
	imports["github.com/jinzhu/gorm"] = struct{}{}
	return imports
}

func render(ctx pongo2.Context, schemas map[string]model.Table) {
	var render = pongo2render.NewRender(arg.TmplDir)
	for path, content := range tpls {
		var globalImports = loadImports()
		//删除自己所在的包，防止循环引用
		delete(globalImports, arg.Module+"/app/"+filepath.Dir(path))
		for tableName, schema := range schemas {
			schema.Imports = globalImports
			var hasOpenId, hasDeleteTime bool
			for _, value := range schema.Columns {
				if value.CamelName == "DeleteTime" {
					schema.Imports["time"] = struct{}{}
					hasDeleteTime = true
				}
			}
			schema.Imports["time"] = struct{}{}

			for _, value := range schema.Columns {
				if value.CamelName == "OpenId" && value.ColumnKey != "PRI" {
					hasOpenId = true
				}
			}
			ctx["tableName"] = tableName
			ctx["camelTableName"] = snakeToCamel(tableName)
			ctx["lcamelTableName"] = lowerFirst(snakeToCamel(tableName))
			ctx["hasOpenId"] = hasOpenId
			ctx["hasDeleteTime"] = hasDeleteTime
			ctx["imports"] = schema.Imports
			ctx["columns"] = schema.Columns
			ctx["hasPrimaryKey"] = schema.HasPrimaryKey
			ctx["camelPrimaryKey"] = schema.CamelPrimaryKey
			ctx["primaryKey"] = schema.PrimaryKey
			ctx["primaryKeyType"] = schema.PrimaryKeyType
			buf, err := render.TemplateFromString(content).Execute(ctx)
			if err = write(filepath.Join(arg.Out, getPath(path, tableName)), buf); err != nil {
				log.Panicln("[render] write err :", err.Error(), path, tableName, buf)
				return
			}
		}

	}
}

//写byte到文件
func write(filename string, buff string) (err error) {
	filePath := path.Dir(filename)
	err = createPath(filePath)
	if err != nil {
		err = errors.New("write create path " + err.Error())
		return
	}

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		err = errors.New("write create file " + err.Error())
		return
	}
	//格式化代码
	bts, err := format.Source([]byte(buff))
	if err != nil {
		err = errors.New("format buf error" + err.Error())
		return
	}

	err = ioutil.WriteFile(filename, bts, 0644)
	if err != nil {
		err = errors.New("write write file" + err.Error())
		return
	}
	return

}
