package rapier

import "testing"

func Test_Expr_Time_Function(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		wantVars []any
		want     string
	}{
		{
			name: "UnixTimestamp use UnixTimestamp()",
			expr: UnixTimestamp(),
			want: "UNIX_TIMESTAMP()",
		},
		{
			name:     "UnixTimestamp use UnixTimestamp(date)",
			expr:     UnixTimestamp("2005-03-27 03:00:00").Mul(99),
			want:     "(UNIX_TIMESTAMP(?))*?",
			wantVars: []any{"2005-03-27 03:00:00", int64(99)},
		},
		{
			name:     "FromUnixTime use FROM_UNIXTIME(date)",
			expr:     FromUnixTime(1680743769),
			wantVars: []any{int64(1680743769)},
			want:     "FROM_UNIXTIME(?)",
		},
		{
			name:     "FromUnixTime use FROM_UNIXTIME(date,format)",
			expr:     FromUnixTime(1680743769, "%Y%m%d"),
			wantVars: []any{int64(1680743769), "%Y%m%d"},
			want:     "FROM_UNIXTIME(?, ?)",
		},
		{
			name:     "FROM_DAYS",
			expr:     FromDays(10000),
			wantVars: nil,
			want:     "FROM_DAYS(10000)",
		},
		{
			name:     "CURDATE",
			expr:     CurDate(),
			wantVars: nil,
			want:     "CURDATE()",
		},
		{
			name:     "CURTIME",
			expr:     CurTime(),
			wantVars: nil,
			want:     "CURTIME()",
		},
		{
			name:     "NOW",
			expr:     Now(),
			wantVars: nil,
			want:     "NOW()",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExpr(t, tt.expr, tt.want, tt.wantVars)
		})
	}
}
