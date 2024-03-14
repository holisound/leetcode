## Leetcode: Lead a way of sloving problems
> 算法思考/思维 60%
- 1、读题、审题，信息（已知条件）提取（包括：正文、示例和提示）
    1、读题之后，有比较清晰的解题思路
    2、【无法理解】题目中已给出的已经条件对于解题思路的帮助和意图
    3、晦涩的表达措辞可能是一种误导（思路陷阱）
    4、通过阅读题解辅助理解题意，大致能看懂题解中的解题思路
    5、别人的思路只能作为参考，想方设法把别人思路中的亮点进行提炼转化，补充优化到你自己的解题思路
    6、任何一个经过加工、咀嚼、修饰的结论或者思路都不可能是真正属于你的；解题过程中自己踩坑避坑的过程，是真正属于你的解题思路。
    7、从题目要求的【输出结果】入手。通过输出结果的数据结构，能帮助了解题意。
> 动手编码/调试 40%
- 写程序
- 建立节点映射边长，超内存，矩阵超了，用哈希表可以过
- 枚举子串边界和枚举子串元素
```golang
//枚举子串边界
for i:=0;i<n;i++{
    for j:=0;j<i+1;j++{
    }
}
//枚举子串元素
for i:=0;i<n;i++{
    for j:=i;j<n;j++{
    }
}
// 三元问题 转化成 二元

for i := 0; i < m; i++ {
    // 顺时针坐标转置 i, j -> j, m-1-i
    col := m - 1 - i
    for j := 0; j <= n; j++ { // 元素标记位和越界标记位
        if j == n || box[i][j] == '*' {

}}}
// byte类型，不能直接二进制运算 
// 数组从后往前遍历的好处：在遍历同时需要删除元素的时候。不会影响后续元素的索引相对位置。

// 枚举二进制子集
for i:=1;i<=n;i++{
    for j:=i;j>0;j=(j-1)&i{
        
    }
}
```