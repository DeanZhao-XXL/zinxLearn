package znet

import (
	"fmt"
	"net"
	"time"
	"ziface"
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

//==========定义当前客户端连接的handle api ======================

//==========实现 ziface.IServer里的全部接口方法==================

//开启网络服务
func (s *Server) Start() {
	fmt.Printf("[START] Server listener at IP:%s , Port %d , is starting\n", s.IP, s.Port)

	//开启一个go去做服务端listener业务
	go func() {
		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))

		if err != nil {
			fmt.Println("resolve tcp addr err:", err)
			return
		}

		//2 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listen", s.IPVersion, "err", err)
			return
		}
		//已经监听成功
		fmt.Println("start zwdnx server", s.Name, "succ,now listening")

		//3 启动server网络连接业务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			//todo Server.Start() 设置服务器最大连接控制,如果超过最大连接,那就关闭连接.
			//todo 处理新连接请求的 业务 方法,此时应该有handler和conn是绑定的

			// 我们这里啊按时做一个最大512字节的回显服务
			go func() {
				//不断地循环从客户端获取数据
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}
					//回显
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[stop] Zinx server , name", s.Name)

	//todo Server.Stop()将其他需要清理的连接信息或者其他信息,也要一并停止或者清理
}

func (s *Server) Serve() {
	s.Start()

	//todo Server.Serve() 是否在启动服务的时候,还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主GO退出, listener的go将会退出
	for {
		time.Sleep(10 * time.Second)
	}
}

//创建一个服务器句柄
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}
	return s
}
