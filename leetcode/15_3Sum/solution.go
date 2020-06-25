package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}
	p1val := nums[0] - 1
	for p1 := 0; p1 < len(nums); p1++ {
		if nums[p1] == p1val {
			continue
		}
		p1val = nums[p1]
		p2 := p1 + 1
		p3 := len(nums) - 1
		for p2 < p3 {
			val := nums[p1] + nums[p2] + nums[p3]
			if val == 0 {
				result = append(result, []int{nums[p1], nums[p2], nums[p3]})
			}
			if val <= 0 {
				p2val := nums[p2]
				for p2 < len(nums) && nums[p2] == p2val {
					p2++
				}
			}
			if val >= 0 {
				p3val := nums[p3]
				for p3 >= 0 && nums[p3] == p3val {
					p3--
				}
			}
		}
	}
	return result
}

func main() {

}
