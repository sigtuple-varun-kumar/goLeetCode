package leetcode

import (
	"fmt"
)

type Node struct {
	val  interface{}
	next *Node
}

type Stack struct {
	top *Node
	len int
}

func (s *Stack) Push(value interface{}) {
	newNode := Node{
		val:  value,
		next: s.top,
	}
	s.top = &newNode
	s.len++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.len == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	value := s.top.val
	s.top = s.top.next
	s.len--
	return value, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	return s.top.val, nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {
	return s.len
}

type parenthesis struct {
	Type     rune
	Location int
}

func minRemoveToMakeValid(s string) string {
	parethesisStack := Stack{}
	for i, char := range s {
		if char == '(' {
			parethesisStack.Push(parenthesis{
				Type:     '(',
				Location: i,
			})
		}
		if char == ')' {
			if parethesisStack.IsEmpty() {
				parethesisStack.Push(parenthesis{
					Type:     ')',
					Location: i,
				})
			}
			// we have two paths if last one was ( then we remove that from stack otherwise we push this one
			lastParenthesis, _ := parethesisStack.Peek()
			// we have handles empty error above
			if lastParenthesis.(parenthesis).Type == '(' {
				parethesisStack.Pop()
			} else {
				parethesisStack.Push(parenthesis{
					Type:     ')',
					Location: i,
				})
			}
		}
	}
	c := []rune(s)
	for !parethesisStack.IsEmpty() {
		val, err := parethesisStack.Pop()
		if err != nil {
			break
		}
		indexToRemove := val.(parenthesis).Location
		// Convert string to runes for manipulation
		c[indexToRemove] = 0
	}
	ans := make([]rune, 0, len(c))
	for _, char := range c {
		if char != 0 {
			ans = append(ans, char)
		}
	}
	return string(ans)
}
