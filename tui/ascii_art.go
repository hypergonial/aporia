package tui

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
A struct that makes loading and printing ascii art easy.
*/
type asciiArt struct {
	strLines []string
	lines    int
	cols     int
}

func newAscii(art string) asciiArt {
	lines := strings.Split(art, "\n")

	longestLine := utf8.RuneCountInString(lines[0])

	for _, line := range lines[1:] {
		if len(line) > longestLine {
			longestLine = utf8.RuneCountInString(line)
		}
	}

	return asciiArt{
		strLines: lines,
		cols:     longestLine,
		lines:    len(lines),
	}
}

func (self *asciiArt) draw(termSize TermSize) {
	linesSkip := maxInt((self.lines-termSize.Lines)/2, 0)
	colsSkip := maxInt((self.cols-termSize.Cols)/2, 0)

	startLine := maxInt((termSize.Lines-self.lines)/2, 0)
	startCol := maxInt((termSize.Cols-self.cols)/2, 0)

	for i := 0; i < startLine; i++ {
		fmt.Print("\n\r")
	}

	for i := 0; i < termSize.Lines && i+linesSkip < self.lines; i++ {
		maxSize := minInt(termSize.Cols + colsSkip, len(self.strLines[i+linesSkip]))
		fmt.Print(strings.Repeat(" ", startCol), self.strLines[i+linesSkip][colsSkip:maxSize], "\n\r")
	}
}
