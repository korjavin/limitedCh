package main

import "testing"

var sizes = []int{1, 100, 2000}

func BenchmarkLen(b *testing.B) {
	for s := range sizes {
		ch := newLimitedChan(s)
		for i := 0; i < b.N; i++ {
			ch.WriteWithLen(i)
		}
	}
}
func BenchmarkDefault(b *testing.B) {
	for s := range sizes {
		ch := newLimitedChan(s)
		for i := 0; i < b.N; i++ {
			ch.WriteWithDefault(i)
		}
	}
}
