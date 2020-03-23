package utils

import face "fenginx/finterface"

type ServerInfo struct {
	Name    string
	Host    string
	Port    string
	Version string
}
type Global struct {
	ServerInfo
	TCPServer face.IServer
}

func (g *Global) LoadConfig() {
	//todo 从配置文件加载配置
}

var G *Global

func init() {
	G = &Global{
		ServerInfo: ServerInfo{
			Name: "default_server",
			Host: "0.0.0.0",
			Port: "8889",
		},
	}

}
