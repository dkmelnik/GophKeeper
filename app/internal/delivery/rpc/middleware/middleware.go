package middleware

// HandlerFunc is a function type that matches the signature of RPC methods.
type HandlerFunc func(args interface{}, reply interface{}) error

// Middleware defines a function to process middleware.
type Middleware func(HandlerFunc) HandlerFunc

// Chain applies middlewares to an RPC handler.
func Chain(handler HandlerFunc, middlewares ...Middleware) HandlerFunc {
	// Apply middlewares in reverse order.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
