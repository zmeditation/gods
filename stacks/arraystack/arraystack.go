/*
Copyright (c) 2015, Emir Pasic
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// Implementation of stack using a slice.
// Structure is not thread safe.
// References: http://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29

package arraystack

import (
	"fmt"
	"github.com/emirpasic/gods/stacks"
	"strings"
)

func assertInterfaceImplementation() {
	var _ stacks.Interface = (*Stack)(nil)
}

type Stack struct {
	elements []interface{}
	top      int
}

// Instantiates a new empty stack
func New() *Stack {
	return &Stack{top: -1}
}

// Pushes a value onto the top of the stack
func (stack *Stack) Push(value interface{}) {
	// Increase when capacity is reached by a factor of 1.5 and add one so it grows when size is zero
	if stack.top+1 >= cap(stack.elements) {
		currentSize := len(stack.elements)
		sizeIncrease := int(1.5*float32(currentSize) + 1.0)
		newSize := currentSize + sizeIncrease
		newItems := make([]interface{}, newSize, newSize)
		copy(newItems, stack.elements)
		stack.elements = newItems
	}
	stack.top += 1
	stack.elements[stack.top] = value
}

// Pops (removes) top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (stack *Stack) Pop() (value interface{}, ok bool) {
	if stack.top >= 0 {
		value, ok = stack.elements[stack.top], true
		stack.top -= 1
		return
	}
	return nil, false
}

// Returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (stack *Stack) Peek() (value interface{}, ok bool) {
	if stack.top >= 0 {
		return stack.elements[stack.top], true
	}
	return nil, false
}

// Returns true if stack does not contain any elements.
func (stack *Stack) Empty() bool {
	return stack.Size() == 0
}

// Returns number of elements within the stack.
func (stack *Stack) Size() int {
	return stack.top + 1
}

// Removes all elements from the stack.
func (stack *Stack) Clear() {
	stack.top = -1
	stack.elements = []interface{}{}
}

func (stack *Stack) String() string {
	str := "ArrayStack\n"
	values := []string{}
	for _, value := range stack.elements {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}
