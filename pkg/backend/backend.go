package backend

import "io"

type Backend interface {
	io.ReaderAt
	io.WriterAt

	Size() (int64, error)
	Sync() error
}

type Trimable interface {
	Trim(offset int64, length int64) error
}
