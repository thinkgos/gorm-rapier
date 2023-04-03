package assist

import (
	"fmt"
	"strings"

	"gorm.io/gorm/clause"
)

// UnixTimestamp use UNIX_TIMESTAMP([date])
func UnixTimestamp(date ...string) Int64 {
	if len(date) > 0 {
		return Int64{expr{e: clause.Expr{SQL: "UNIX_TIMESTAMP(?)", Vars: []any{date[0]}}}}
	}
	return Int64{expr{e: clause.Expr{SQL: "UNIX_TIMESTAMP()"}}}
}

// FromUnixTime use FROM_UNIXTIME(unix_timestamp[,format])
func FromUnixTime(date int64, format ...string) String {
	if len(format) > 0 && strings.TrimSpace(format[0]) != "" {
		return String{expr{e: clause.Expr{SQL: "FROM_UNIXTIME(?, ?)", Vars: []any{date, format[0]}}}}
	}
	return String{expr{e: clause.Expr{SQL: "FROM_UNIXTIME(?)", Vars: []any{date}}}}
}

// FromDays equal to FROM_DAYS(value)
func FromDays(value int) Time {
	return Time{expr{e: clause.Expr{SQL: fmt.Sprintf("FROM_DAYS(%d)", value)}}}
}

// CurDate return result of CURDATE()
func CurDate() Time {
	return Time{expr{e: clause.Expr{SQL: "CURDATE()"}}}
}

// CurTime return result of CURTIME()
func CurTime() Time {
	return Time{expr{e: clause.Expr{SQL: "CURTIME()"}}}
}

// Now return result of NOW()
func Now() Time {
	return Time{expr{e: clause.Expr{SQL: "NOW()"}}}
}
