package server

type Options struct{
	TcpPort int
}

func (s *Options) Init(){
	s.TcpPort = 443
}