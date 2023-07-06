package assist

import (
	"testing"

	"gorm.io/gorm/clause"
)

func Test_intoClauseColumn(t *testing.T) {
	column := intoClauseColumn("t", "field", func(c clause.Column) clause.Column {
		c.Alias = "value"
		return c
	})
	if column.Alias != "value" {
		t.Error("must be a column alias <value>")
	}
}
