package assist

func (x *Executor[T]) IntoSubQueryExpr() Field {
	return SubQuery(x.IntoDB())
}

func (x *Executor[T]) IntoExistExpr() Expr {
	return Exist(x.IntoDB())
}

func (x *Executor[T]) IntoNotExistExpr() Expr {
	return NotExist(x.IntoDB())
}
