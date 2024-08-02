package noc

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type trieNode struct {
	children map[string]*trieNode
	methods  map[string]HandlerFunc
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[string]*trieNode),
		methods:  make(map[string]HandlerFunc),
	}
}

type trieRouter struct {
	root *trieNode
}

func NewTrieRouter() *trieRouter {
	return &trieRouter{
		root: newTrieNode(),
	}
}

func (r *trieRouter) addRoute(method string, path string, handler HandlerFunc) {
	node := r.root
	segments := splitPath(path)
	for _, segment := range segments {
		if segment == "*" {
			// 处理动态路径参数
			node.children[segment] = newTrieNode()
			node = node.children[segment]
			break
		} else {
			if child, exists := node.children[segment]; exists {
				node = child
			} else {
				newNode := newTrieNode()
				node.children[segment] = newNode
				node = newNode
			}
		}
	}
	node.methods[method] = handler
}

func splitPath(path string) []string {
	if path == "/" {
		return []string{""}
	}
	return strings.Split(strings.Trim(path, "/"), "/")
}

func (r *trieRouter) findRoute(method string, path string) (handler HandlerFunc, params map[string]string) {
	node := r.root
	segments := splitPath(path)
	params = make(map[string]string)

	for i, segment := range segments {
		if child, exists := node.children[segment]; exists {
			node = child
		} else if child, exists := node.children["*"]; exists {
			// 动态路径参数
			params[fmt.Sprintf("param%d", i)] = segment
			node = child
		} else {
			return nil, nil
		}
	}

	if handler, ok := node.methods[method]; ok {
		return handler, params
	}

	return nil, nil
}

func (r *trieRouter) Route(method string, path string, handler HandlerFunc, filters ...FilterFunc) {
	for i := len(filters) - 1; i >= 0; i-- {
		handler = filters[i](handler)
	}
	r.addRoute(method, path, handler)
}

func (r *trieRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, params := r.findRoute(req.Method, req.URL.Path)
	if handler == nil {
		http.NotFound(w, req)
		return
	}

	ctx := context.WithValue(req.Context(), "params", params)

	handler(NewContext(w, req.WithContext(ctx)))
}
