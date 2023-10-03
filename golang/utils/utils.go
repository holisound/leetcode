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
	nrow, ncol int
	visited    [][]int
	Matrix     any
}

func NewBFSUtil(matrix any) BFSUtil {
	var nrow, ncol int
	if _, ok := matrix.([][]int); ok {
		nrow, ncol = len(matrix.([][]int)), len(matrix.([][]int)[0])
	} else if _, ok := matrix.([][]byte); ok {
		nrow, ncol = len(matrix.([][]byte)), len(matrix.([][]byte)[0])
	}
	visited := make([][]int, nrow)
	for i := range visited {
		visited[i] = make([]int, ncol)
	}
	return BFSUtil{nrow, ncol, visited, matrix}

}

func (bfu *BFSUtil) IsGoodPos(pos *Pos) bool {
	return pos.R > -1 && pos.R < bfu.nrow && pos.C > -1 && pos.C < bfu.ncol
}
func (bfu *BFSUtil) IsFirstVisited(pos *Pos) bool {
	return bfu.visited[pos.R][pos.C] == 1
}
func (bfu *BFSUtil) ValidateAddons(pos *Pos) bool {
	// 自定义方法
	return true
}
func (bfu *BFSUtil) Validate(pos *Pos) bool {
	return bfu.IsFirstVisited(pos) && bfu.ValidateAddons(pos)
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
			bfu.visited[newPos.R][newPos.C] += 1
			res = append(res, newPos)
		}
	}
	return
}
