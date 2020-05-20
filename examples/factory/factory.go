package main

import (
	"gokit/patterns"
)

func main() {
	patterns.NewStore(patterns.DiskStorage)
	patterns.NewStore(patterns.DiskStorage)
	patterns.NewStore(patterns.DiskStorage)
}

