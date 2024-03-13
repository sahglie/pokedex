package repo

import "fmt"

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
`, p.Name, p.Height, p.Weight)

	return s
}
