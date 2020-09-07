package main

import "znet"

//Server 模块的测试函数
func main() {
	//1 创建一个server 句柄s
	s := znet.NewServer("[zwd v0.1]")

	//2 开启服务
	s.Serve()

}
