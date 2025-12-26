package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadLines(file *os.File) ([]string, error) {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func MustReadLines(file *os.File) []string {
	lines, err := ReadLines(file)
	if err != nil {
		log.Default().Fatalf("could not read file \"%s\": %s\n", file.Name(), err)
	}
	return lines
}

func ReadFile(fp string) ([]string, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadLines(f)
}

// Debugging

// DebugLogger wraps a logger with a Debug() print that only prints if its
// debug flag is set
type DebugLogger struct {
	*log.Logger
	isDebug bool
}

func (l *DebugLogger) Debug(v ...any) {
	if l.isDebug {
		l.Print(v...)
	}
}

func (l *DebugLogger) Debugf(format string, v ...any) {
	if l.isDebug {
		l.Printf(format, v...)
	}
}

func NewDebugLogger(out io.Writer, prefix string, flag int, isDebug bool) *DebugLogger {
	return &DebugLogger{
		Logger:  log.New(out, prefix, flag),
		isDebug: isDebug,
	}
}

// Slice utils - higher order slice functions

// Filter - return a slice with all items from a slice passing a predicate
func Filter[T any](s []T, f func(i T) bool) []T {
	res := []T{}
	for _, i := range s {
		if f(i) {
			res = append(res, i)
		}
	}
	return res
}
