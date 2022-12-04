/*
https://leetcode.cn/problems/maximum-matrix-sum/solution/
1975. 最大方阵和
*/ 
func maxMatrixSum(matrix [][]int) int64 {
    cntNeg := 0
    minAbs := 0x3f3f3f
    n := len(matrix)
    tot:=0
    for i:=0;i<n;i++{
        for j:=0;j<n;j++{
            if matrix[i][j]<0{
                cntNeg++
            }
            tot+=abs(matrix[i][j])
            if abs(matrix[i][j])<minAbs{
                minAbs=abs(matrix[i][j])
            }
        }
    }
    if cntNeg&1==1{
        return int64(tot-minAbs*2)
    }
    return int64(tot)
}

func abs(a int)int{
    if a<0{
        return -a
    }
    return a
}
