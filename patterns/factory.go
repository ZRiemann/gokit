package patterns

import (
	"fmt"
	"io"
)

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage( /*...*/ )
	case DiskStorage:
		return newDiskStorage( /*...*/ )
	default:
		return newTempStorage( /*...*/ )
	}
}

func newMemoryStorage() Store {
	fmt.Println("newMemoryStorage")
	return nil
}

func newDiskStorage() Store{
	fmt.Println("newDiskStorage")
	return nil
}

func newTempStorage() Store{
	fmt.Println("newTempStorage")
	return nil
}
