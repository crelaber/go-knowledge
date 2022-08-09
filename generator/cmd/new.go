package cmd

import (
	"github.com/spf13/cobra"
	"go-knowledge/generator/internal/gen"
	"go-knowledge/generator/pkg/arg"
	"log"
)

var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new project from database",
	Run:   newProject,
}

func init() {
	newCmd.PersistentFlags().StringVarP(&arg.Database, "db", "d", "shop", "指定数据库名称")
	newCmd.PersistentFlags().StringVarP(&arg.MySQL, "mysql", "m", "root:root@mysql+tcp(127.0.0.1:3306)/information_schema", "指定存储")
	newCmd.PersistentFlags().StringVarP(&arg.Out, "out", "o", "./dist", "指定输出目录")
	newCmd.PersistentFlags().StringVarP(&arg.SshTunnel, "ssh", "s", "", "开启ssh隧道")
	newCmd.PersistentFlags().StringVarP(&arg.Table, "table", "t", "%%", "指定表名")
	newCmd.PersistentFlags().StringVarP(&arg.Module, "module", "M", "github.com/crelaber/shop", "指定module项目")
	newCmd.PersistentFlags().StringVarP(&arg.TmplDir, "tmpl-dir", "T", "tmpl", "指定渲染模板目录")
	RootCmd.AddCommand(newCmd)
}

func newProject(cmd *cobra.Command, args []string) {
	//根据数据库解析得到mysql的table schema
	tableSchema, err := gen.GetTableSchemas(arg.MySQL, arg.Database, arg.Table)
	if err != nil {
		log.Panic("[GetTableSchemas] getSchema fail", err.Error())
		return
	}

	schemaTps := gen.GetSchemaTpls(tableSchema)
	gen.Render(schemaTps)
}
