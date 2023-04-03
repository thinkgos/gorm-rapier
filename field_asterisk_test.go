package assist

import "testing"

func Test_Expr_Asterisk(t *testing.T) {
	tests := []struct {
		name         string
		Expr         Expr
		ExpectedVars []any
		Result       string
	}{
		{
			Expr:   Star,
			Result: "*",
		},
		{
			Expr:   NewAsterisk("user"),
			Result: "`user`.*",
		},
		{
			Expr:   Star.Count(),
			Result: "COUNT(*)",
		},
		{
			Expr:   Star.Distinct().Count(),
			Result: "COUNT(DISTINCT *)",
		},
		{
			Expr:   NewAsterisk("user").Count(),
			Result: "COUNT(`user`.*)",
		},
		{
			Expr:   NewAsterisk("user").Distinct().Count(),
			Result: "COUNT(DISTINCT `user`.*)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.Expr, tt.Result, tt.ExpectedVars)
		})
	}
}
