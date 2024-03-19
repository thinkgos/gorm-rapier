package rapier

import "gorm.io/gorm/clause"

// EmptyExpr return a empty expression. it is nil
func EmptyExpr() Expr { return expr{e: clause.Expr{}} }

// Or return or condition
// form example: (`id1` OR `id2`)
func Or(exprs ...Expr) Expr {
	return &expr{e: clause.Or(IntoExpression(exprs...)...)}
}

// And return and condition
// form example: (`id1` AND `id2`)
func And(exprs ...Expr) Expr {
	return &expr{e: clause.And(IntoExpression(exprs...)...)}
}

// Not return not condition
// form example: NOT (`id1` AND `id2`)
func Not(exprs ...Expr) Expr {
	return &expr{e: clause.Not(IntoExpression(exprs...)...)}
}
