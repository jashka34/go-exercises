package solution15

func RemoveFirstElement(slice []int) []int {
	if len(slice) == 0 {
		return []int{}
	} else {
		return slice[1:]
	}
}
