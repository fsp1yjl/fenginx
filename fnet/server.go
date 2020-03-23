package fnet

import (
	face "fenginx/finterface"
	"fenginx/utils"
	"fmt"
	"net"
)

type Server struct {
	Name      string
	IpVersion string
	Ip        string
	Port      string
	Router    face.IRouter
}

func (s *Server) Start() {
	go func() {
		fmt.Println("STARTING ......")
		bindAddress := s.Ip + ":" + s.Port
		addr, err := net.ResolveTCPAddr(s.IpVersion, bindAddress)
		if err != nil {
			fmt.Println("resolve tcp addr error", err)
			return
		}

		listenner, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			fmt.Println(" tcp listener error", err)
			return
		}
		fmt.Println("server listen success")
		connID := 0
		for {

			c, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("accept error:", err)
				continue
			}

			conn := NewConnection(c, connID, s.Router)
			go conn.Start()

			connID++

		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	// do someting

	select {}
}

func (s *Server) AddRouter(r face.IRouter) {
	s.Router = r
}

func NewServer() face.IServer {
	s := &Server{
		Name:      utils.G.Name,
		IpVersion: "tcp4",
		Ip:        utils.G.Host,
		Port:      utils.G.Port,
	}

	return s
}
