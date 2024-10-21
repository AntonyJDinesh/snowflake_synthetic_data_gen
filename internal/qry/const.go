package qry

type ColumnDataType string

const (
	ColumnDataTypeNumber    ColumnDataType = "NUMBER"
	ColumnDataTypeVarchar   ColumnDataType = "VARCHAR"
	ColumnDataTypeFloat     ColumnDataType = "FLOAT"
	ColumnDataTypeTimestamp ColumnDataType = "TIMESTAMP"
	ColumnDataTypeDate      ColumnDataType = "DATE"
	ColumnDataTypeChar      ColumnDataType = "CHAR"
	ColumnDataTypeBoolean   ColumnDataType = "BOOLEAN"

	ColumnDataTypeSeq  ColumnDataType = "SEQ"
	ColumnDataTypeUUID ColumnDataType = "UUID"

	ColumnDataTypeInteger ColumnDataType = "INTEGER"
	ColumnDataTypeBigint  ColumnDataType = "BIGINT"
	ColumnDataTypeDouble  ColumnDataType = "DOUBLE"
)

type Table struct {
	DatabaseName string
	SchemaName   string
	TableName    string
}

type Column struct {
	Name         string
	DataType     ColumnDataType
	LengthMin    int
	LengthMax    int
	RangeMin     int
	RangeMax     int
	DecimalScale int
}

func IsLastIndex(i int, l int) bool {
	return i+1 == l
}
