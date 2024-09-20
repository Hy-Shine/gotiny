package orm

import (
	"strings"

	"gitee.com/chunanyong/zorm"
)

// Where for sql where
// example: Where(finder, "col1=? and col2=?", 1, "3") --> where col1=? and col2=? args:[1, "3"]
func Where(finder *zorm.Finder, cond string, values ...interface{}) *zorm.Finder {
	query, _ := finder.GetSQL()
	query = strings.ToLower(query)
	if strings.Contains(query, "where") {
		finder.Append("and "+cond, values...)
	} else {
		finder.Append("where "+cond, values...)
	}

	return finder
}

// Or adds sql or condition
// example: Or(finder, "id>?", 1) --> or id>? args:[1]
// example: Or(finder, "id=? and name != ?", 10, "John") --> or (id=? and name!=?) args:[10, "John"]
func Or(finder *zorm.Finder, cond string, values ...interface{}) *zorm.Finder {
	if strings.Trim(cond, " ") == "" {
		return finder
	}

	if strings.Count(cond, "?") > 1 {
		finder.Append("or ("+cond+")", values...)
	} else {
		finder.Append("or "+cond, values...)
	}

	return finder
}

// Order adds order columns
// example: Order(finder, "col1 desc,col2", "col3") -- > order by col1 desc, col2, col3
func Order(finder *zorm.Finder, cols ...string) *zorm.Finder {
	if len(cols) == 0 {
		return finder
	}

	var builder strings.Builder
	builder.WriteString("order by ")
	builder.WriteString(cols[0])
	if len(cols) > 1 {
		length := len(cols)
		for i := 1; i < length; i++ {
			builder.WriteString("," + cols[i])
		}
	}

	return finder.Append(builder.String())
}
