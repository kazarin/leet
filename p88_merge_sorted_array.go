package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, m)
	copy(nums, nums1)

	a := 0

	b := 0
	var mm int

	for i := 0; i < m+n; i++ {
		if a > m-1 {
			mm = nums2[b]
			b++
		} else if b > n-1 {
			mm = nums[a]
			a++
		} else if nums[a] > nums2[b] {
			mm = nums2[b]
			b++
		} else {
			mm = nums[a]
			a++
		}
		nums1[i] = mm

	}

	return

}
