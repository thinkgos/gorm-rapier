package rapier

import (
	"strings"

	"gorm.io/gorm/clause"
)

type WhenThenExpr Expr

func WhenThen(condition Expr, result Expr) WhenThenExpr {
	return expr{
		e: clause.Expr{
			SQL:                "WHEN ? THEN ?",
			Vars:               []any{condition, result},
			WithoutParentheses: false,
		},
	}
}

/*
CASE

	WHEN condition1 THEN result1
	WHEN condition2 THEN result2
	...

END
*/
func CaseWhen(wts ...WhenThenExpr) Field {
	sqlBuilder := strings.Builder{}
	sqlBuilder.Grow(5 + 2*len(wts) + 6)
	vars := make([]any, 0, len(wts))
	sqlBuilder.WriteString("(CASE")
	for _, wt := range wts {
		sqlBuilder.WriteString(" ?")
		vars = append(vars, wt)
	}
	sqlBuilder.WriteString(" END)")
	return Field{
		expr{
			e: clause.Expr{
				SQL:                sqlBuilder.String(),
				Vars:               vars,
				WithoutParentheses: false,
			},
		},
	}
}

/*
CASE

	WHEN condition1 THEN result1
	WHEN condition2 THEN result2
	...
	ELSE result
END
*/

func CaseWhenElse(elseResult Expr, wts ...WhenThenExpr) Field {
	sqlBuilder := strings.Builder{}
	sqlBuilder.Grow(5 + 2*len(wts) + 12)
	vars := make([]any, 0, len(wts)+1)
	sqlBuilder.WriteString("(CASE")
	for _, ee := range wts {
		sqlBuilder.WriteString(" ?")
		vars = append(vars, ee)
	}
	vars = append(vars, elseResult)
	sqlBuilder.WriteString(" ELSE ? END)")
	return Field{
		expr{
			e: clause.Expr{
				SQL:                sqlBuilder.String(),
				Vars:               vars,
				WithoutParentheses: false,
			},
		},
	}
}
