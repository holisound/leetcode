import lls "github.com/emirpasic/gods/stacks/linkedliststack"

func summaryRanges(nums []int) []string {
	stack := lls.New()
	ans := []string{}

	for i := 0; i <= len(nums); i++ {
		for val, ok := stack.Peek(); ok && (i == len(nums) || val != nums[i]-1); val, ok = stack.Peek() {
			if stack.Size() > 1 {
				ans = append(ans, fmt.Sprintf("%d->%d", val.(int)-stack.Size()+1, val))
			} else {
				ans = append(ans, fmt.Sprintf("%d", val))
			}

			stack.Clear()
		}
		if i < len(nums) {
			stack.Push(nums[i])
		}

	}

	return ans
}
