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
	"fmt"
	"html/template"
	"strings"
	"unicode"

	clStrings "github.com/go-corelibs/strings"
)

func ParseHtmlTagInlineKey(input string) (key string, ok bool) {
	// `^([a-zA-Z][-a-zA-Z0-9]+)$`
	if len(input) == 0 {
		return
	}
	var build []rune
	for idx, r := range input {
		if unicode.IsDigit(r) || r == '-' {
			if idx == 0 {
				return "", false
			}
		} else if !unicode.IsLetter(r) {
			return "", false
		}
		build = append(build, r)
	}
	return string(build), true
}

func ParseHtmlTagKeyValue(input string) (key, value string, ok bool) {
	// `^([a-zA-Z][-a-zA-Z0-9]+)=(.+?)$`
	if key, value, ok = strings.Cut(input, "="); ok {
		value = clStrings.TrimQuotes(value)
		key, ok = ParseHtmlTagInlineKey(key)
	}
	if !ok {
		key, value = "", ""
	}
	return
}

func parseAndUpdateHtmlTagAttributes(attributes map[string]interface{}, raw string) (e error) {
	for _, part := range strings.Split(raw, " ") {
		if part == "" {
			continue
		}
		var ok bool
		var key, quoted string
		if key, ok = ParseHtmlTagInlineKey(part); ok {
			attributes[key] = nil
		} else if key, quoted, ok = ParseHtmlTagKeyValue(part); ok {
			unquoted := clStrings.TrimQuotes(quoted)
			attributes[key] = unquoted
		} else {
			e = fmt.Errorf(`unsupported HTMLAttr format: %v`, part)
			return
		}
	}
	return
}

func ParseHtmlTagAttributes(input interface{}) (attributes map[string]interface{}, err error) {
	attributes = make(map[string]interface{})
	switch v := input.(type) {
	case string:
		err = parseAndUpdateHtmlTagAttributes(attributes, v)
	case template.HTML:
		err = parseAndUpdateHtmlTagAttributes(attributes, string(v))
	case template.HTMLAttr:
		err = parseAndUpdateHtmlTagAttributes(attributes, string(v))
	case []string:
		for _, tha := range v {
			if err = parseAndUpdateHtmlTagAttributes(attributes, tha); err != nil {
				return
			}
		}
	case []template.HTML:
		for _, tha := range v {
			if err = parseAndUpdateHtmlTagAttributes(attributes, string(tha)); err != nil {
				return
			}
		}
	case []template.HTMLAttr:
		for _, tha := range v {
			if err = parseAndUpdateHtmlTagAttributes(attributes, string(tha)); err != nil {
				return
			}
		}
	default:
		err = fmt.Errorf("unknown input type: (%T) %+v", v, v)
	}
	return
}
