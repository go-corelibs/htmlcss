[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/htmlcss)
[![codecov](https://codecov.io/gh/go-corelibs/htmlcss/graph/badge.svg?token=2xdELREztB)](https://codecov.io/gh/go-corelibs/htmlcss)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/htmlcss)](https://goreportcard.com/report/github.com/go-corelibs/htmlcss)

# htmlcss - HTML/CSS related utilities

A collection of utilities for working with HTML and CSS related things.

# Installation

``` shell
> go get github.com/go-corelibs/htmlcss@latest
```

# Examples

## AddClassNames

``` go
func main() {
    updated := htmlcss.AddClassNamed("original classes", "one", "two", "classes")
    // updated == "original classes one two"
}
```

## ParseClass

``` go
func main() {
    classes := ParseClass("one two")
    original := classes.String()
    // original == "one two"
    classes.Add("more")
    classes.Rem("two")
    modified := classes.String()
    // modified == "one more"
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2024 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
