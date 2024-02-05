package solutions

import "sort"

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	merging, scanning := 0, 0
	for scanning < len(intervals) {
		if intervals[merging][1] >= intervals[scanning][0] {
			if intervals[scanning][1] > intervals[merging][1] {
				intervals[merging][1] = intervals[scanning][1]
			}
		} else {
			merging++
			intervals[merging] = intervals[scanning]
		}
		scanning++
	}
	return intervals[:merging+1]
}
