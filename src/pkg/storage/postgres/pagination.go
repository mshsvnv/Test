package postgres

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
)

type SortDirection int8

const (
	ASC SortDirection = iota
	DESC
)

func (s SortDirection) String() string {
	switch s {
	case DESC:
		return "DESC"
	default:
		return "ASC"
	}
}

func SortDirectionFromString(dir string) SortDirection {
	switch dir {
	case "ASC":
		return ASC
	default:
		return DESC
	}
}

type SortOptions struct {
	Direction SortDirection
	Columns   []string
}

func (s SortOptions) Format() string {
	return fmt.Sprintf("%s %s", strings.Join(s.Columns, ","), s.Direction.String())
}

type FilterOptions struct {
	Column  string
}

type Pagination struct {
	Filter     FilterOptions
	Sort       SortOptions
}

func (p *Pagination) ToSQL(s squirrel.SelectBuilder) squirrel.SelectBuilder {
	s = s.OrderBy(p.Sort.Format())

	return s
}