package qnet

import (
	"com.smyx/QNet/utils"
	"errors"
	"fmt"
	"net"
)

// Server represents a server instance and its basic configuration.
type Server struct {
	Name     string // Server name
	IP       string // Server IP address
	Port     int    // Server port number
	Protocol string // Server IP version, e.g., "tcp4" for IPv4, "tcp6" for IPv6
}

// Start initializes the server, sets up the listening address, and starts listening for connections.
// This method runs in a new goroutine and does not block the caller.
func (s *Server) Start() {
	go func() {
		// Combine IP and port to form the listening address
		address := fmt.Sprintf("%s:%d", s.IP, s.Port)
		tcpAddr, err := net.ResolveTCPAddr("tcp", address)
		if err != nil {
			fmt.Printf("[QNet] Server [%s] failed to start: %v\n", s.Name, err)
			return
		}
		// Listen on the specified IP version and address
		tcp, err := net.ListenTCP(s.Protocol, tcpAddr)
		if err != nil {
			fmt.Printf("[QNet] Server [%s] failed to listen: %v\n", s.Name, err)
			return
		}
		fmt.Printf("[QNet] Server [%s] is listening...\n", s.Name)
		for {
			// Accept new connections
			conn, err := tcp.AcceptTCP()
			if err != nil {
				fmt.Printf("[QNet] Server [%s] failed to accept connection: %v\n", s.Name, err)
				continue
			}
			// TODO Read received data
			nc := NewConnection(conn, utils.GenerateConnID(), CallBackToClient)
			nc.Start()
		}
	}()
}

// TODO 临时实现
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallBackToClient ... ")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

// Stop stops the server and prints a stop message.
// In the current implementation, it only prints the stop message and does not release resources.
func (s *Server) Stop() {
	fmt.Printf("[QNet] Server [%s] has stopped...\n", s.Name)
	// TODO Release connection resources, etc.
}

// Serve starts the server.
// It calls the Start method to initialize and start the server.
func (s *Server) Serve() {
	s.Start()
}

// NewServer creates and initializes a new Server instance.
// This function takes serverName as a parameter to set the server's name.
// The return value is a pointer to a Server type, containing the basic configuration of the server.
// Note that all server instances will use the default local address "127.0.0.1" and port 8080,
// as well as TCP protocol as the IP version.
func NewServer(serverName string) *Server {
	return &Server{
		Name:     serverName,
		IP:       "127.0.0.1",
		Port:     8080,
		Protocol: "tcp",
	}
}
