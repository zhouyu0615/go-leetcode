package main

import (
	"fmt"
	"sort"
	"testing"
)

type question1 struct {
	para1
	ans1
}

// para 是参数
// one 代表第一个参数
type para1 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one []int
}

func testEq(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func Test_TwoSums(t *testing.T) {

	qs := []question1{
		{
			para1{[]int{47, 2, 4}, 6},
			ans1{[]int{1, 2}},
		},

		{
			para1{[]int{3, 2, 4}, 5},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans1{[]int{1, 3}},
		},

		{
			para1{[]int{0, 1}, 1},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 4}, 5},
			ans1{[]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		_, p := q.ans1, q.para1
		actual := twoSum(p.nums, p.target)
		if !testEq(actual, q.ans1.one) {
			t.Fatalf("【input】:%v       【output】:%v,  expect:%v \n", p, actual, q.ans1.one)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, actual)
	}
	fmt.Printf("\n\n\n")
}
