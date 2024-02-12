package solutions

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	i, i1, i2 := m+n-1, m-1, n-1
	for i1 >= 0 && i2 >= 0 {
		if nums1[i1] > nums2[i2] {
			nums1[i] = nums1[i1]
			i1--
		} else {
			nums1[i] = nums2[i2]
			i2--
		}
		i--
	}
	for i2 >= 0 {
		nums1[i2] = nums2[i2]
		i2--
	}
}
