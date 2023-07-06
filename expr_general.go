package assist

import (
	"time"

	"gorm.io/gorm/clause"
)

// IFNULL(expr,?)
func (e expr) innerIfNull(value any) expr {
	e.e = clause.Expr{
		SQL:  "IFNULL(?,?)",
		Vars: []any{e.RawExpr(), value},
	}
	return e
}

// expr = ?
func (e expr) innerEq(value any) expr {
	e.e = clause.Eq{Column: e.RawExpr(), Value: value}
	return e
}

// expr <> ?
func (e expr) innerNeq(value any) expr {
	e.e = clause.Neq{Column: e.RawExpr(), Value: value}
	return e
}

// expr > ?
func (e expr) innerGt(value any) expr {
	e.e = clause.Gt{Column: e.RawExpr(), Value: value}
	return e
}

// use expr >= ?
func (e expr) innerGte(value any) expr {
	e.e = clause.Gte{Column: e.RawExpr(), Value: value}
	return e
}

// expr < ?
func (e expr) innerLt(value any) expr {
	e.e = clause.Lt{Column: e.RawExpr(), Value: value}
	return e
}

// expr <= ?
func (e expr) innerLte(value any) expr {
	e.e = clause.Lte{Column: e.RawExpr(), Value: value}
	return e
}

// expr IN(?,?...)
func (e expr) innerIn(values []any) expr {
	e.e = clause.IN{Column: e.RawExpr(), Values: values}
	return e
}

// expr IN(?,?...)
func (e expr) innerInAny(value any) expr {
	e.e = intoClauseIN(e.RawExpr(), value)
	return e
}

// expr IN(?,?...)
func (e expr) innerNotIn(values []any) expr {
	e.e = clause.NotConditions{
		Exprs: []clause.Expression{
			clause.IN{Column: e.RawExpr(), Values: values},
		},
	}
	return e
}

// expr IN(?,?...)
func (e expr) innerNotInAny(value any) expr {
	e.e = clause.NotConditions{
		Exprs: []clause.Expression{
			intoClauseIN(e.RawExpr(), value),
		},
	}
	return e
}

// expr LIKE ?
func (e expr) innerLike(value any) expr {
	e.e = clause.Like{
		Column: e.RawExpr(),
		Value:  value,
	}
	return e
}

// expr NOT LIKE ?
func (e expr) innerNotLike(value any) expr {
	e.e = clause.NotConditions{
		Exprs: []clause.Expression{clause.Like{
			Column: e.RawExpr(),
			Value:  value,
		},
		},
	}
	return e
}

// expr REGEXP ?
func (e expr) innerRegexp(value any) expr {
	e.e = clause.Expr{
		SQL:  "? REGEXP ?",
		Vars: []any{e.RawExpr(), value},
	}
	return e
}

// NOT(expr REGEXP ?)
func (e expr) innerNotRegexp(value any) expr {
	e.e = clause.NotConditions{
		Exprs: []clause.Expression{
			clause.Expr{
				SQL:  "? REGEXP ?",
				Vars: []any{e.RawExpr(), value},
			},
		},
	}
	return e
}

// expr BETWEEN ? AND ?
func (e expr) innerBetween(left, right any) expr {
	e.e = clause.Expr{
		SQL:  "? BETWEEN ? AND ?",
		Vars: []any{e.RawExpr(), left, right},
	}
	return e
}

// NOT (expr BETWEEN ? AND ?)
func (e expr) innerNotBetween(left, right any) expr {
	e.e = clause.NotConditions{
		Exprs: []clause.Expression{
			clause.Expr{
				SQL:  "? BETWEEN ? AND ?",
				Vars: []any{e.RawExpr(), left, right},
			},
		},
	}
	return e
}

// SUM(expr)
func (e expr) innerSum() expr {
	e.e = clause.Expr{
		SQL:  "SUM(?)",
		Vars: []any{e.RawExpr()},
	}
	return e
}

