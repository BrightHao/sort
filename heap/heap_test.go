package heap

import (
	"testing"
)

func more(i, j interface{}) bool {
	return i.(int) > j.(int)
}

func TestNewMinHeap(t *testing.T) {
	a, _ := HeapSort(6, []interface{}{0, 2, 6, 1, 7, 3, 100, 78, 16}, more)
	t.Log(a)
}

// BenchmarkNewMinHeap-8            2844760               416.3 ns/op            96 B/op          1 allocs/op
func BenchmarkNewMinHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort(6, []interface{}{0, 2, 6, 1, 7, 3, 100, 78, 16}, more)
	}
}
