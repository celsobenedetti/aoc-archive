package matrix

type MinHeap struct {
	list   []*Node
	length int
}

func CreateHeap(n int) MinHeap {
	return MinHeap{list: make([]*Node, n), length: 0}
}

func (h MinHeap) At(i int) *Node {
	return h.list[i]
}

func (h MinHeap) Parent(i int) int {
	return (i - 1) / 2
}

func (h MinHeap) LeftChild(i int) int {
	return (i * 2) + 1
}

func (h MinHeap) RightChild(i int) int {
	return (i * 2) + 2
}

func (h *MinHeap) swap(parent, child int) {
	h.list[parent], h.list[child] = h.list[child], h.list[parent]
}

func (h *MinHeap) heapifyDown(i int) {
	if i >= h.length {
		return
	}

	leftIdx, rightIdx := h.LeftChild(i), h.RightChild(i)

	if rightIdx >= h.length || leftIdx >= h.length {
		return
	}

	leftValue, rightValue, currValue := h.At(leftIdx).Dist, h.At(rightIdx).Dist, h.At(i).Dist

	if leftValue >= rightValue && currValue >= rightValue {
		h.swap(i, rightIdx)
		h.heapifyDown(rightIdx)
	}
	if rightValue >= leftValue && currValue >= leftValue {
		h.swap(i, leftIdx)
		h.heapifyDown(leftIdx)
	}
}

func (h *MinHeap) heapifyUp(i int) {
	if i == 0 {
		return
	}
	parent := h.Parent(i)
	if h.At(i).Dist < h.At(parent).Dist {
		h.swap(parent, i)
		h.heapifyUp(parent)
	}
}

func (h *MinHeap) Insert(node *Node) {
	if len(h.list) == h.length {
		h.list = append(h.list, make([]*Node, h.length)...)
	}

	h.list[h.length] = node
	h.length++
	h.heapifyUp(h.length - 1)
}

func (h *MinHeap) Pop() *Node {
	if h.length == 0 {
		return nil
	}

	node := h.At(0)

	lastIndex := h.length - 1
	h.swap(0, lastIndex)
	h.list[lastIndex] = nil
	h.length--

	h.heapifyDown(0)
	return node
}
