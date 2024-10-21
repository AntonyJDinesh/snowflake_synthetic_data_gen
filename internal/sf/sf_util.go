package sf

import (
	"database/sql"
	"fmt"
)

type TableMetaData struct {
	Table   *Table
	Coulmns []*Column
}

func (tmd *TableMetaData) String() string {
	return fmt.Sprintf("{Table: %s,\nCoulmns: %s}", tmd.Table, tmd.Coulmns)
}

type Column struct {
	Name     string
	DataType string
	IsPKey   bool
	IsUKey   bool
}

func (col *Column) String() string {
	return fmt.Sprintf("{Name: %v, DataType: %v, IsPKey: %v, IsUKey: %v}\n", col.Name, col.DataType, col.IsPKey, col.IsUKey)
}

type Table struct {
	DatabaseName string
	SchemaName   string
	TableName    string
}

func (tbl *Table) String() string {
	return fmt.Sprintf("{DatabaseName: %v, SchemaName: %v, TableName: %v}", tbl.DatabaseName, tbl.SchemaName, tbl.TableName)
}

func GetTableMetaData(tbl *Table, db *sql.DB) (*TableMetaData, error) {
	rows, err := db.Query(fmt.Sprintf("desc TABLE %s.%s.%s", tbl.DatabaseName, tbl.SchemaName, tbl.TableName))
	if err != nil {
		return nil, err
	}

	cmd := &TableMetaData{}
	cmd.Table = &Table{}
	*cmd.Table = *tbl
	cmd.Coulmns = make([]*Column, 0)

	for rows.Next() {
		var name, typ, kind, nullable, pkey, ukey string
		var defa, check, exp, comment, policy_name, privacy_domain sql.NullString

		err = rows.Scan(&name, &typ, &kind, &nullable, &defa, &pkey, &ukey, &check, &exp, &comment, &policy_name, &privacy_domain)

		if err != nil {
			return nil, err
		}

		col := &Column{Name: name, DataType: typ}

		if pkey == "Y" {
			col.IsPKey = true
		}

		if ukey == "Y" {
			col.IsUKey = true
		}

		cmd.Coulmns = append(cmd.Coulmns, col)
	}

	return cmd, nil
}

func ExceQry(qry string, db *sql.DB) error {
	_, err := db.Exec(qry)
	if err != nil {
		return err
	}

	return nil
}
