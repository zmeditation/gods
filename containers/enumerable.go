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

// Enumerable functions for ordered containers.
// Ruby's enumerable inspired package.

package containers

type Enumerable interface {
	// Calls the given function once for each element, passing that element's index(key) and value.
	Each(func(index interface{}, value interface{}))

	// Invokes the given function once for each element and returns a
	// container containing the values returned by the given function.
	Map(func(index interface{}, value interface{}) interface{}) Container

	// Returns a new container containing all elements for which the given function returns a true value.
	Select(func(index interface{}, value interface{}) bool) Container

	// Passes each element of the collection to the given function and
	// returns true if the function ever returns true for any element.
	Any(func(index interface{}, value interface{}) bool) bool

	// Passes each element of the collection to the given function and
	// returns true if the function returns true for all elements.
	All(func(index interface{}, value interface{}) bool) bool

	// Passes each element of the collection to the given function and returns
	// the first for which the function is true or nil,nil otherwise if no element
	// matches the criteria.
	Find(func(index interface{}, value interface{}) bool) (index interface{}, value interface{})
}
