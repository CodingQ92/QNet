package iface

import "net"

// IConnection is an interface that defines the operations for managing network connections.
// It provides methods to start and stop connections, retrieve TCP connection information,
// get a unique connection ID, obtain the remote address, and send data.
type IConnection interface {
	// Start initializes the network connection.
	Start()
	// Stop terminates the network connection.
	Stop()
	// GetTCPConnection returns the current TCP connection object.
	GetTCPConnection() *net.TCPConn
	// GetConnID returns a unique identifier for the current connection.
	GetConnID() uint32
	// RemoteAddr returns the remote network address of the current connection.
	RemoteAddr() net.Addr
	// Send sends data over the current connection.
	// data: The byte slice of the data to be sent.
	// Returns an error if data sending fails.
	Send(data []byte) error
}

// HandleFunc is a callback function type for processing data received over a TCP connection.
// It takes a TCPConn object and a byte slice as parameters, and returns an error.
// conn: The current TCP connection object.
// data: The received byte slice of data.
// Returns an error if data processing fails.
type HandleFunc func(*net.TCPConn, []byte, int) error
