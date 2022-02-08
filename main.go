package main

import (
	"fmt"
	"sort"
	"strings"
)

type s struct {
	key string
	val int
}

func main() {
	// str := "a b c a b a f r s a c s s s s"
	str := "most computer software used to be written with its ultimate users clearly in mind. a computer installation had its own programmers who would write user-specific software. even if software from outside was used, it was frequently modified locally to suit the idiosyncrasies or particular applications of each installation. computers were expensive, programmers less so. but computers have become much cheaper, and are now common place at work and at home. there is a corresponding explosive growth in the amount of software available. computers and software are becoming mass-market merchandise. users frequently buy general-purpose software, expecting it to include enough functionality to cover their applications. the user no longer has a programming staff to fall back on. how then can he make a general-purpose system into a special-purpose one that suits his needs? the system might provide the mechanism for the user to do his task, but if it requires seventeen general-purpose operations to do what could take onespecial-purpose operation, the user will be frustrated and annoyed. one solution is to let the user do his own programming. certainly we should not expect the average user to write programs with the same easeas professional programmers. (see [anderson] and [cuff] for general discussions about the problems of introducing programming to the average person.) what we can do instead is to find ways to make programming easier, to couch it in terms with which the user is familiar, and to reduce the mental burdenof programming. programming by the user has already appeared in a number of different ways. the degree of user programmability can vary from choosing among options the system provides, to systems which require programming by the user as an integral part of their functionality. for user programming in its simplest form, systems provide options or modes from which the user can choose in order to customize the operation of the program. for instance, a program that prints might provide an option of single or double spacing. though it might be stretching a point to call this real programming, the system is supplying a very simple, specialized programming language."
	// str := "As you can expect, there is no fixed number of words that a paragraph should have. A rule of thumb: the paragraphs are usually about 100 to 200 words long, which is about 6-8 sentences. Nevertheless, it all depends on the ideas, and ideas come in many sizes. So the paragraph can be long enough to reach its end. But we can’t ignore the fact that readers don’t like seeing blocks of paragraphs. And it’s for that reason you should stick to one concept or idea per paragraph to avoid creating very long paragraphs."

	ch := make(chan map[string]int)
	go rangeMap(str, 20, ch)

	var p []s

	for key, value := range <-ch {
		p = append(p, s{
			key: key,
			val: value,
		})
	}

	sort.SliceStable(p, func(i, j int) bool {
		return p[i].val < p[j].val
	})

	fmt.Println(p)

	// tmp, _ := json.Marshal(result)
	// fmt.Println(string(tmp))
}

func rangeMap(s string, count int, ch chan map[string]int) {
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
	ch <- result
}
