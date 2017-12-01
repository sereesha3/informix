package cache

// A Cacher is responsible for utilizing Queuers to cache messages across
// any number of producing channels, providing a fast read only view of the
// data.
type Cacher interface {

	// TODO:
	Keys() []string

	// Register lets the caller start a caching routine on the
	// provided directive.
	Register(*Directive)

	// Unregister lets the caller stop a caching routine that is currently
	// running for the provided directive.
	Unregister(*Directive)
}