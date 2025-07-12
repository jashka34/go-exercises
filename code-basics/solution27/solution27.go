package solution27

type Parent struct {
	Name     string
	Children []Child
}

type Child struct {
	Name string
	Age  int
}

func CopyParent(p *Parent) Parent {
	ret := Parent{}
	if p == nil {
		return ret
	}
	// fmt.Printf("p: %+v\n", p)
	ret.Name = p.Name
	for _, ch := range p.Children {
		// fmt.Printf("%+v\n", p)
		// fmt.Printf("%d: \n", i)
		rch := Child{}
		rch.Name = ch.Name
		rch.Age = ch.Age
		ret.Children = append(ret.Children, rch)
	}

	return ret
}
