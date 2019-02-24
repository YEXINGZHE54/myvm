package unsafe

import (
	"fmt"
)

var (
	memories = map[int64][]byte{}
	lastaddr = int64(1024)
)

func alloc(size int64) (addr int64) {
	memories[lastaddr] = make([]byte, size)
	addr = lastaddr
	lastaddr = lastaddr + size
	return
}

func at(addr int64) []byte {
	for start, buf := range memories {
		end := start + int64(len(buf))
		if start <= addr && addr < end {
			return buf[addr - start:]
		}
	}
	panic(fmt.Sprintf("invalid address %d", addr))
}

func free(addr int64) {
	if _, ok := memories[addr]; !ok {
		panic(fmt.Sprintf("invalid free addr: %d", addr))
	}
	delete(memories, addr)
}