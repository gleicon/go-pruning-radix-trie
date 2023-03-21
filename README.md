## go-pruning-radix-trie

### What it is all about ?

Go implementation of [Pruning Radix Trie, originally by Wolf Garbe](https://github.com/wolfgarbe/PruningRadixTrie). The breakdown from the author can be [found in this blog post](https://seekstorm.com/blog/pruning-radix-trie/).

In the author's word: 

```
A Radix Trie or Patricia Trie is a space-optimized trie (prefix tree). A Pruning Radix trie is a novel Radix trie algorithm, that allows pruning of the Radix trie and early termination of the lookup.
```

It is geared for applications like autocomplete, fast fuzzy searches and search that starts with short prefixes and are too costly for a regular Patricia Trie, hence pruning and fine-grained lookup control to deal with bigger datasets.

### Code

[![Go Report Card](https://goreportcard.com/badge/github.com/gleicon/go-pruning-radix-trie?style=flat-square)](https://goreportcard.com/report/github.com/gleicon/go-pruning-radix-trie)
