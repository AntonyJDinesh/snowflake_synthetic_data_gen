package qry

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
)

const SynQryTmplName = "synthetic_qry.tmpl"
const SynQryTmplFile = "internal/qry/tmpl/synthetic_qry.tmpl"

type SynQryData struct {
	Table        *Table
	TotalRecords int
	Columns      []*Column
}

var syn_qry_tmpl *template.Template

func init() {
	func_map := template.FuncMap{"col_gen": SynQryColumnGenerator, "last_idx": IsLastIndex}
	var err error
	syn_qry_tmpl, err = template.New(SynQryTmplName).Funcs(func_map).ParseFiles(SynQryTmplFile)

	if err != nil {
		panic(err)
	}
}

func GetSynQry(qry_data *SynQryData) (qry string, err error) {
	if qry_data == nil {
		err = errors.New("SynQryData is nil")
		return
	}

	var buf bytes.Buffer

	err = syn_qry_tmpl.ExecuteTemplate(&buf, SynQryTmplName, qry_data)
	if err != nil {
		panic(err)
	}

	qry = buf.String()

	return
}

func SynQryColumnGenerator(column *Column) (col_gen_fn string, err error) {
	switch column.DataType {
	case ColumnDataTypeSeq:
		col_gen_fn = "ROW_NUMBER() over (order by SEQ1())"
	case ColumnDataTypeUUID:
		col_gen_fn = "UUID_STRING()"
	case ColumnDataTypeVarchar:
		col_gen_fn = fmt.Sprintf("RANDSTR(UNIFORM(%d, %d, RANDOM()), RANDOM())", column.LengthMin, column.LengthMax)
	case ColumnDataTypeInteger:
		col_gen_fn = fmt.Sprintf("UNIFORM(%d, %d, RANDOM())", column.RangeMin, column.RangeMax)
	case ColumnDataTypeFloat:
		col_gen_fn = fmt.Sprintf("ABS(TRUNC(NORMAL(100, 100, RANDOM()), %d))", column.DecimalScale)
	case ColumnDataTypeBoolean:
		col_gen_fn = "TO_BOOLEAN(UNIFORM(0, 1, random()))"
	case ColumnDataTypeDate:
		col_gen_fn = "DATEADD('day', uniform(-3650, 0, random()), CURRENT_DATE())"
	default:
		err = errors.ErrUnsupported
	}

	return
}
