package leetcode

func push(queue []int, val int) {
	queue[1] = queue[0]
	queue[0] = val
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	p1, p2, visited, n := 0, 0, 0, len(nums1)+len(nums2)
	target := n/2 + 1
	queue := make([]int, 2, 2)
	for visited <= target && visited < n {
		visited++
		if p1 == len(nums1) {
			push(queue, nums2[p2])
			p2++
		} else if p2 == len(nums2) {
			push(queue, nums1[p1])
			p1++
		} else if nums1[p1] < nums1[p2] {
			push(queue, nums1[p1])
			p1++
		} else {
			push(queue, nums2[p2])
			p2++
		}
	}

}
