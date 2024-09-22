package main

import "fmt"

func threeSum(nums []int) [][]int {
	var res [][]int
	nums = sortNums(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := 0 - nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[left]+nums[right] > target {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func sortNums(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

//import (
//	"sort"
//)
//
//func threeSum(nums []int) [][]int {
//	var res [][]int
//	sort.Ints(nums) // 使用内置的排序函数进行升序排序
//	for i := 0; i < len(nums); i++ {
//		// 跳过重复的元素
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		target := 0 - nums[i]
//		left, right := i+1, len(nums)-1
//		for left < right {
//			sum := nums[left] + nums[right]
//			if sum == target {
//				res = append(res, []int{nums[i], nums[left], nums[right]})
//				left++
//				right--
//				// 跳过重复的元素
//				for left < right && nums[left] == nums[left-1] {
//					left++
//				}
//				for left < right && nums[right] == nums[right+1] {
//					right--
//				}
//			} else if sum < target {
//				left++
//			} else {
//				right--
//			}
//		}
//	}
//	return res
//}

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	fmt.Println(threeSum(nums))
}
