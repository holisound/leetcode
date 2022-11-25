/*
https://leetcode.cn/problems/filter-restaurants-by-vegan-friendly-price-and-distance/
1333. 餐厅过滤器
*/
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
    idTo := map[int][]int{}
    indices := []int{}
    for _, r := range restaurants {
        if (r[2] != 1 && 1 == veganFriendly) || r[3]>maxPrice || r[4] > maxDistance {
            continue
        }
        indices = append(indices, r[0])
        idTo[r[0]] = []int{r[0], r[1]}
    }
    sort.Slice(indices, func (x, y int)bool{
        i, j := indices[x], indices[y]
        a, b := idTo[i], idTo[j]
        if a[1] == b[1] {
            return a[0] > b[0]
        }
        return a[1] > b[1]
    })
    return indices
}
