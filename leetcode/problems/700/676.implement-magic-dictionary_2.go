package problems

/*
676. Implement Magic Dictionary
Implement a magic directory with buildDict, and search methods.

For the method buildDict, you'll be given a list of non-repetitive words to build a dictionary.

For the method search, you'll be given a word, and judge whether if you modify exactly one character into another character in this word, the modified word is in the dictionary you just built.

Example 1:
Input: buildDict(["hello", "leetcode"]), Output: Null
Input: search("hello"), Output: False
Input: search("hhllo"), Output: True
Input: search("hell"), Output: False
Input: search("leetcoded"), Output: False
Note:
You may assume that all the inputs are consist of lowercase letters a-z.
For contest purpose, the test data is rather small by now. You could think about highly efficient algorithm after the contest.
Please remember to RESET your class variables declared in class MagicDictionary, as static/class variables are persisted across multiple test cases. Please see here for more details.
*/

//MagicDictionary struct
type MagicDictionary struct {
	dicts map[int][][]rune
}

//Constructor Initialize your data structure here.
func Constructor() MagicDictionary {
	return MagicDictionary{}
}

//BuildDict Build a dictionary through a list of words
func (m *MagicDictionary) BuildDict(dict []string) {
	m.dicts = make(map[int][][]rune, len(dict))
	for _, d := range dict {
		l := len(d)
		arr, ok := m.dicts[l]
		if ok {
			arr = append(arr, []rune(d))
			m.dicts[l] = arr
		} else {
			arr = [][]rune{[]rune(d)}
			m.dicts[l] = arr
		}
	}
}

//Search Returns if there is any word in the trie that equals to the given word after modifying exactly one character
func (m *MagicDictionary) Search(word string) bool {
	l := len(word)
	arr, ok := m.dicts[l]
	if ok {
		for _, v := range arr {
			matchOne := 0
			words := []rune(word)
			for i := 0; i < l; i++ {
				if v[i] != words[i] {
					matchOne++
					if matchOne > 1 {
						break
					}
				}
			}
			if matchOne == 1 {
				return true
			}
		}
	}

	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */
