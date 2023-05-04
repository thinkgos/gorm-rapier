package assist

import (
	"testing"

	"gorm.io/gorm"
)

func Test_Columns_SubQuery(t *testing.T) {
	var dummy []Dict

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []interface{}
		want     string
	}{
		{
			name: "in invalid query",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(xx_Dict.Id.IntoColumns().In(nil)).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "in sub query",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(xx_Dict.Id.IntoColumns().In(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id)).Where(xx_Dict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` IN (SELECT `dict`.`id` FROM `dict` WHERE `dict`.`score` > ?)",
		},
		{
			name: "in sub query - (multiple fields)",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(NewColumns(xx_Dict.Id, xx_Dict.Name).In(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id, xx_Dict.Name)).Where(xx_Dict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT * FROM `dict` WHERE (`dict`.`id`,`dict`.`name`) IN (SELECT `dict`.`id`,`dict`.`name` FROM `dict` WHERE `dict`.`score` > ?)",
		},
		{
			name: "in sub query - (no field)",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(NewColumns().In(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id)).Where(xx_Dict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "in value",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(xx_Dict.Id.IntoColumns().In(Values([]any{1, 100}))).
				Find(&dummy),
			wantVars: []any{1, 100},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` IN (?,?)",
		},
		{
			name: "in value - multiple field",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(NewColumns(xx_Dict.Id, xx_Dict.Score).In(Values([][]any{{100, 200}, {1, 2}}))).
				Find(&dummy),
			wantVars: []any{100, 200, 1, 2},
			want:     "SELECT * FROM `dict` WHERE (`dict`.`id`, `dict`.`score`) IN ((?,?),(?,?))",
		},
		{
			name: "in value - (no field)",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(NewColumns().In(Values([][]any{{100, 200}, {1, 2}}))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "not in sub query",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(xx_Dict.Id.IntoColumns().NotIn(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id)).Where(xx_Dict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT * FROM `dict` WHERE NOT `dict`.`id` IN (SELECT `dict`.`id` FROM `dict` WHERE `dict`.`score` > ?)",
		},
		{
			name: "not in sub query(no fields)",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(NewColumns().NotIn(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id)).Where(xx_Dict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE NOT ",
		},
		{
			name: "not in value",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(NewColumns(xx_Dict.Id, xx_Dict.Score).NotIn(Values([][]any{{100, 200}, {1, 2}}))).
				Find(&dummy),
			wantVars: []any{100, 200, 1, 2},
			want:     "SELECT * FROM `dict` WHERE NOT (`dict`.`id`, `dict`.`score`) IN ((?,?),(?,?))",
		},
		{
			name: "not in value(no fields)",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(NewColumns().NotIn(Values([][]any{{100, 200}, {1, 2}}))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE NOT ",
		},
		{
			name: "eq",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(
					xx_Dict.Id.IntoColumns().Eq(
						newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Max())),
					),
				).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` = (SELECT MAX(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "eq - (no field)",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(
					NewColumns().Eq(
						newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Max())),
					),
				).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "neq",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(xx_Dict.Id.IntoColumns().Neq(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` <> (SELECT MAX(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "neq(no fields)",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(NewColumns().Neq(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "gt",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(xx_Dict.Id.IntoColumns().Gt(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` > (SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "gt - (no field)",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(NewColumns().Gt(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "gte",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(xx_Dict.Id.IntoColumns().Gte(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` >= (SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "gte - (no field)",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(NewColumns().Gte(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "lt",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(xx_Dict.Id.IntoColumns().Lt(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` < (SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "lt - (no field)",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(NewColumns().Lt(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "lte",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(xx_Dict.Id.IntoColumns().Lte(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` <= (SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "lte - (no field)",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(NewColumns().Lte(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "exist",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(Exist(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE EXISTS(SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "not exist",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(NotExist(newDb().Scopes(xx_Dict.Xc_Model(), Select(xx_Dict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE NOT EXISTS(SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_Field_IntoColumns(t *testing.T) {
	var dummy []Dict

	fieldRaw := NewRaw("`age`")
	fieldField := NewField("", "age")
	fieldInt := NewInt("", "age")
	fieldInt8 := NewInt8("", "age")
	fieldInt16 := NewInt16("", "age")
	fieldInt32 := NewInt32("", "age")
	fieldInt64 := NewInt64("", "age")
	fieldUint := NewUint("", "age")
	fieldUint8 := NewUint8("", "age")
	fieldUint16 := NewUint16("", "age")
	fieldUint32 := NewUint32("", "age")
	fieldUint64 := NewUint64("", "age")
	fieldBool := NewBool("", "age")
	fieldString := NewString("", "age")
	fieldBytes := NewBytes("", "age")
	fieldFloat32 := NewFloat32("", "age")
	fieldFloat64 := NewFloat64("", "age")
	fieldDecimal := NewDecimal("", "age")
	fieldTime := NewTime("", "age")
	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "raw IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldRaw.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldRaw.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "field IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldField.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldField.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "int IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldInt.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldInt.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "int8 IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldInt8.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldInt8.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "int16 IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldInt16.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldInt16.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "int32 IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldInt32.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldInt32.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "int64 IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldInt64.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldInt64.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "uint IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldUint.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldUint.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "uint8 IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldUint8.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldUint8.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "uint16 IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldUint16.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldUint16.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "uint32 IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldUint32.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldUint32.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "uint64 IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldUint64.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldUint64.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},

		{
			name: "bool IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldBool.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldBool.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "string IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldString.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldString.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "bytes IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldBytes.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldBytes.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "float32 IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldFloat32.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldFloat32.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "float64 IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldFloat64.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldFloat64.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "decimal IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldDecimal.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldDecimal.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
		{
			name: "time IntoColumns",
			db: newDb().Model(xx_Dict.X_Model()).
				Where(fieldTime.IntoColumns().Eq(newDb().Model(xx_Dict.X_Model()).Scopes(Select(fieldDecimal.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `age` = (SELECT MAX(`age`) FROM `dict`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
			})
		})
	}
}
