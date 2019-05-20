// Copyright (c) 2019 TerserGo
// 2019-05-20 10:42
// schema/column_schema.go

package schema

import "strings"

type ColumnSchema struct {
	TableName      string
	Name           string
	VarName        string
	Comment        string
	ColumnType     string // 列具体数据类型 varchar(64), int(11)
	DataType       string // 数据类型(无精度) varchar, int
	DataTypeLength int    // 数据精度
	DataTypeScale  int    // 数据小数精度
	DefaultValue   string
	IsNullable     bool
	IsPrimaryKey   bool
	IsEnum         bool
	GoDataType     string // Golang对应的基础类型
	LabelTag       string
}

func (c *ColumnSchema) SetIsPrimaryKey(v interface{}) {
	c.IsPrimaryKey = equalToString(v, "PRI")
}

func (c *ColumnSchema) SetIsNullable(v interface{}) {
	c.IsNullable = equalToString(v, "YES")
}

func (c *ColumnSchema) SetDataTypeLength(v interface{}) {
	i, err := toInt(v)
	if err == nil && i > 0 {
		c.DataTypeLength = i
	}
}

func (c *ColumnSchema) SetDataTypeScale(v interface{}) {
	i, err := toInt(v)
	if err == nil && i > 0 {
		c.DataTypeScale = i
	}
}

func (c *ColumnSchema) InitGoDataType() {
	if len(c.GoDataType) > 0 {
		return
	}
	c.LabelTag = "`"

	baseType, exists := dataTypeMaps[c.DataType]

	if !exists {
		c.GoDataType = c.DataType
		return
	}

	c.GoDataType = baseType

	switch strings.ToLower(c.DataType) {
	case "int":
		if strings.Index(c.ColumnType, "unsigned") > 0 {
			c.GoDataType = "u" + baseType
		}
	case "enum":
		c.initEnumType(c.ColumnType)
	default:
	}

}

func (c *ColumnSchema) initEnumType(enumType string) {

}