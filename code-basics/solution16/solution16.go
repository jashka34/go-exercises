package solution16

func Map(strs []string, mapFunc func(s string) string) []string {
	var ret []string
	for _, s := range strs {
		ret = append(ret, mapFunc(string(s)))
	}
	return ret
}
