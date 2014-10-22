package house

import (
	"fmt"
	"strings"
)

func Embed(relPhrase, nounPhrase string) string {
	return relPhrase + " " + nounPhrase
}

func Verse(subject string, relPhrases []string, nounPhrase string) string {
	all := []string{subject}
	all = append(all, relPhrases...)
	all = append(all, nounPhrase)
	return strings.Join(all, " ")
}

var data = []struct{ object, join string }{
	{"house that Jack built.", ""},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

func verse(j int) string {
	lines := []string{fmt.Sprintf("This is the %s", data[j].object)}

	for k := j; k > 0; k-- {
		lines = append(lines, fmt.Sprintf("that %s the %s", data[k].join, data[k-1].object))
	}
	return strings.Join(lines, "\n")
}

func Song() string {
	var verses []string
	for j := range data {
		verses = append(verses, verse(j))
	}
	return strings.Join(verses, "\n\n")
}
