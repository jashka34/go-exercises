package solution18

import "sort"

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })
	return userIDs
}
