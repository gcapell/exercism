package foodchain

import (
	"fmt"
	"strings"
)

var TestVersion = 1

var data = []struct{ name, line, extra string }{
	{},
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die.", ""},
	{"spider", "It wriggled and jiggled and tickled inside her.", " that wriggled and jiggled and tickled inside her"},
	{"bird", "How absurd to swallow a bird!", ""},
	{"cat", "Imagine that, to swallow a cat!", ""},
	{"dog", "What a hog, to swallow a dog!", ""},
	{"goat", "Just opened her throat and swallowed a goat!", ""},
	{"cow", "I don't know how she swallowed a cow!", ""},
}

func Verse(v int) string {
	if v == 8 {
		return "I know an old lady who swallowed a horse.\nShe's dead, of course!"
	}

	lines := []string{
		fmt.Sprintf("I know an old lady who swallowed a %s.", data[v].name),
		data[v].line,
	}

	for j := v; j > 1; j-- {
		lines = append(lines, fmt.Sprintf("She swallowed the %s to catch the %s%s.", data[j].name, data[j-1].name, data[j-1].extra))
	}
	if v != 1 {
		lines = append(lines, data[1].line)
	}
	return strings.Join(lines, "\n")
}

func Verses(min, max int) string {
	var verses []string
	for j := min; j <= max; j++ {
		verses = append(verses, Verse(j))
	}
	return strings.Join(verses, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}
