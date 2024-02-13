// Copyright (c) 2024  The Go-CoreLibs Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package htmlcss

import (
	"strings"
	"sync"

	"github.com/go-corelibs/slices"
)

type CssClass interface {
	// Has returns true if the given class name is present within this
	// CssClass	instance
	Has(name string) (present bool)
	// Add appends the given class name to this CssClass instance
	Add(name string)
	// Rem removes the given class name from this CssClass instance
	Rem(name string)
	// List returns a list of the class names within this CssClass instance,
	// in the order they were added
	List() (list []string)
	// Apply takes an HTML tag's class attribute value and adds all classes
	// to this CssClass instance
	Apply(class string)
	// String returns the HTML class attribute value for this CssClass
	// instance
	String() (class string)
}

type cssClass struct {
	src    string
	order  []string
	lookup map[string]struct{}

	m *sync.RWMutex
}

func ParseClass(class string) (c CssClass) {
	c = &cssClass{
		lookup: make(map[string]struct{}),
		m:      &sync.RWMutex{},
	}
	c.Apply(class)
	return
}

func (c *cssClass) Has(name string) (present bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	_, present = c.lookup[name]
	return
}

func (c *cssClass) Add(name string) {
	if name == "" {
		return
	}
	c.m.Lock()
	defer c.m.Unlock()
	name = strings.TrimSpace(name)
	if _, present := c.lookup[name]; !present {
		c.order = append(c.order, name)
		c.lookup[name] = struct{}{}
	}
	return
}

func (c *cssClass) Rem(name string) {
	c.m.Lock()
	defer c.m.Unlock()
	c.order = slices.Prune(c.order, name)
	delete(c.lookup, name)
	return
}

func (c *cssClass) List() (list []string) {
	c.m.RLock()
	defer c.m.RUnlock()
	list = c.order[:]
	return
}

func (c *cssClass) Apply(class string) {
	class = strings.TrimSpace(class)
	for _, name := range strings.Split(class, " ") {
		if !c.Has(name) {
			c.Add(name)
		}
	}
}

func (c *cssClass) String() (class string) {
	c.m.RLock()
	defer c.m.RUnlock()
	class = strings.Join(c.order, " ")
	return
}
