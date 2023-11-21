package main

import "fmt"

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return[]int{}
}

func main() {
    result := twoSum([]int{1, 2, 4, 6, 4}, 10)
    fmt.Println(result)
}