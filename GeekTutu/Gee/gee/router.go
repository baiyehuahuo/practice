package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node       // Method root, GET POST ...
	handlers map[string]HandlerFunc // Get handler by pattern
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 解析 pattern 若解析到 '*' 停止解析
func parsePattern(pattern string) []string {
	var parts []string
	vs := strings.Split(pattern, "/")
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

// 添加路由方法 parttern 与 handler 对应
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)

	parts := parsePattern(pattern)
	key := method + "-" + pattern
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// 获取路由
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path) // 解析各部分
	params := make(map[string]string) // 解析参数
	root, ok := r.roots[method]       // '根'
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)
	if n == nil {
		return nil, nil
	}
	parts := parsePattern(n.pattern)
	for idx, part := range parts {
		if part[0] == ':' { // :param 直接对应该搜索
			params[part[1:]] = searchParts[idx]
		}
		if part[0] == '*' && len(part) > 1 {
			// * 后续停止查询
			params[part[1:]] = strings.Join(searchParts[idx:], "/")
			break
		}
	}
	return n, params
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n == nil {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		return
	}
	c.Params = params
	key := c.Method + "-" + n.pattern
	r.handlers[key](c)
}
