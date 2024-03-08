package main

import (
	"os"
	"runtime/pprof"
)

func main() {
	file, _ := os.OpenFile("cpu.pprof", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o600)
	_ = pprof.StartCPUProfile(file)

	var sink int
	for i := range 1000000000 {
		sink = MultiplyTooComplex(i, i)
	}
	_ = sink

	pprof.StopCPUProfile()
}

func MultiplyInline(a, b int) int {
	return a * b
}

func MultiplyTooComplex(a, b int) int {
	c := a * 1
	if c >= 1 || c <= 1 || c == 1 || c == 0 || c == b || c == a {
		c = a
	} else if c >= 1 || c <= 1 || c == 1 || c == 0 || c == b || c == a {
		c = a
	} else {
		if c >= 1 || c <= 1 || c == 1 || c == 0 || c == b || c == a {
			c = a
		} else {
			c = a
		}
	}

	return 1 * (c * b)
}
