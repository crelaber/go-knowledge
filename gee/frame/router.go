package frame

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	//log.Printf("Router %4s - %s", method, pattern)
	//key := method + "-" + pattern
	//r.handlers[key] = handler
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

//获取路由
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}

			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
			}
		}
		return n, params
	}
	return nil, nil

}

//所有的路由的处理方法
func (r *router) handle(c *Context) {
	//key := c.Method + "-" + c.Path
	//if handler, ok := r.handlers[key]; ok {
	//	handler(c)
	//} else {
	//	c.String(http.StatusNotFound, "404 NOT FOUND %s\n", c.Path)
	//}

	//router.go的变化比较小，比较重要的一点是，在调用匹配到的handler前，
	//将解析出来的路由参数赋值给了c.Params。这样就能够在handler中，通过Context对象访问到具体的值了。
	//添加中间件后，将从路由匹配得到的 Handler 添加到 c.handlers列表中，执行c.Next()。
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND %s\n", c.Path)
	}
	c.Next()
}