// value:
// time.Duration - DATE_ADD(expr, INTERVAL ? MICROSECOND)
// other = expr+?
func (e expr) innerAdd(value any) expr {
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

// value:
// time.Duration - DATE_SUB(expr, INTERVAL ? MICROSECOND)
// other = expr-?
func (e expr) innerSub(value any) expr {
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

// col*? or (expr)*?
func (e expr) innerMul(value any) expr {
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

// col/? or (expr)/?
func (e expr) innerDiv(value any) expr {
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

// col%? or (expr)%?
func (e expr) innerMod(value any) expr {
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

// col DIV ? or (expr) DIV ?
func (e expr) innerFloorDiv(value any) expr {
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

// FLOOR(expr)
func (e expr) innerFloor() expr {
	e.e = clause.Expr{
		SQL:  "FLOOR(?)",
		Vars: []any{e.RawExpr()},
	}
	return e
}

// ROUND(expr, ?)
func (e expr) innerRound(decimals int) expr {
	e.e = clause.Expr{
		SQL:  "ROUND(?, ?)",
		Vars: []any{e.RawExpr(), decimals},
	}
	return e
}

// innerFindInSet equal to FIND_IN_SET(expr, targetList)
func (e expr) innerFindInSet(targetList string) expr {
	e.e = clause.Expr{
		SQL:  "FIND_IN_SET(?, ?)",
		Vars: []any{e.RawExpr(), targetList},
	}
	return e
}

// innerFindInSetWith equal to FIND_IN_SET(target, expr)
func (e expr) innerFindInSetWith(target string) expr {
	e.e = clause.Expr{
		SQL:  "FIND_IN_SET(?, ?)",
		Vars: []any{target, e.RawExpr()},
	}
	return e
}

// col>>? or (expr)>>?
func (e expr) innerRightShift(value any) expr {
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

// col<<? or (expr)<<?
func (e expr) innerLeftShift(value any) expr {
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

// col^? or (expr)^?
func (e expr) innerBitXor(value any) expr {
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

// col&? or (expr)&?
func (e expr) innerBitAnd(value any) expr {
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

// col|? or (expr)|?
func (e expr) innerBitOr(value any) expr {
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

// ~col or ~(expr)
func (e expr) innerBitFlip() expr {
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

// expr AND ?
func (e expr) innerAnd(value any) expr {
	e.e = clause.Expr{
		SQL:  "? AND ?",
		Vars: []any{e.RawExpr(), value},
	}
	return e
}

// expr OR ?
func (e expr) innerOr(value any) expr {
	e.e = clause.Expr{
		SQL:  "? OR ?",
		Vars: []any{e.RawExpr(), value},
	}
	return e
}

// expr XOR ?
func (e expr) innerXor(value any) expr {
	e.e = clause.Expr{
		SQL:  "? XOR ?",
		Vars: []any{e.RawExpr(), value},
	}
	return e
}

// TRIM(BOTH ? FROM ?)
func (e expr) innerTrim(remStr string) expr {
	e.e = clause.Expr{SQL: "TRIM(BOTH ? FROM ?)", Vars: []any{remStr, e.RawExpr()}}
	return e
}

// TRIM(LEADING ? FROM ?)
func (e expr) innerLTrim(remStr string) expr {
	e.e = clause.Expr{SQL: "TRIM(LEADING ? FROM ?)", Vars: []any{remStr, e.RawExpr()}}
	return e
}

// TRIM(TRAILING ? FROM ?)
func (e expr) innerRTrim(remStr string) expr {
	e.e = clause.Expr{SQL: "TRIM(TRAILING ? FROM ?)", Vars: []any{remStr, e.RawExpr()}}
	return e
}

// TRIM(?)
func (e expr) innerTrimSpace() expr {
	e.e = clause.Expr{SQL: "TRIM(?)", Vars: []any{e.RawExpr()}}
	return e
}

// LTRIM(?)
func (e expr) innerLTrimSpace() expr {
	e.e = clause.Expr{SQL: "LTRIM(?)", Vars: []any{e.RawExpr()}}
	return e
}

// RTRIM(?)
func (e expr) innerRTrimSpace() expr {
	e.e = clause.Expr{SQL: "RTRIM(?)", Vars: []any{e.RawExpr()}}
	return e
}
