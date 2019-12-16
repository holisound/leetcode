class MapSum:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.trie = {}

    def insert(self, key: str, val: int) -> None:
        def is_key(node, key):
            for u in key:
                if u not in node:
                    return False
                node = node[u]
            return "is_key" in node
        
        node = self.trie
        
        if is_key(node, key):
            for u in key:
                node[u]['value'] = val
                node = node[u]
        else:
            for u in key:
                if u not in node:
                    node[u] = {"value": 0}
                node[u]['value'] += val
                node = node[u]
            node['is_key'] = {}

    def sum(self, prefix: str) -> int:
        node = self.trie
        for p in prefix:
            if not p in node:
                return 0
            node = node[p]
        return node["value"]
