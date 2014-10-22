package ocr

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const digits = `
 _     _  _     _  _  _  _  _ 
| |  | _| _||_||_ |_   ||_||_|
|_|  ||_  _|  | _||_|  ||_| _|
                              `

var table = map[string]string{}

func init() {
	lines, err := splitLines(digits)
	if err != nil {
		log.Fatal(err)
	}
	for pos, s := range lines[0] {
		table[s] = strconv.Itoa(pos)
	}
}

// splitLines divides s into 4-row OCR lines, then into 3x3=9 char digits.
func splitLines(s string) ([][]string, error) {
	lines := strings.Split(s, "\n")
	if len(lines[0]) == 0 {
		lines = lines[1:]
	}
	if len(lines)%4 != 0 {
		return nil, fmt.Errorf("%d lines in %q", len(lines), lines)
	}
	var ocrLines [][]string
	for len(lines) > 0 {
		chars, err := splitChars(lines[:4])
		if err != nil {
			return nil, err
		}
		ocrLines = append(ocrLines, chars)
		lines = lines[4:]
	}
	return ocrLines, nil
}

func splitChars(lines []string) ([]string, error) {
	if len(lines) != 4 {
		return nil, fmt.Errorf("expected 4 lines, got %v", lines)
	}
	first, last := true, 0
	for pos, line := range lines {
		if first {
			last = len(line)
			if last%3 != 0 {
				return nil, fmt.Errorf("line %q length %d != 0 mod 3", line, last)
			}
			first = false
			continue
		}
		if len(line) != last {
			return nil, fmt.Errorf("line %d %q different length %d!=%d",
				pos, line, len(line), last)
		}
	}
	reply := []string{}
	for pos := 0; pos < len(lines[0]); pos += 3 {
		word := []string{}
		if lines[3][pos:pos+3] != "   " {
			reply = append(reply, "?????????")
			continue
		}
		for j := 0; j < 3; j++ {
			word = append(word, lines[j][pos:pos+3])
		}
		reply = append(reply, strings.Join(word, ""))
	}
	return reply, nil
}

func Recognize(s string) []string {
	lines, err := splitLines(s)
	if err != nil {
		log.Fatal(err)
	}
	reply := []string{}
	for _, line := range lines {
		word := ""
		for _, digit := range line {
			if ocr, ok := table[digit]; ok {
				word += ocr
			} else {
				word += "?"
			}
		}
		reply = append(reply, word)
	}
	return reply
}
