package main

import (
	"runtime/debug"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	debug.SetGCPercent(-1)
	main()
}
