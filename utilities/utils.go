package utilities

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type Stack[T any] []T

// IsEmpty checks if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack[T]) Push(element T) {
	*s = append(*s, element) // Simply append the new value to the end of the stack
}

// Pop Removes and return top element of stack. Return false if stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	var result T
	if s.IsEmpty() {
		return result, false
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		result = (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]    // Remove it from the stack by slicing it off.
		return result, true
	}
}

func ReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	Check(err)
	return data
}

func ScannerFromFile(path string) *bufio.Scanner {
	data, err := os.ReadFile(path)
	Check(err)
	s := string(data)
	scanner := bufio.NewScanner(strings.NewReader(s))
	return scanner
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
