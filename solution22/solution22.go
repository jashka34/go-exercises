package solution22

func shiftASCII(s string, step int) string {
	var ret string
	// fmt.Println(s)
	for i := range s {
		// for i := 0; i < len(s); i++ {
		// fmt.Println(s[i])
		ret += string(s[i] + byte(step))
	}
	return ret
}
