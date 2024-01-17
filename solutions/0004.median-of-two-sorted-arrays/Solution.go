package solutions

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 > l2 {
		l1, l2 = l2, l1
		nums1, nums2 = nums2, nums1
	}
	odd := ((l1 + l2) & 1) == 1
	half := (l1 + l2) >> 1
	if l1 == 0 {
		if odd {
			return float64(nums2[half])
		} else {
			return (float64(nums2[half-1]) + float64(nums2[half])) / 2
		}
	}

	left, right := 0, l1
	var a, b, x, y int
	for left < right {
		m1 := (left + right) >> 1
		m2 := half - m1
		b = nums1[m1]
		if m1 == 0 {
			a = nums2[m2-1]
		} else {
			a = nums1[m1-1]
		}
		if m2 == 0 {
			x, y = a, nums2[m2]
		} else if m2 == l2 {
			x, y = nums2[m2-1], b
		} else {
			x, y = nums2[m2-1], nums2[m2]
		}

		if x <= b && a <= y {
			break
		} else if x > b {
			left = m1 + 1
		} else {
			right = m1
		}
	}

	if left == l1 {
		a, b, y = nums1[l1-1], nums2[half-l1], nums2[half-l1]
		if half-l1 == 0 {
			x = a
		} else {
			x = nums2[half-l1-1]
		}
	}

	if odd {
		return float64(min(b, y))
	} else {
		return (float64(max(a, x)) + float64(min(b, y))) / 2
	}
}
