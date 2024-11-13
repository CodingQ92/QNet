package qnet

import (
	"com.smyx/QNet/iface"
	"fmt"
	"net"
)

type Connection struct {
	ConnID    uint32
	isClose   bool
	Conn      *net.TCPConn
	handleApi iface.HandleFunc
	ExistChan chan struct{}
}

func NewConnection(conn *net.TCPConn, connId uint32, handleApi iface.HandleFunc) *Connection {
	return &Connection{
		ConnID:    connId,
		isClose:   false,
		Conn:      conn,
		handleApi: handleApi,
		ExistChan: make(chan struct{}),
	}
}

func (c *Connection) Start() {
	fmt.Println("ConnId [", c.ConnID, "] Connection.Start()...")
	defer func() {
		c.Stop()
		fmt.Println("ConnId [", c.ConnID, "] Connection.Stop()...")
	}()
	for {
		buf := make([]byte, 512)
		n, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("Connection.Read() error:", err)
			continue
		}
		err = c.handleApi(c.Conn, buf[:n], n)
		if err != nil {
			fmt.Println("Connection.HandleApi() error:", err)
			break
		}
	}
}

func (c *Connection) Stop() {
	fmt.Println("ConnId [", c.ConnID, "] Connection.Stop()...")
	if c.isClose {
		return
	}
	c.isClose = true
	err := c.Conn.Close()
	if err != nil {
		fmt.Println("Connection.Stop() error:", err)
		return
	}
	c.ExistChan <- struct{}{}
	close(c.ExistChan)
}

func (c *Connection) Send(data []byte) error {
	return nil
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
