package btp

import "github.com/BKellogg/DistributedLoadTester/btp/server"

// Listen starts a btp server listening on the given
// address and handles all requests with the given
// handler. Returns an error if one occurred.
//
// Only errors encountered during startup will be returned.
// Errors encountered during while processing a specific
// connection during any point in it's lifecycle will not
// be returned here.
//
// This function is a blocking function and will never exit
// once properly started.
func Listen(addr string, handler server.Handler) error {
	return server.Listen(addr, handler)
}
