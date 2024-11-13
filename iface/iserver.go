package iface

// IServer is an interface that defines the standard operations of a server.
// It includes methods to start the server, stop the server, and provide service.
type IServer interface {
	// Start initializes the server's operational environment and begins listening for requests.
	// This method does not accept parameters or return values.
	Start()

	// Stop safely shuts down the server, releasing resources occupied during operation.
	// This method does not accept parameters or return values.
	Stop()

	// Serve is the main method for the server to provide services.
	// It handles incoming requests and processes them.
	// This method does not accept parameters or return values.
	Serve()
}
