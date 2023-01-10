/*
Suppose the content library contains the following titles: "duel", "dule", "speed", "spede", "deul", "cars".
How would you efficiently implement a functionality so that if a user misspell speed as spede,
they are shown the correct title?

We want to split the list of titles into sets of words so that all words in a set are anagrams.
In the above list, there are three sets: {"duel", "dule", "deul"}, {"speed", "spede"}, and {"cars"}.
Search results should comprise all members of the set that the search string is found in.
We should pre-compute these sets instead of forming them when the user searches a title.
*/

package main

import (
	"fmt"
	"strconv"
)

// this function is to get count vector of the word ex: abbccd :: 122100000000000000000000
func getVector(query string) string {
	key := ""
	countVector := make([]int, 26)
	for _, ele := range query {
		index := ele - 'a'
		countVector[index]++
	}
	for i := range countVector {
		key += strconv.Itoa(countVector[i])
	}
	return key
}

// this func is to compute map of anagaram of the word with actual word ex: speed -> {speed, spede, eepds}, duel -> {duel,dule}
func computeGroups(titles []string) map[string][]string {

	res := make(map[string][]string) //make(map[key_type]value_type, value_type is []string)

	for _, t := range titles {
		key := getVector(t)
		res[key] = append(res[key], t)
	}
	return res
}

func main() {
	titles := []string{"duel", "dule", "speed", "spede", "duel", "cars"}
	output := computeGroups(titles)

	query := "sca"
	queryVector := getVector(query)

	if v, found := output[queryVector]; found {
		fmt.Println("the elements are:: ")
		fmt.Println(v)
	} else {
		fmt.Println("no matching element!!")
	}
}
