package leetcode

import "unicode"

type State int

const (
	LETTER State = iota
	HYPHEN
	WORD_WITH_HYPHEN
	SPACE
	PUNCTUATION
	BLANK
	INVALID
)

func countValidWords(sentence string) int {
	res := 0
	curState := BLANK
	for _, c := range sentence {
		if unicode.IsLetter(c) {
			curState = handleLetter(curState)
		} else if unicode.IsDigit(c) {
			curState = INVALID
		} else if c == '-' {
			curState = handleHyphen(curState)
		} else if unicode.IsPunct(c) {
			curState = handlePunctuation(curState)
		} else {
			if curState != INVALID && curState != SPACE && curState != BLANK && curState != HYPHEN {
				res++
			}
			curState = SPACE
		}
	}
	if curState != INVALID && curState != SPACE && curState != BLANK {
		res++
	}
	return res
}

func handleLetter(curState State) State {
	if curState == HYPHEN || curState == WORD_WITH_HYPHEN {
		return WORD_WITH_HYPHEN
	} else if curState == INVALID || curState == PUNCTUATION {
		return INVALID
	} else {
		return LETTER
	}
}

func handleHyphen(curState State) State {
	if curState != LETTER {
		return INVALID
	}
	return HYPHEN
}

func handlePunctuation(curState State) State {
	if curState != LETTER && curState != WORD_WITH_HYPHEN && curState != BLANK && curState != SPACE {
		return INVALID
	}
	return PUNCTUATION
}

func CountValidWords(sentence string) int {
	return countValidWords(sentence)
}
