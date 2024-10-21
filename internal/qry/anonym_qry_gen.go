package qry

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
)

const anonym_qry_tmpl_name = "anonymized_qry.tmpl"
const anonym_qry_tmpl_file = "internal/qry/tmpl/anonymized_qry.tmpl"

type AnonymQryData struct {
	SrcTable *Table
	DstTable *Table
	Columns  []*Column
}

var anonym_qry_tmpl *template.Template

func init() {
	func_map := template.FuncMap{"col_gen": AnonymQryColumnGenerator, "last_idx": IsLastIndex}
	var err error
	anonym_qry_tmpl, err = template.New(anonym_qry_tmpl_name).Funcs(func_map).ParseFiles(anonym_qry_tmpl_file)

	if err != nil {
		panic(err)
	}
}

func GetAnonymQry(qry_data *AnonymQryData) (qry string, err error) {
	if qry_data == nil {
		err = errors.New("SynQryData is nil")
		return
	}

	var buf bytes.Buffer

	err = anonym_qry_tmpl.ExecuteTemplate(&buf, anonym_qry_tmpl_name, qry_data)
	if err != nil {
		panic(err)
	}

	qry = buf.String()

	return
}

func AnonymQryColumnGenerator(column *Column) (col_gen_fn string, err error) {
	switch column.DataType {
	case ColumnDataTypeNumber:
		col_gen_fn = fmt.Sprintf("anonym_factor(%s, $passphrase)", column.Name)
	case ColumnDataTypeVarchar:
		col_gen_fn = fmt.Sprintf("randstr(uniform(5, 10, anonym_factor(%s, $passphrase)), anonym_factor(%s, $passphrase))", column.Name, column.Name)
	case ColumnDataTypeBoolean:
		col_gen_fn = column.Name
	case ColumnDataTypeFloat:
		col_gen_fn = fmt.Sprintf("NORMAL(100, 100, anonym_factor(%s, $passphrase))", column.Name)
	case ColumnDataTypeDate:
		col_gen_fn = fmt.Sprintf("DATEADD('day', uniform(-3650, 0, random()), %s)", column.Name)
	default:
		err = errors.ErrUnsupported
	}

	return
}
