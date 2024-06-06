package printer

import "github.com/jubenitezg/dupl/syntax"

type ReadFile func(filename string) ([]byte, error)

type Printer interface {
	PrintHeader() error
	PrintClones(dups [][]*syntax.Node) error
	PrintFooter() error
}
