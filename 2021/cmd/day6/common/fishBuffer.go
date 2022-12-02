package common

type FishBuffer struct {
	Buffer []int
	Count  int
}

var length = 9

func MakeFishBuffer(input []int) (f FishBuffer) {
	f.Buffer = make([]int, length)
	f.Count = len(input)

	for _, v := range input {
		f.Buffer[v]++
	}

	return f
}

func (f FishBuffer) RunRoundLifeCycle() (result FishBuffer) {
	result.Buffer = make([]int, length)
	numberCreated := f.Buffer[0]

	for i := 0; i < length-1; i++ {
		result.Buffer[i] = f.Buffer[i+1]
	}

	result.Buffer[6] += numberCreated
	result.Buffer[length-1] = numberCreated

	result.Count = f.Count + numberCreated

	return result
}
