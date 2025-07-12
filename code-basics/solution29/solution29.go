package solution29

type Person struct {
	Age uint8
}

type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	ret := make(map[uint8]int, len(pl))
	for _, p := range pl {
		// fmt.Printf("%d %+v\n", i, p)
		ret[p.Age]++
	}
	return ret
}
