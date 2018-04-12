package btp

// Handler defines the type of function that can be
// used as a BTP Handler
type Handler func(ResponseWriter, *Request)
