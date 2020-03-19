package fnet

import (
	face "fenginx/finterface"
	"fmt"
	"net"
)

type Server struct {
	Name      string
	IpVersion string
	Ip        string
	Port      string
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

		for {

			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("accept error:", err)
				continue
			}
			go func() {
				for {
					buf := make([]byte, 512)

					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("read error:", err)
						continue
					}

					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write error:", err)
						continue
					}
				}

			}()
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

func NewServer(name string) face.IServer {
	s := &Server{
		Name:      name,
		IpVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      "8889",
	}

	return s
}
