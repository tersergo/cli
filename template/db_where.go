package template

var DBWhere = `// auto-generated by terser-cli 
// struct: DBWhere Where查询字段类
// type TableField string 表字段
package model

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type TableField string

type DBWhere struct {
	Name  TableField
	Query string
	Binds []interface{}
}

// sql: queryField AND queryField
func DBWhereAnd(dbWhere ...*DBWhere) *DBWhere {
	if len(dbWhere) == 0 {
		return nil
	}

	return mergeDBWhere(" AND ", dbWhere)
}

// sql: queryField OR queryField
func DBWhereOr(dbWhere ...*DBWhere) *DBWhere {
	if len(dbWhere) == 0 {
		return nil
	}

	return mergeDBWhere(" OR ", dbWhere)
}

func mergeDBWhere(joinSep string, fieldList []*DBWhere) *DBWhere {
	var (
		parts []string
		binds []interface{}
	)

	for _, field := range fieldList {
		if field == nil || field.Query == "" {
			continue
		}

		parts = append(parts, field.Query)
		if field.HasBind() {
			binds = append(binds, field.Binds)
		}
	}

	length := len(parts)
	if length == 0 {
		return nil
	}

	where := strings.Join(parts, joinSep)
	if length > 1 {
		where = fmt.Sprintf("( %s )", where)
	}

	return &DBWhere{
		Query: where,
		Binds: binds,
	}
}

// sql: name > value
func (field TableField) GT(value interface{}) *DBWhere {
	if value == nil {
		return nil
	}
	return field.WhereFormat("%s > ?", value)
}

// sql: name >= value
func (field TableField) GTE(value interface{}) *DBWhere {
	if value == nil {
		return nil
	}
	return field.WhereFormat("%s >= ?", value)
}

// sql: name < value
func (field TableField) LT(value interface{}) *DBWhere {
	if value == nil {
		return nil
	}
	return field.WhereFormat("%s < ?", value)
}

// sql: name <= value
func (field TableField) LTE(value interface{}) *DBWhere {
	if value == nil {
		return nil
	}
	return field.WhereFormat("%s <= ?", value)
}

// sql: name IS NULL
func (field TableField) IsNull() *DBWhere {
	return field.WhereFormat("%s IS NULL")
}

// sql: name IS NOT NULL
func (field TableField) NotNull() *DBWhere {
	return field.WhereFormat("%s IS NOT NULL")
}

// sql: name LIKE %value%
func (field TableField) Like(value interface{}) *DBWhere {
	if value == nil {
		return nil
	}
	v := fmt.Sprintf("%%%s%%", value)
	return field.WhereFormat("%s LIKE ?", v)
}

// sql: name = value
func (field TableField) Equal(value interface{}) *DBWhere {
	if value == nil {
		return nil
	}
	return field.WhereFormat("%s = ?", value)
}

// sql: name <> value
func (field TableField) NotEqual(value interface{}) *DBWhere {
	if value == nil {
		return nil
	}
	return field.WhereFormat("%s <> ?", value)
}

// sql: name IN (values)
func (field TableField) In(values ...interface{}) *DBWhere {
	switch len(values) {
	case 0:
		return nil
	case 1:
		return field.WhereFormat("%s IN (?)", values[0])
	default:
		return field.WhereFormat("%s IN (?)", values)
	}
}

// sql: FIND_IN_SET(value, name)
func (field TableField) FindInSet(value interface{}) *DBWhere {
	if value == nil {
		return nil
	}
	return field.WhereFormat("FIND_IN_SET(?, %s)", value)
}

// sql: name BETWEEN begin AND end
func (field TableField) Between(begin, end interface{}) *DBWhere {
	if begin == nil || end == nil {
		return nil
	}
	return field.WhereFormat("%s BETWEEN ? AND ?", begin, end)
}

func (field TableField) WhereFormat(format string, value ...interface{}) *DBWhere {
	if field == "" {
		return nil
	}

	return &DBWhere{
		Name:  field,
		Query: fmt.Sprintf(format, field),
		Binds: value,
	}
}

func (query *DBWhere) HasBind() bool {
	return len(query.Binds) > 0
}

func (query *DBWhere) SetDBWhere(db *gorm.DB) *gorm.DB {
	if query.Query == "" {
		return db
	}

	if query.HasBind() {
		return db.Query(query.Query, query.Binds...)
	}

	return db.Query(query.Query)
}

`