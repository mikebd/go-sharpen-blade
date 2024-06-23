package ngram

import (
	"strings"
)

/*
Question:

An n-gram, which is used in linguistics, refers to a contiguous sequence of N
items from a given sequence of text.

Example:
Suppose we have the following sentences.
"the cow jumped over the moon"
"the cow and the moon"

2-grams would be:
"the cow"
"cow jumped"
"jumped over"
"over the"
"the moon"
"cow and"
"and the"

Given a list of reviews, we want the frequency of all N-grams, where N is a
parameter. If we use the following reviews:
"the cow jumped over the moon"
"the cow and the moon"
And we use N=2, then the expected result is:
"the cow" ⇒ 2
"cow jumped" ⇒ 1
"jumped over" ⇒ 1
"over the" ⇒ 1
"the moon" ⇒ 2
"cow and" ⇒ 1
"and the" ⇒ 1
*/

type NgramResult map[string]int

const TwoGramSize = 2

func TwoGram(reviews []string) NgramResult {
	result := make(NgramResult)
	for _, review := range reviews {
		words := strings.Split(review, " ")
		for index := range words {
			ngram := strings.Join(words[index:index+TwoGramSize], " ")
			count := result[ngram]
			result[ngram] = count + 1
			if index+TwoGramSize == len(words) {
				break
			}
		}
	}
	return result
}
