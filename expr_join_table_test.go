package rapier

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		name     string
		join     JoinTableExpr
		wantSQL  string
		wantVars []any
	}{
		{
			name: "LEFT JOIN",
			join: JoinTableExpr{
				Join: clause.Join{
					Type:  clause.LeftJoin,
					Table: clause.Table{},
					ON: clause.Where{
						Exprs: []clause.Expression{
							clause.Eq{
								Column: clause.Column{
									Table: "d",
									Name:  "pid",
								},
								Value: 1000,
							},
						},
					},
				},
				TableExpr: TableExpr(From{
					Alias:    "d",
					SubQuery: db.Table("dict").Where("name = ?", "Tom"),
				})(db.Session(&gorm.Session{NewDB: true})).Statement.TableExpr,
			},
			wantSQL:  "LEFT JOIN (SELECT * FROM `dict` WHERE name = ?) AS `d` ON `d`.`pid` = ?",
			wantVars: []any{"Tom", 1000},
		},
		{
			name: "RIGHT JOIN",
			join: JoinTableExpr{
				Join: clause.Join{
					Type:  clause.RightJoin,
					Table: clause.Table{},
					ON: clause.Where{
						Exprs: []clause.Expression{
							clause.Eq{
								Column: clause.Column{
									Table: "d",
									Name:  "pid",
								},
								Value: 1000,
							},
						},
					},
				},
				TableExpr: TableExpr(From{
					Alias:    "d",
					SubQuery: db.Table("dict").Where("name = ?", "Tom"),
				})(db.Session(&gorm.Session{NewDB: true})).Statement.TableExpr,
			},
			wantSQL:  "RIGHT JOIN (SELECT * FROM `dict` WHERE name = ?) AS `d` ON `d`.`pid` = ?",
			wantVars: []any{"Tom", 1000},
		},
		{
			name: "INNER JOIN",
			join: JoinTableExpr{
				Join: clause.Join{
					Type:  clause.InnerJoin,
					Table: clause.Table{},
					ON: clause.Where{
						Exprs: []clause.Expression{
							clause.Eq{
								Column: clause.Column{
									Table: "d",
									Name:  "pid",
								},
								Value: 1000,
							},
						},
					},
				},
				TableExpr: TableExpr(From{
					Alias:    "d",
					SubQuery: db.Table("dict").Where("name = ?", "Tom"),
				})(db.Session(&gorm.Session{NewDB: true})).Statement.TableExpr,
			},
			wantSQL:  "INNER JOIN (SELECT * FROM `dict` WHERE name = ?) AS `d` ON `d`.`pid` = ?",
			wantVars: []any{"Tom", 1000},
		},
		{
			name: "CROSS JOIN",
			join: JoinTableExpr{
				Join: clause.Join{
					Type:  clause.CrossJoin,
					Table: clause.Table{},
					ON: clause.Where{
						Exprs: []clause.Expression{
							clause.Eq{
								Column: clause.Column{
									Table: "d",
									Name:  "pid",
								},
								Value: 1000,
							},
						},
					},
				},
				TableExpr: TableExpr(From{
					Alias:    "d",
					SubQuery: db.Table("dict").Where("name = ?", "Tom"),
				})(db.Session(&gorm.Session{NewDB: true})).Statement.TableExpr,
			},
			wantSQL:  "CROSS JOIN (SELECT * FROM `dict` WHERE name = ?) AS `d` ON `d`.`pid` = ?",
			wantVars: []any{"Tom", 1000},
		},
		{
			name: "INNER JOIN",
			join: JoinTableExpr{
				Join: clause.Join{
					Type:  clause.InnerJoin,
					Table: clause.Table{},
					Using: []string{"pid", "name"},
				},
				TableExpr: TableExpr(From{
					Alias:    "d",
					SubQuery: db.Table("dict").Where("name = ?", "Tom"),
				})(db.Session(&gorm.Session{NewDB: true})).Statement.TableExpr,
			},
			wantSQL:  "INNER JOIN (SELECT * FROM `dict` WHERE name = ?) AS `d` USING (`pid`,`name`)",
			wantVars: []any{"Tom"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt := NewStatement()
			tt.join.Build(stmt)
			if gotSQL := stmt.SQL.String(); tt.wantSQL != stmt.SQL.String() {
				t.Errorf("\nSQL:\n\twant: %v\n\tgot: %v", tt.wantSQL, gotSQL)
			}
			if !reflect.DeepEqual(stmt.Vars, tt.wantVars) {
				t.Errorf("\nVars:\n\twant: %+v\n\tgot: %+v", tt.wantSQL, stmt.Vars)
			}
		})
	}
}
