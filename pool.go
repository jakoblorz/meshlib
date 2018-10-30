package meshlib

import (
	"fmt"
	"io"
	"net/rpc"
)

// Pool manages active *rpc.Client. It also serves as register for all current nodes
// in the topology. Initializing a new address should not automatically create a client/connection.
type Pool interface {
	io.Closer

	// Init initializes a specific address; this could be a net.Addr
	// or any other fmt.Stringer interface
	Init(fmt.Stringer)

	// Clear removes a specific address from the pool; all pooled connections
	// will be closed
	Clear(fmt.Stringer)

	// Aquire pulls a *rpc.Client from the pool and returns a return callback
	// which needs to be called once the client is not needed again
	Aquire(fmt.Stringer) (*rpc.Client, func(), error)

	// Release can be called to return a *rpc.Client back to the pool manually
	Release(fmt.Stringer, *rpc.Client) error

	// List returns a list of initialized addresses
	List() []fmt.Stringer

	// Call is a shorthand func allowing the call of once rpc call, already integrated
	// into the pool aquire - release flow
	Call(fmt.Stringer, string, interface{}, interface{}) error
}
