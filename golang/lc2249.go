/*
https://leetcode.cn/problems/count-lattice-points-inside-a-circle/
2249. 统计圆内格点数目
*/
func countLatticePoints(circles [][]int) int {
    res := 0
    for i := 0; i <= 200; i++ {
        for j := 0; j <= 200; j++ {
            for _, c := range circles {
                if isPointInCircle(i, j, c[0], c[1], c[2]) {
                    res++
                    break
                }
            }
        }
    }
    return res
}

func isPointInCircle(px, py, cx, cy, cr int) bool {
    /* 判断点(px,py)是否在以(cx,cy)为中心，半径cr的圆内（上）*/
    return (px-cx)*(px-cx)+(py-cy)*(py-cy) <= cr*cr
}
