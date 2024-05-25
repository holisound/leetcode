package main

/*
示例 1：

输入：nums = [3,4,1,2,6], queries = [[0,4]]

输出：[false]

解释：

子数组是 [3,4,1,2,6]。2 和 6 都是偶数。

示例 2：

输入：nums = [4,3,1,6], queries = [[0,2],[2,3]]

输出：[false,true]

解释：

	子数组是 [4,3,1]。3 和 1 都是奇数。因此这个查询的答案是 false。
	子数组是 [1,6]。只有一对：(1,6)，且包含了奇偶性不同的数字。因此这个查询的答案是 true。

提示：

	1 <= nums.length <= 105
	1 <= nums[i] <= 105
	1 <= queries.length <= 105
	queries[i].length == 2
	0 <= queries[i][0] <= queries[i][1] <= nums.length - 1
*/
func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = 1
	cnt := 1
	for i := 1; i < n; i++ {
		if nums[i-1]&1^nums[i]&1 == 1 {
			cnt++
		} else {
			cnt = 1
		}
		prefix[i] = cnt
	}
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = prefix[q[1]] >= q[1]-q[0]+1
	}
	return ans
}

func main() {
}
