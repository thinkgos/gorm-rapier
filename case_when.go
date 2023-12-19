package rapier

import (
	"strings"

	"gorm.io/gorm/clause"
)

type CaseWhen struct {
	wts    []*clause.Expr
	result Expr
}

// NewCaseWhen new case when clause
// CASE
//
//	WHEN condition1 THEN result1
//	WHEN condition2 THEN result2
//	...
//	[ELSE result]
//
// END
func NewCaseWhen() *CaseWhen {
	return &CaseWhen{
		wts: make([]*clause.Expr, 0, 16),
	}
}

// WhenThen add `WHEN condition THEN result`
func (c *CaseWhen) WhenThen(condition Expr, result Expr) *CaseWhen {
	c.wts = append(c.wts, &clause.Expr{
		SQL:                "WHEN ? THEN ?",
		Vars:               []any{condition, result},
		WithoutParentheses: false,
	})
	return c
}

// Else add `ELSE result“
func (c *CaseWhen) Else(result Expr) *CaseWhen {
	c.result = result
	return c
}

func (c *CaseWhen) Build() Field {
	b := &strings.Builder{}
	b.Grow(5 + 2*len(c.wts) + 12)
	vars := make([]any, 0, len(c.wts)+1)
	b.WriteString("(CASE")
	for _, wt := range c.wts {
		b.WriteString(" ?")
		vars = append(vars, wt)
	}
	if c.result != nil {
		vars = append(vars, c.result)
		b.WriteString(" ELSE ?")
	}
	b.WriteString(" END)")
	return Field{
		expr{
			e: clause.Expr{
				SQL:                b.String(),
				Vars:               vars,
				WithoutParentheses: false,
			},
		},
	}
}
