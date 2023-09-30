package main

type Trie struct {
	nodes [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, c := range word {
		if cur.nodes[c-'a'] == nil {
			cur.nodes[c-'a'] = &Trie{}
		}
		cur = cur.nodes[c-'a']
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.SearchCommon(word)
	return node != nil && node.isEnd
}

func (this *Trie) SearchCommon(word string) *Trie {
	cur := this
	for _, c := range word {
		if cur.nodes[c-'a'] == nil {
			return nil
		}
		cur = cur.nodes[c-'a']
	}
	return cur
}

func (this *Trie) StartsWith(word string) bool {
	return nil != this.SearchCommon(word)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
