package main

import (
	"sort"

	. "github.com/holisound/leetcode/utils"
)

/*
https://leetcode.cn/problems/get-watched-videos-by-your-friends/
1311. 获取你好友已观看的视频
*/
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	cnt := make(map[string]int)
	vis := make(map[int]bool)
	vis[id] = true

	que := new(Queue)
	que.Push(id)
	for i := 0; i < level; i++ {
		size := que.Size()
		for j := 0; j < size; j++ {
			for _, nid := range friends[que.Poll()] {
				if !vis[nid] {
					vis[nid] = true
					que.Push(nid)
				}
			}
		}
	}
	for que.Size() > 0 {
		for _, v := range watchedVideos[que.Poll()] {
			cnt[v]++
		}
	}

	ans := []string{}
	for k := range cnt {
		ans = append(ans, k)
	}
	sort.Slice(ans, func(i, j int) bool {
		if cnt[ans[i]] == cnt[ans[j]] {
			return ans[i] < ans[j]
		}
		return cnt[ans[i]] < cnt[ans[j]]
	})
	return ans
}
