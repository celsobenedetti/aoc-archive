package common

type FishList []int

func (f FishList) GetLen() int {
	return len(f)
}

func (f *FishList) addFish(n int) {
	*f = append(*f, n)
}

func (f *FishList) AddNewFish() {
	f.addFish(8)
}

func MakeFishList(fishes []int) (list FishList) {
	for _, fish := range fishes {
		list.addFish(fish)
	}
	return list
}

func (fishes FishList) RunDayLifeClycle() FishList {
	newFishes := make(FishList, 0)

	for i := 0; i < len(fishes); i++ {
		if fishes[i] == 0 {
			fishes[i] = 6
			newFishes.AddNewFish()
		} else {
			fishes[i]--
		}
		newFishes = append(newFishes, fishes[i])
	}
	return newFishes
}
