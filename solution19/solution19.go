package solution19

func UniqueUserIDs(userIDs []int64) []int64 {
	if len(userIDs) <= 1 {
		return userIDs
	}
	// sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })

	uniq := 0
	for i := range userIDs {
		// fmt.Println(i, ":", id, "uniq:", uniq)
		if i > 0 {
			if userIDs[i] != userIDs[uniq] {
				uniq++
				userIDs[uniq] = userIDs[i]
				// fmt.Println("userIDs:", userIDs)
			}
		}
	}

	return userIDs[:uniq+1]
}
