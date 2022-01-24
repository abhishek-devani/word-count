package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

const serviceName = "Medium service"
const serviceDescription = "Simple service, just for fun"

func main() {
	str := "As you can expect, there is no fixed number of words that a paragraph should have. A rule of thumb: the paragraphs are usually about 100 to 200 words long, which is about 6-8 sentences. Nevertheless, it all depends on the ideas, and ideas come in many sizes. So the paragraph can be long enough to reach its end. But we can’t ignore the fact that readers don’t like seeing blocks of paragraphs. And it’s for that reason you should stick to one concept or idea per paragraph to avoid creating very long paragraphs."
	result := rangeMap(str, 10)
	tmp, _ := json.Marshal(result)
	fmt.Println(string(tmp))
}

func rangeMap(s string, count int) map[string]int {
	words := strings.Split(s, " ")

	// count occurences
	m := make(map[string]int)
	for _, word := range words {
		_, ok := m[word]
		if !ok {
			m[word] = 1
		} else {
			m[word]++
		}
	}
	// fmt.Println(m)

	// Takes max values
	// counts := make(map[string]int)
	// for key, value := range m {
	// 	if value > 0 {
	// 		counts[key] = value
	// 	}
	// }
	// fmt.Println(counts)

	// Sorts By Value
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return m[keys[i]] > m[keys[j]] })
	// fmt.Println(keys)

	// Builds result map
	result := make(map[string]int)
	for _, key := range keys {
		result[key] = m[key]
		count--
		if count == 0 {
			break
		}

	}
	return result
}
