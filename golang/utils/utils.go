package utils

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Max(a ...int) int {
	x := a[0]
	for _, n := range a {
		if n > x {
			x = n
		}
	}
	return x
}

func Min(a ...int) int {
	x := a[0]
	for _, n := range a {
		if n < x {
			x = n
		}
	}
	return x
}

func Reverse(bits []byte) {
	for i, j := 0, len(bits)-1; i < j; i, j = i+1, j-1 {
		bits[i], bits[j] = bits[j], bits[i]
	}
}

// 曼哈顿距离
func ManhattanDist(x1, y1, x2, y2 int) int {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	return x2 - x1 + y2 - y1
}

var (
	dx [4]int = [4]int{0, 0, 1, -1}
	dy [4]int = [4]int{1, -1, 0, 0}
)

type Pos struct {
	R, C, Value int
}
type BFSUtil struct {
	R, C    int
	Visited [][]int
	Matrix  any
}

func (bfu *BFSUtil) IsGoodPos(pos *Pos) bool {
	return pos.R > -1 && pos.R < bfu.R && pos.C > -1 && pos.C < bfu.C
}
func (bfu *BFSUtil) IsFirstVisited(pos *Pos) bool {
	return bfu.Visited[pos.R][pos.C] == 1
}

func (bfu *BFSUtil) SetValueInt(pos *Pos, value int) {
	mtx := bfu.Matrix.([][]int)
	mtx[pos.R][pos.C] = value
}
func (bfu *BFSUtil) SetValueByte(pos *Pos, value byte) {
	mtx := bfu.Matrix.([][]byte)
	mtx[pos.R][pos.C] = value
}
func (bfu *BFSUtil) GetNeighbours(pos *Pos) (res []*Pos) {
	for i := 0; i < 4; i++ {
		newPos := &Pos{pos.R + dx[i], pos.C + dy[i], pos.Value + 1}
		if bfu.IsGoodPos(newPos) {
			bfu.Visited[newPos.R][newPos.C] += 1
			res = append(res, newPos)
		}
	}
	return
}

func Swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
