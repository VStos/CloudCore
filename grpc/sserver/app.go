package sserver

type OOptions struct {
	TcpPort int
}

func (s *OOptions) Init() {
	s.TcpPort = 443
}
