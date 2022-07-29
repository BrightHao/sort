package heap

import "errors"

type MinHeap struct {
	Size int
	Heap []interface{}
	More func(x, y interface{}) bool
}

func HeapSort(size int, arr []interface{}, f func(x, y interface{}) bool) ([]interface{}, error) {
	if size == 0 {
		return nil, errors.New("empty heap")
	}
	h := &MinHeap{Size: size, Heap: make([]interface{}, size), More: f}
	for _, i := range arr {
		h.add(i)
	}
	return h.Heap, nil
}

func (h *MinHeap) heapify(i int) {
	len := h.Size
	heap := h.Heap

	left := 2*i + 1
	right := 2*i + 2
	if left < len && h.More(heap[i], heap[left]) {
		h.swap(left, i)
		h.heapify(left)
	}
	if right < len && h.More(heap[i], heap[right]) {
		h.swap(right, i)
		h.heapify(right)
	}
}

func (h *MinHeap) add(obj interface{}) {
	size := h.Size
	heap := h.Heap
	if heap[0] == nil {
		for i := 0; i < size; i++ {
			heap[i] = obj
		}
		return
	}

	if h.More(obj, heap[0]) {
		heap[0] = obj
	}

	h.heapify(0)
}

func (h *MinHeap) swap(x, y int) {
	nums := h.Heap
	nums[x], nums[y] = nums[y], nums[x]
}
