package main

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	mostPopular := 0
	cntViews := map[string]int{}
	targets := map[string]bool{}
	for i, view := range views {
		cntViews[creators[i]] += view
		if cntViews[creators[i]] > mostPopular {
			mostPopular = cntViews[creators[i]]
			targets = make(map[string]bool)
			targets[creators[i]] = true
		} else if cntViews[creators[i]] == mostPopular {
			targets[creators[i]] = true
		}
	}
	creatorViewId := map[string]string{}
	mostViews := map[string]int{}
	for i, c := range creators {
		if _, ok := targets[c]; ok {
			if _, ok := mostViews[c]; mostViews[c] < views[i] || !ok {
				mostViews[c], creatorViewId[c] = views[i], ids[i]
			} else if views[i] == mostViews[c] && ids[i] < creatorViewId[c] {
				creatorViewId[c] = ids[i]
			}
		}
	}
	res := [][]string{}
	for t := range targets {
		res = append(res, []string{t, creatorViewId[t]})
	}

	return res

}
