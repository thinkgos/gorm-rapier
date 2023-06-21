package assist

import (
	"time"

	"gorm.io/gorm/clause"
)

func (e expr) ifNull(value any) expr {
	e.e = clause.Expr{
		SQL:  "IFNULL(?,?)",
		Vars: []any{e.RawExpr(), value},
	}
	return e
}

func (e expr) inAny(value any) expr {
	e.e = intoInExpr(e.RawExpr(), value)
	return e
}

func (e expr) notInAny(value any) expr {
	e.e = clause.Not(intoInExpr(e.RawExpr(), value))
	return e
}

func (e expr) regexp(value any) expr {
	e.e = clause.Expr{
		SQL:  "? REGEXP ?",
		Vars: []any{e.RawExpr(), value},
	}
	return e
}

func (e expr) notRegexp(value any) expr {
	e.e = clause.Not(clause.Expr{
		SQL:  "? REGEXP ?",
		Vars: []any{e.RawExpr(), value},
	})
	return e
}

func (e expr) between(values []any) expr {
	e.e = clause.Expr{
		SQL:  "? BETWEEN ? AND ?",
		Vars: append([]any{e.RawExpr()}, values...),
	}
	return e
}

func (e expr) notBetween(values []any) expr {
	e.e = clause.Not(clause.Expr{
		SQL:  "? BETWEEN ? AND ?",
		Vars: append([]any{e.RawExpr()}, values...),
	})
	return e
}

func (e expr) sum() expr {
	e.e = clause.Expr{
		SQL:  "SUM(?)",
		Vars: []any{e.RawExpr()},
	}
	return e
}

func (e expr) add(value any) expr {
	switch v := value.(type) {
	case time.Duration:
		e.e = clause.Expr{
			SQL:  "DATE_ADD(?, INTERVAL ? MICROSECOND)",
			Vars: []any{e.RawExpr(), v.Microseconds()},
		}
	default:
		e.e = clause.Expr{
			SQL:  "?+?",
			Vars: []any{e.RawExpr(), value},
		}
	}
	return e
}

func (e expr) sub(value any) expr {
	switch v := value.(type) {
	case time.Duration:
		e.e = clause.Expr{
			SQL:  "DATE_SUB(?, INTERVAL ? MICROSECOND)",
			Vars: []any{e.RawExpr(), v.Microseconds()},
		}
	default:
		e.e = clause.Expr{
			SQL:  "?-?",
			Vars: []any{e.RawExpr(), value},
		}
	}
	return e
}

func (e expr) mul(value any) expr {
	if e.e == nil {
		e.e = clause.Expr{
			SQL:  "?*?",
			Vars: []any{e.col, value},
		}
	} else {
		e.e = clause.Expr{
			SQL:  "(?)*?",
			Vars: []any{e.e, value},
		}
	}
	return e
}

func (e expr) div(value any) expr {
	if e.e == nil {
		e.e = clause.Expr{
			SQL:  "?/?",
			Vars: []any{e.col, value},
		}
	} else {
		e.e = clause.Expr{
			SQL:  "(?)/?",
			Vars: []any{e.e, value},
		}
	}
	return e
}

func (e expr) mod(value any) expr {
	if e.e == nil {
		e.e = clause.Expr{
			SQL:  "?%?",
			Vars: []any{e.col, value},
		}
	} else {
		e.e = clause.Expr{
			SQL:  "(?)%?",
			Vars: []any{e.e, value},
		}
	}
	return e
}

func (e expr) floorDiv(value any) expr {
	if e.e == nil {
		e.e = clause.Expr{
			SQL:  "? DIV ?",
			Vars: []any{e.col, value},
		}
	} else {
		e.e = clause.Expr{
			SQL:  "(?) DIV ?",
			Vars: []any{e.e, value},
		}
	}
	return e
}

func (e expr) floor() expr {
	e.e = clause.Expr{
		SQL:  "FLOOR(?)",
		Vars: []any{e.RawExpr()},
	}
	return e
}

func (e expr) round(decimals int) expr {
	e.e = clause.Expr{
		SQL:  "ROUND(?, ?)",
		Vars: []any{e.RawExpr(), decimals},
	}
	return e
}

// findInSet equal to FIND_IN_SET(expr, targetList)
func (e expr) findInSet(targetList string) expr {
	e.e = clause.Expr{
		SQL:  "FIND_IN_SET(?, ?)",
		Vars: []any{e.RawExpr(), targetList},
	}
	return e
}

// findInSetWith equal to FIND_IN_SET(target, expr)
func (e expr) findInSetWith(target string) expr {
	e.e = clause.Expr{
		SQL:  "FIND_IN_SET(?, ?)",
		Vars: []any{target, e.RawExpr()},
	}
	return e
}

func (e expr) rightShift(value any) expr {
	if e.e == nil {
		e.e = clause.Expr{
			SQL:  "?>>?",
			Vars: []any{e.col, value},
		}
	} else {
		e.e = clause.Expr{
			SQL:  "(?)>>?",
			Vars: []any{e.e, value},
		}
	}
	return e
}

func (e expr) leftShift(value any) expr {
	if e.e == nil {
		e.e = clause.Expr{
			SQL:  "?<<?",
			Vars: []any{e.col, value},
		}
	} else {
		e.e = clause.Expr{
			SQL:  "(?)<<?",
			Vars: []any{e.e, value},
		}
	}
	return e
}

func (e expr) bitXor(value any) expr {
	if e.e == nil {
		e.e = clause.Expr{
			SQL:  "?^?",
			Vars: []any{e.col, value},
		}
	} else {
		e.e = clause.Expr{
			SQL:  "(?)^?",
			Vars: []any{e.e, value},
		}
	}
	return e
}

func (e expr) bitAnd(value any) expr {
	if e.e == nil {
		e.e = clause.Expr{
			SQL:  "?&?",
			Vars: []any{e.col, value},
		}
	} else {
		e.e = clause.Expr{
			SQL:  "(?)&?",
			Vars: []any{e.e, value},
		}
	}
	return e
}

func (e expr) bitOr(value any) expr {
	if e.e == nil {
		e.e = clause.Expr{
			SQL:  "?|?",
			Vars: []any{e.col, value},
		}
	} else {
		e.e = clause.Expr{
			SQL:  "(?)|?",
			Vars: []any{e.e, value},
		}
	}
	return e
}

func (e expr) bitFlip() expr {
	if e.e == nil {
		e.e = clause.Expr{
			SQL:  "~?",
			Vars: []any{e.col},
		}
	} else {
		e.e = clause.Expr{
			SQL:  "~(?)",
			Vars: []any{e.RawExpr()},
		}
	}
	return e
}

func (e expr) and(value any) expr {
	e.e = clause.Expr{
		SQL:  "? AND ?",
		Vars: []any{e.RawExpr(), value},
	}
	return e
}

func (e expr) or(value any) expr {
	e.e = clause.Expr{
		SQL:  "? OR ?",
		Vars: []any{e.RawExpr(), value},
	}
	return e
}

func (e expr) xor(value any) expr {
	e.e = clause.Expr{
		SQL:  "? XOR ?",
		Vars: []any{e.RawExpr(), value},
	}
	return e
}
