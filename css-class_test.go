// Copyright (c) 2024  The Go-Enjin Authors
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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCssClass(t *testing.T) {
	Convey("ParseClass", t, func() {
		c := ParseClass("one two")
		So(c, ShouldNotBeNil)
		So(c.String(), ShouldEqual, "one two")
	})

	Convey("Has", t, func() {
		c := ParseClass("one")
		So(c, ShouldNotBeNil)
		So(c.Has("one"), ShouldBeTrue)
		So(c.Has("two"), ShouldBeFalse)
	})

	Convey("Add", t, func() {
		c := ParseClass("one")
		So(c, ShouldNotBeNil)
		So(c.String(), ShouldEqual, "one")
		c.Add("one")
		So(c.String(), ShouldEqual, "one")
		c.Add("two")
		So(c.String(), ShouldEqual, "one two")
		c.Add("")
		So(c.String(), ShouldEqual, "one two")
	})

	Convey("Rem", t, func() {
		c := ParseClass("one two")
		So(c, ShouldNotBeNil)
		c.Rem("more")
		So(c.String(), ShouldEqual, "one two")
		c.Rem("two")
		So(c.String(), ShouldEqual, "one")
	})

	Convey("List", t, func() {
		c := ParseClass("one two")
		So(c, ShouldNotBeNil)
		So(c.List(), ShouldEqual, []string{"one", "two"})
	})

	Convey("Apply", t, func() {
		c := ParseClass("one two")
		So(c, ShouldNotBeNil)
		c.Apply("two more")
		So(c.String(), ShouldEqual, "one two more")
	})
}
