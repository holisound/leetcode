/*
https://leetcode.cn/problems/smallest-subsequence-of-distinct-characters/
1081. 不同字符的最小子序列
*/
func smallestSubsequence(s string) string {
    inStack:=map[int]bool{}
    st := &Stack{}
    for i, c := range s {
        k:=int(c)
        if inStack[k] {
            continue
        }
        for st.Size()>0 && st.Peek()>k && strings.IndexRune(s[i+1:], rune(st.Peek()))!=-1{
            inStack[st.Pop()]=false
        }

        st.Push(k)
        inStack[k]=true
    }
    sb:=&strings.Builder{}
    for st.Size()>0 {
        sb.WriteRune(rune(st.Pop()))
    }
    return reverseString(sb.String())
}

func reverseString(s string)string{
    chars:=[]byte(s)
    for i,j:=0,len(s)-1;i<j;i,j=i+1,j-1{
        chars[i],chars[j]=chars[j],chars[i]
    }
    return string(chars)
}
// 栈, 将int作为中间数据类型，在rune或byte之间转换
type Stack struct {
    data    []int
    sizeCur int
    sizeMax int
}
func (s *Stack) Push(x int){
    s.sizeCur++
    if s.sizeCur > s.sizeMax {
        s.data = append(s.data, x)
        s.sizeMax = s.sizeCur
    } else {
        s.data[s.sizeCur-1] = x
    }
}
func (s *Stack) Pop() int{
    res := s.data[s.sizeCur-1]
    s.sizeCur--
    return res
}
func (s *Stack) Size() int {
    return s.sizeCur
}
func (s *Stack) Peek() int{
    return s.data[s.sizeCur-1]
}
