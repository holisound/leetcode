/*
https://leetcode.cn/problems/maximum-sum-of-an-hourglass/
2428. 沙漏的最大总和
*/


func maxSum(grid [][]int) int {
    offsets :=[][]int{
        {-1,-1},{-1,0},{-1,1},
        {1,-1},{1,0},{1,1},
    }
    res:=0
    n, m := len(grid), len(grid[0])
    for i:=0;i<n;i++{
        for j:=0;j<m;j++{
            flag:=false
            s:=grid[i][j]
            for _, offset := range offsets {
                r, c := i+offset[0], j+offset[1]
                if r>=n || r<0 || c>=m || c<0{
                    flag=true
                    break
                } else {
                    s += grid[r][c]
                }
            }
            if !flag && s > res{
                res = s
            }
        }
    }
    return res
}
