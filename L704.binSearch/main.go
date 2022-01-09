package main

//import "fmt"

func search(nums []int, target int) int {
	var low, high, mid int
	low = 0
	high = len(nums) - 1

	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		}
	}

	return -1
}

/*
func main() {

	var nums []int = []int{1, 3, 4, 5, 6, 67, 78}

	fmt.Printf("result = %d", search(nums, 7))

}
*/
