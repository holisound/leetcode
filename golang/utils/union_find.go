package utils

type UnionFind struct {
	fa []int
}

func NewUnionFind(n int) *UnionFind {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return &UnionFind{fa}
}
func (uf *UnionFind) Find(x int) int {
	fa := uf.fa
	if fa[x] != x {
		fa[x] = uf.Find(fa[x])
	}
	return fa[x]
}
func (uf *UnionFind) Merge(from, to int) {
	uf.fa[uf.Find(from)] = uf.Find(to)
}
func (uf *UnionFind) Connect(from, to int) bool {
	return uf.Find(from) == uf.Find(to)
}
