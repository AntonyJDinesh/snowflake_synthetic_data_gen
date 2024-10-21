package main

import (
	"fmt"

	"redhat.com/ddis/synthetic_data_gen/internal/qry"
)

func main() {
	qry_str, _ := qry.GetSynQry(
		&qry.SynQryData{
			Table: &qry.Table{
				DatabaseName: "SANDBOX_DB_ADINESH",
				SchemaName:   "TEST",
				TableName:    "PARTY",
			},
			TotalRecords: 100,
			Columns: []*qry.Column{
				{
					Name:     "id",
					DataType: qry.ColumnDataTypeSeq,
				},
				{
					Name:      "name",
					DataType:  qry.ColumnDataTypeVarchar,
					LengthMin: 5,
					LengthMax: 10,
				},
				{
					Name:     "is_active",
					DataType: qry.ColumnDataTypeBoolean,
				},
				{
					Name:     "no_of_employee",
					DataType: qry.ColumnDataTypeInteger,
					RangeMin: 5000,
					RangeMax: 10000,
				},
				{
					Name:         "annual_turnover_in_millions_usd",
					DataType:     qry.ColumnDataTypeFloat,
					DecimalScale: 2,
				},
				{
					Name:     "date_incorporation",
					DataType: qry.ColumnDataTypeDate,
				},
			},
		},
	)

	fmt.Println(qry_str)

	// db, err := sf.GetConnection()
	// if err != nil {
	// 	panic(err)
	// }

	// tbl := &sf.Table{DatabaseName: "SANDBOX_DB_ADINESH", SchemaName: "TEST", TableName: "PARTY"}
	// cmd, err := sf.GetTableMetaData(tbl, db)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(cmd)

	ano_qry, _ := qry.GetAnonymQry(&qry.AnonymQryData{
		SrcTable: &qry.Table{DatabaseName: "SANDBOX_DB_ADINESH", SchemaName: "TEST", TableName: "PARTY"},
		DstTable: &qry.Table{DatabaseName: "SANDBOX_DB_ADINESH", SchemaName: "TEST", TableName: "PARTY_ANONYM"},
		Columns: []*qry.Column{
			{
				Name:     "ID",
				DataType: qry.ColumnDataTypeNumber,
			},
			{
				Name:     "NAME",
				DataType: qry.ColumnDataTypeVarchar,
			},
			{
				Name:     "IS_ACTIVE",
				DataType: qry.ColumnDataTypeBoolean,
			},
			{
				Name:     "NO_OF_EMPLOYEE",
				DataType: qry.ColumnDataTypeNumber,
			},
			{
				Name:     "ANNUAL_TURNOVER_IN_MILLIONS_USD",
				DataType: qry.ColumnDataTypeFloat,
			},
			{
				Name:     "DATE_INCORPORATION",
				DataType: qry.ColumnDataTypeDate,
			},
		},
	})

	fmt.Println(ano_qry)
}
