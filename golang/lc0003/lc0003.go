package lc0003

/*
https://leetcode.cn/problems/longest-substring-without-repeating-characters/
3. Longest Substring Without Repeating Characters
*/

func lengthOfLongestSubstring(s string) int {
/*
we're supposed to find the longest substring without repeating characters.
so we want to keep the substring without repeating characters
Brute Force
    first, we can generate all the substrings
    and then we filter out the ones with repeating characters
    and next we can find out the longest substring from the remains
HashMap & Two Pointers
    we can remeber the indices of last occurrence of characters 

*/
    ans := 0
    lastOcc := map[rune]int{}
    k := 0
    for i, c := range s{
        if j, ok:=lastOcc[c]; ok && j + 1 > k{
            k = j+1
        }
        if i-k+1>ans{
            ans=i-k+1
        }
        lastOcc[c] = i
    }
    return ans
}
