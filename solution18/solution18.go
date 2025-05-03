package solution18

import (
	"fmt"
	"sort"
)

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })

	// prevId int;
	for i, id := range userIDs {
		fmt.Println(i, ":", id)
		if i > 0 {
			if userIDs[i-1] == userIDs[i] {
				fmt.Println("ops: ", userIDs[i])
			}
		}
	}

	return userIDs
}
