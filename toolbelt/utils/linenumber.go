package utils

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type LineNumberFinder struct {
	data []byte
}

func NewLineNumberFinder(data []byte) *LineNumberFinder {
	return &LineNumberFinder{data: data}
}

func (l *LineNumberFinder) FindLineNumber(s string) int {
	r := bytes.NewReader(l.data)
	r.Seek(0, io.SeekStart)

	scanner := bufio.NewScanner(r)

	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), s) {
			return line
		}

		line++
	}

	return 0
}
