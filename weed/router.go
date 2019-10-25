package weed

import "net/http"

type HandlerFunc http.HandlerFunc

type HandlersChain []HandlerFunc

type RouterGroup struct {
	Handlers HandlersChain
	basePath string
	root     bool
}
