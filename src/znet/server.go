package znet

import (
	"fmt"
	"net"
)

//IServer接口实现,定义一个Server服务类
type Server struct {
	//服务器名称
	Name string
	//tpc4 or other
	IPVersion string
	//服务绑定的ip地址
	IP string
	//服务绑定的端口
	Port int
}

//==========实现 ziface.IServer里的全部接口方法==================

//开启网络服务
func (s *Server) Start() {
	fmt.Printf("[START] Server listener at IP:%s , Port %d , is starting\n", s.IP, s.Port)

	//开启一个go去做服务端listener业务
	go func() {
		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s : %d", s.IP, s.Port))
	}()
}
