package source

import "io"

// Source represents image source.
type Source interface {
	// Random returns an image handle randomly choiced.
	Random() (io.ReadCloser, error)
}
