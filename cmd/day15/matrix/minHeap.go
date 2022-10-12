package matrix

type MinHeap struct {
	list []*Node
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
	if i >= len(h.list) {
		return
	}

	leftIdx, rightIdx := h.LeftChild(i), h.RightChild(i)

	if rightIdx >= len(h.list) || leftIdx >= len(h.list) {
		return
	}

	leftValue, rightValue, currValue := h.At(leftIdx).Weight, h.At(rightIdx).Weight, h.At(i).Weight

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
	if h.At(i).Weight < h.At(parent).Weight {
		h.swap(parent, i)
		h.heapifyUp(parent)
	}
}

func (h *MinHeap) Insert(node *Node) {
	h.list = append(h.list, node)
	h.heapifyUp(len(h.list) - 1)
}

func (h *MinHeap) Pop() *Node {
	if len(h.list) == 0 {
		return nil
	}

	node := h.At(0)

	lastIndex := len(h.list) - 1
	h.list[0] = h.list[lastIndex]
	h.list = h.list[:lastIndex]

	h.heapifyDown(0)
	return node
}
