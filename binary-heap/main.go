package main

func insert(heap []int, element int) {
	heap = append(heap, element)
	makeHeapCompliant(heap)
}

func makeHeapCompliant(heap []int) {
	if len(heap) == 1 {
		return
	}

	if heap[len(heap) - 1] > heap[posParent] {
	}
}

func main() {

}
