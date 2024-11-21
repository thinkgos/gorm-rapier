package rapier

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
				wantVars: []any{10, 10},
				want:     "SELECT * FROM `dict` LIMIT ? OFFSET ?",
			},
			{
				name:     "page < 1",
				db:       newDb().Model(Dict{}).Clauses(Paginate(0, 10)).Find([]Dict{}),
				wantVars: []any{10},
				want:     "SELECT * FROM `dict` LIMIT ?",
			},
			{
				name:     "perPage < 1",
				db:       newDb().Model(Dict{}).Clauses(Paginate(1, 0)).Find([]Dict{}),
				wantVars: []any{50},
				want:     "SELECT * FROM `dict` LIMIT ?",
			},
			{
				name:     "perPage > DefaultMaxPerPage",
				db:       newDb().Model(Dict{}).Clauses(Paginate(1, DefaultMaxPerPage+1)).Find([]Dict{}),
				wantVars: []any{500},
				want:     "SELECT * FROM `dict` LIMIT ?",
			},
			{
				name:     "customer perPage > maxPerPage",
				db:       newDb().Model(Dict{}).Clauses(Paginate(1, 201, 200)).Find([]Dict{}),
				wantVars: []any{200},
				want:     "SELECT * FROM `dict` LIMIT ?",
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
			})
		}
	})

	t.Run("Pagination", func(t *testing.T) {
		tests := []struct {
			name     string
			db       *gorm.DB
			wantVars []any
			want     string
		}{
			{
				name:     "in range",
				db:       newDb().Model(Dict{}).Scopes(Pagination(2, 10)).Find([]Dict{}),
				wantVars: []any{10, 10},
				want:     "SELECT * FROM `dict` LIMIT ? OFFSET ?",
			},
			{
				name:     "page < 1",
				db:       newDb().Model(Dict{}).Scopes(Pagination(0, 10)).Find([]Dict{}),
				wantVars: []any{10},
				want:     "SELECT * FROM `dict` LIMIT ?",
			},
			{
				name:     "perPage < 1",
				db:       newDb().Model(Dict{}).Scopes(Pagination(1, 0)).Find([]Dict{}),
				wantVars: []any{50},
				want:     "SELECT * FROM `dict` LIMIT ?",
			},
			{
				name:     "perPage > DefaultMaxPerPage",
				db:       newDb().Model(Dict{}).Scopes(Pagination(1, DefaultMaxPerPage+1)).Find([]Dict{}),
				wantVars: []any{500},
				want:     "SELECT * FROM `dict` LIMIT ?",
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
			})
		}
	})
	t.Run("Limit", func(t *testing.T) {
		tests := []struct {
			name     string
			db       *gorm.DB
			wantVars []any
			want     string
		}{
			{
				name:     "in range",
				db:       newDb().Model(Dict{}).Scopes(Limit(2, 10)).Find([]Dict{}),
				wantVars: []any{10, 10},
				want:     "SELECT * FROM `dict` LIMIT ? OFFSET ?",
			},
			{
				name:     "page < 1",
				db:       newDb().Model(Dict{}).Scopes(Limit(0, 10)).Find([]Dict{}),
				wantVars: []any{10},
				want:     "SELECT * FROM `dict` LIMIT ?",
			},
			{
				name:     "perPage < 1",
				db:       newDb().Model(Dict{}).Scopes(Limit(1, 0)).Find([]Dict{}),
				wantVars: nil,
				want:     "SELECT * FROM `dict`",
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
			})
		}
	})
}
