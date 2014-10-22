		// return tab[r]
		// return tab[r]
package anagram

import (
    "sort"
    "strings"
)

type runeSlice []rune

func (r runeSlice) Len() int           { return len(r) }
func (r runeSlice) Less(i, j int) bool { return r[i] < r[j] }
func (r runeSlice) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func sameSlice(rs, rt runeSlice) bool {
     for i := range rs {
        if rs[i] != rt[i] {
            return false
        }
    }
    return true
}

func Detect(subject string, candidates []string) []string {
    anagrams := make([]string, 0, len(candidates))
    subject = strings.ToLower(subject)
    rs := runeSlice(subject)
    sort.Sort(rs)
    for _, c := range candidates {
    	if len(c) != len(subject) { 
    		continue
    	}
        c = strings.ToLower(c)
        if c == subject {
        		continue
        	}
        cs := runeSlice(c)
        sort.Sort(cs)
        if sameSlice(rs, cs) {
            anagrams = append(anagrams, c)
        }
    }
    return anagrams
}
