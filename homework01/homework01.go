package homework01

import (
	"fmt"
	"sort"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// TODO: implement
	if len(nums) == 0 {
		return 0
	}
	resultMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		key := nums[i]
		if _, ok := resultMap[key]; ok {
			delete(resultMap, key)
		} else {
			resultMap[key] = 1
		}
	}
	for key := range resultMap {
		return key
	}
	return 0
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	fmt.Println(x)
	reverse := 0
	originalX := x
	for {
		a := x % 10
		b := x / 10
		reverse = reverse*10 + a
		if b == 0 {
			break
		}
		x = x / 10
	}
	// TODO: implement
	return reverse == originalX
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	// TODO: implement
	db := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack []rune

	for _, char := range s {
		if _, exists := db[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != db[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// TODO: implement
	maxPrefix := ""
	for _, str := range strs {
		if len(maxPrefix) == 0 {
			maxPrefix = str
			continue
		}
		for i := 0; i < len(maxPrefix) && i < len(str); i++ {
			if maxPrefix[i] != str[i] {
				maxPrefix = maxPrefix[:i]
				break
			}
		}
	}

	return maxPrefix
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// TODO: implement
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)

	return digits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// TODO: implement

	x := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[x] {
			x++
			nums[x] = nums[i]
		}
	}
	return x + 1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// TODO: implement
	if len(intervals) <= 1 {
		return intervals
	}
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{}
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= current[1] {
			if intervals[i][1] > current[1] {
				current[1] = intervals[i][1]
			}
		} else {
			result = append(result, current)
			current = intervals[i]
		}
	}
	result = append(result, current)
	return result
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// TODO: implement
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil
}
