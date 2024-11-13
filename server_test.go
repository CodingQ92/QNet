package QNet

import (
	"com.smyx/QNet/qnet"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	// 启动server
	qs := qnet.NewServer("QNet")
	qs.Serve()
	// 启动client
	go ClientTest()
	time.Sleep(time.Second * 12)
}

func ClientTest() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		_, _ = conn.Write([]byte(fmt.Sprintf("hello world >>> %d", i)))
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("[QNet] 客户端 [%s] 收到: %s\n", conn.LocalAddr(), string(buf[:n]))
		time.Sleep(time.Second * 1)
	}
}
