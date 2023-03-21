package go_pruning_trie

type Node struct {
	List []map[string]*Node
	//Does this node represent the last character in a word?
	//0: no word; >0: is word (termFrequencyCount)
	TermFrequencyCount         uint64
	TermFrequencyCountChildMax uint64
}

func newNode(termFrequencyCount uint64) (*Node, error) {
	node := Node{TermFrequencyCount: termFrequencyCount, TermFrequencyCountChildMax: 0}

	return &node, nil
}

type PruningRadixTrie struct {
	TermCount       uint64
	TermCountLoaded uint64
	trie            *Node
}

func newPruningRadixTrie() (*PruningRadixTrie, error) {
	rootNode, _ := newNode(0)
	prt := PruningRadixTrie{TermCount: 0, TermCountLoaded: 0, trie: rootNode}
	return &prt, nil
}

func PruningTrieFind(token string) string {
	return "lero"
}
