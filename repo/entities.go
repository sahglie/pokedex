package repo

import (
	"fmt"
	"strings"
)

type Pokemon struct {
	Id     int
	Name   string
	Height int
	Weight int
	Stats  []string
	Types  []string
}

func (p *Pokemon) String() string {
	s := fmt.Sprintf(`
Name: %s
Height: %d
Weight: %d
Stats:
%v
Types:
%v
`, p.Name, p.Height, p.Weight, p.fmtStats(), p.fmtTypes())

	return s
}

func (p *Pokemon) fmtStats() string {
	b := strings.Builder{}
	for _, e := range p.Stats {
		b.WriteString("  - " + e + "\n")
	}
	return strings.TrimRight(b.String(), "\n")
}

func (p *Pokemon) fmtTypes() string {
	b := strings.Builder{}
	for _, e := range p.Types {
		b.WriteString("  - " + e + "\n")
	}
	return strings.TrimRight(b.String(), "\n")
}
