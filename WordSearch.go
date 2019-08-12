package main

/**
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.
*/

import (
	"fmt"
)

func main() {

	var word string
	fmt.Println("Insert Word:")
	fmt.Scanf("%s", &word)
	fmt.Println(wordExist(word))
}

var Letters = map[string]int{"A": 2, "B": 1, "C": 2, "D": 1, "E": 3, "S": 2, "F": 1}

func wordExist(word string) bool {
	for i := 0; i < len(word); i++ {
		board := string(word[i])
		if val, ok := Letters[board]; ok {
			Letters[board] = val - 1
			if Letters[board] == 0 {
				delete(Letters, board)
			}
		}else {
			return false
		}
	} 
	return true
}
