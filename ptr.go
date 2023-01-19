// Package ptr implements a simple pointer instrumentation.
// As it is based on Go generics, the minimal Go version is 1.18.
package ptr

// From will extract a value designated by pointer provider.
// It will panic if the pointer provided is nil and the default isn't supplied.
func From[T any](pointer *T, dflt ...T) T {
	if pointer == (*T)(nil) && len(dflt) > 0 {
		return dflt[0]
	}
	return *pointer
}

// To will return a link to the value.
func To[T any](value T) (pointer *T) { return &value }
