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

func handleFunc(c *net.TCPConn, buf []byte, cnt int) error {

	if _, err := c.Write(buf[:cnt]); err != nil {
		fmt.Println("write error:", err)
		return err
	}
	return nil
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

			conn := NewConnection(c, connID, handleFunc)
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

func NewServer(name string) face.IServer {
	s := &Server{
		Name:      name,
		IpVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      "8889",
	}

	return s
}
