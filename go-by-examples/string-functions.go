package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains ", s.Contains("test", "es"))
	p("Count ", s.Count("test", "t"))
	p("Count ", s.HasPrefix("test", "te"))
	p("Count ", s.HasSuffix("test", "st"))
	p("Index", s.Index("test", "s"))
	p("Join", s.Join([]string{"a", "b"}, "-"))
	p("Repeat", s.Repeat("a", 5))
	p("Replace", s.Replace("Daavid", "a", "i", -1))
	p("Replace", s.Replace("Daavid", "a", "i", 1))
	p("Split", s.Split("a-b-c-d-e", "-"))
	p("Lower", s.ToLower("TEST"))
	p("Upper", s.ToUpper("test"))

	p("Len", len("Hello"))
	p("Char", "hello"[0])
}
