package assist

import (
	"testing"

	"gorm.io/gorm"
)

func Test_Paginate(t *testing.T) {
	t.Run("Paginate", func(t *testing.T) {
		tests := []struct {
			name     string
			db       *gorm.DB
			wantVars []any
			want     string
		}{
			{
				name:     "in range",
				db:       newDb().Model(Dict{}).Clauses(Paginate(2, 10)).Find([]Dict{}),
				wantVars: nil,
				want:     "SELECT * FROM `dict` LIMIT 10 OFFSET 10",
			},
			{
				name:     "page < 1",
				db:       newDb().Model(Dict{}).Clauses(Paginate(0, 10)).Find([]Dict{}),
				wantVars: nil,
				want:     "SELECT * FROM `dict` LIMIT 10",
			},
			{
				name:     "perPage < 1",
				db:       newDb().Model(Dict{}).Clauses(Paginate(1, 0)).Find([]Dict{}),
				wantVars: nil,
				want:     "SELECT * FROM `dict` LIMIT 50",
			},
			{
				name:     "perPage > DefaultMaxPerPage",
				db:       newDb().Model(Dict{}).Clauses(Paginate(1, DefaultMaxPerPage+1)).Find([]Dict{}),
				wantVars: nil,
				want:     "SELECT * FROM `dict` LIMIT 500",
			},
			{
				name:     "customer perPage > maxPerPage",
				db:       newDb().Model(Dict{}).Clauses(Paginate(1, 201, 200)).Find([]Dict{}),
				wantVars: nil,
				want:     "SELECT * FROM `dict` LIMIT 200",
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
			})
		}
	})

	t.Run("Paginate", func(t *testing.T) {
		tests := []struct {
			name     string
			db       *gorm.DB
			wantVars []any
			want     string
		}{
			{
				name:     "in range",
				db:       newDb().Model(Dict{}).Scopes(Pagination(2, 10)).Find([]Dict{}),
				wantVars: nil,
				want:     "SELECT * FROM `dict` LIMIT 10 OFFSET 10",
			},
			{
				name:     "page < 1",
				db:       newDb().Model(Dict{}).Scopes(Pagination(0, 10)).Find([]Dict{}),
				wantVars: nil,
				want:     "SELECT * FROM `dict` LIMIT 10",
			},
			{
				name:     "perPage < 1",
				db:       newDb().Model(Dict{}).Scopes(Pagination(1, 0)).Find([]Dict{}),
				wantVars: nil,
				want:     "SELECT * FROM `dict` LIMIT 50",
			},
			{
				name:     "perPage > DefaultMaxPerPage",
				db:       newDb().Model(Dict{}).Scopes(Pagination(1, DefaultMaxPerPage+1)).Find([]Dict{}),
				wantVars: nil,
				want:     "SELECT * FROM `dict` LIMIT 500",
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
			})
		}
	})
}
