package weed

import "sync"

type Server struct {
	RouterGroup
	pool sync.Pool
}

func New() *Server {
	srv := &Server{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
	}
	srv.pool.New = func() interface{} {
		return 0
	}
	return srv
}

func (s *Server) Use(middleware HandlerFunc) {

}

func Default() *Server {
	srv := New()
	srv.Use(Logger())
	return srv
}

func (s *Server) Run() {

}
