package rapier

import "gorm.io/gorm/clause"

// JoinTableExpr join clause with table expression(sub query...)
type JoinTableExpr struct {
	clause.Join
	TableExpr clause.Expression
}

func (join JoinTableExpr) Build(builder clause.Builder) {
	if join.Type != "" {
		builder.WriteString(string(join.Type)) // nolint:errcheck
		builder.WriteByte(' ')                 // nolint:errcheck
	}

	builder.WriteString("JOIN ") // nolint:errcheck
	if join.TableExpr != nil {
		join.TableExpr.Build(builder)
	}

	if len(join.ON.Exprs) > 0 {
		builder.WriteString(" ON ") // nolint:errcheck
		join.ON.Build(builder)
	} else if len(join.Using) > 0 {
		builder.WriteString(" USING (") // nolint:errcheck
		for idx, c := range join.Using {
			if idx > 0 {
				builder.WriteByte(',') // nolint:errcheck
			}
			builder.WriteQuoted(c)
		}
		builder.WriteByte(')') // nolint:errcheck
	}
}
