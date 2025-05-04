package server

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() error {
	return nil
}

func (s *Server) Run() error {
	return nil
}

func (s *Server) Shutdown() []error {
	return nil
}
