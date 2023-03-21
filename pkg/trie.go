package go_pruning_trie

type NodeList map[string]*Node

type Node struct {
	List NodeList
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

func (prt *PruningRadixTrie) AddTerm(term string, termFrequencyCount uint64) {
	nodeList := NodeList{}
	prt._add_term(prt.trie, term, termFrequencyCount, 0, 0, nodeList)
}

func (prt *PruningRadixTrie) UpdateMaxCounts(nodeList NodeList, termFrequencyCount uint64) {
	// beware the elusive pointers
	for key := range nodeList {
		node := nodeList[key]
		if termFrequencyCount > node.TermFrequencyCountChildMax {
			node.TermFrequencyCountChildMax = termFrequencyCount
		}
	}
}

// (Node curr, String term, long termFrequencyCount, int id, int level, List<Node> nodeList)
func (prt *PruningRadixTrie) _add_term(curr *Node, term string, termFrequencyCount uint64, id int, level int, nodeList NodeList) {

}

func PruningTrieFind(token string) string {
	return "lero"
}
