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
	"html/template"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHtml(t *testing.T) {
	Convey("ParseHtmlTagInlineKey", t, func() {
		parsed, ok := ParseHtmlTagInlineKey("")
		So(ok, ShouldBeFalse)
		So(parsed, ShouldEqual, "")
		parsed, ok = ParseHtmlTagInlineKey("yes")
		So(ok, ShouldBeTrue)
		So(parsed, ShouldEqual, "yes")
		parsed, ok = ParseHtmlTagInlineKey(" nope")
		So(ok, ShouldBeFalse)
		So(parsed, ShouldEqual, "")
		parsed, ok = ParseHtmlTagInlineKey("-nope")
		So(ok, ShouldBeFalse)
		So(parsed, ShouldEqual, "")
		parsed, ok = ParseHtmlTagInlineKey("0nope")
		So(ok, ShouldBeFalse)
		So(parsed, ShouldEqual, "")
		parsed, ok = ParseHtmlTagInlineKey("yes-yes")
		So(ok, ShouldBeTrue)
		So(parsed, ShouldEqual, "yes-yes")
	})

	Convey("ParseHtmlTagKeyValue", t, func() {
		key, value, ok := ParseHtmlTagKeyValue("")
		So(ok, ShouldBeFalse)
		So(key, ShouldEqual, "")
		So(value, ShouldEqual, "")
		key, value, ok = ParseHtmlTagKeyValue("nope")
		So(ok, ShouldBeFalse)
		So(key, ShouldEqual, "")
		So(value, ShouldEqual, "")
		key, value, ok = ParseHtmlTagKeyValue("=nope")
		So(ok, ShouldBeFalse)
		So(key, ShouldEqual, "")
		So(value, ShouldEqual, "")
		key, value, ok = ParseHtmlTagKeyValue("-=nope")
		So(ok, ShouldBeFalse)
		So(key, ShouldEqual, "")
		So(value, ShouldEqual, "")
		key, value, ok = ParseHtmlTagKeyValue("one=two")
		So(ok, ShouldBeTrue)
		So(key, ShouldEqual, "one")
		So(value, ShouldEqual, "two")
		key, value, ok = ParseHtmlTagKeyValue(`one="two"`)
		So(ok, ShouldBeTrue)
		So(key, ShouldEqual, "one")
		So(value, ShouldEqual, "two")
		key, value, ok = ParseHtmlTagKeyValue(`one='two'`)
		So(ok, ShouldBeTrue)
		So(key, ShouldEqual, "one")
		So(value, ShouldEqual, "two")
		key, value, ok = ParseHtmlTagKeyValue(`one = two`)
		So(ok, ShouldBeFalse)
		So(key, ShouldEqual, "")
		So(value, ShouldEqual, "")
	})

	Convey("ParseHtmlTagAttributes", t, func() {

		Convey("string input", func() {
			attrs, err := ParseHtmlTagAttributes("")
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{})
			attrs, err = ParseHtmlTagAttributes([]int{0})
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{})
			attrs, err = ParseHtmlTagAttributes("checked")
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{
				"checked": nil,
			})
			attrs, err = ParseHtmlTagAttributes("checked=true")
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{
				"checked": "true",
			})
			attrs, err = ParseHtmlTagAttributes("checked = nope")
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{"checked": nil})
			attrs, err = ParseHtmlTagAttributes("checked=true nope!")
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{"checked": "true"})
		})

		Convey("template.HTML input", func() {
			attrs, err := ParseHtmlTagAttributes(template.HTML(""))
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{})
			attrs, err = ParseHtmlTagAttributes([]int{0})
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{})
			attrs, err = ParseHtmlTagAttributes(template.HTML("checked"))
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{
				"checked": nil,
			})
			attrs, err = ParseHtmlTagAttributes(template.HTML("checked=true"))
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{
				"checked": "true",
			})
			attrs, err = ParseHtmlTagAttributes(template.HTML("checked = nope"))
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{"checked": nil})
			attrs, err = ParseHtmlTagAttributes(template.HTML("checked=true nope!"))
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{"checked": "true"})
		})

		Convey("template.HTMLAttr input", func() {
			attrs, err := ParseHtmlTagAttributes(template.HTMLAttr(""))
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{})
			attrs, err = ParseHtmlTagAttributes([]int{0})
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{})
			attrs, err = ParseHtmlTagAttributes(template.HTMLAttr("checked"))
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{
				"checked": nil,
			})
			attrs, err = ParseHtmlTagAttributes(template.HTMLAttr("checked=true"))
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{
				"checked": "true",
			})
			attrs, err = ParseHtmlTagAttributes(template.HTMLAttr("checked = nope"))
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{"checked": nil})
			attrs, err = ParseHtmlTagAttributes(template.HTMLAttr("checked=true nope!"))
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{"checked": "true"})
		})

		Convey("[]string input", func() {
			attrs, err := ParseHtmlTagAttributes([]string{""})
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{})
			attrs, err = ParseHtmlTagAttributes([]int{0})
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{})
			attrs, err = ParseHtmlTagAttributes([]string{"checked"})
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{
				"checked": nil,
			})
			attrs, err = ParseHtmlTagAttributes([]string{"checked=true"})
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{
				"checked": "true",
			})
			attrs, err = ParseHtmlTagAttributes([]string{"checked = nope"})
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{"checked": nil})
			attrs, err = ParseHtmlTagAttributes([]string{"checked=true nope!"})
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{"checked": "true"})
		})

		Convey("[]template.HTML input", func() {
			attrs, err := ParseHtmlTagAttributes([]template.HTML{""})
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{})
			attrs, err = ParseHtmlTagAttributes([]int{0})
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{})
			attrs, err = ParseHtmlTagAttributes([]template.HTML{"checked"})
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{
				"checked": nil,
			})
			attrs, err = ParseHtmlTagAttributes([]template.HTML{"checked=true"})
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{
				"checked": "true",
			})
			attrs, err = ParseHtmlTagAttributes([]template.HTML{"checked = nope"})
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{"checked": nil})
			attrs, err = ParseHtmlTagAttributes([]template.HTML{"checked=true nope!"})
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{"checked": "true"})
		})

		Convey("[]template.HTMLAttr input", func() {
			attrs, err := ParseHtmlTagAttributes([]template.HTMLAttr{""})
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{})
			attrs, err = ParseHtmlTagAttributes([]int{0})
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{})
			attrs, err = ParseHtmlTagAttributes([]template.HTMLAttr{"checked"})
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{
				"checked": nil,
			})
			attrs, err = ParseHtmlTagAttributes([]template.HTMLAttr{"checked=true"})
			So(err, ShouldBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{
				"checked": "true",
			})
			attrs, err = ParseHtmlTagAttributes([]template.HTMLAttr{"checked = nope"})
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{"checked": nil})
			attrs, err = ParseHtmlTagAttributes([]template.HTMLAttr{"checked=true nope!"})
			So(err, ShouldNotBeNil)
			So(attrs, ShouldEqual, map[string]interface{}{"checked": "true"})
		})

	})
}
