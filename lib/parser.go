package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ParserAction int

const (
	MoveableAction = iota
	WorldAction
	EnemyAction
	PlayerAction
	BGMusicAction
)

// LineParser represents a parser that reads lines from a file
type LineParser struct {
	filePath string
	file     *os.File
	scanner  *bufio.Scanner
}

// NewLineParser creates a new LineParser for the specified file path
func NewLineParser(filePath string) (*LineParser, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	scanner := bufio.NewScanner(file)
	return &LineParser{
		filePath: filePath,
		file:     file,
		scanner:  scanner,
	}, nil
}

// Next returns the next line from the file and a boolean indicating whether there are more lines
func (p *LineParser) Next() (string, bool) {
	if p.scanner.Scan() {
		return p.scanner.Text(), true
	}
	// Check for errors during scanning
	if err := p.scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
	}
	return "", false
}

// Close closes the file
func (p *LineParser) Close() error {
	return p.file.Close()
}

func (p *LineParser) SplitLine(line string) []string {
	return strings.Fields(line) // This handles multiple spaces, tabs, etc.
}
