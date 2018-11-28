package gorp

type RedshiftDialect struct {
	PostgresDialect
}

func (d RedshiftDialect) InsertAutoIncrToTarget(exec SqlExecutor, insertSql string, _ interface{}, params ...interface{}) error {
	rows, err := exec.Query(insertSql, params...)
	defer rows.Close()
	return err
}

func (d RedshiftDialect) AutoIncrBindValue() string {
	// returning an empty string here makes gorp skip adding this column to the INSERT SQL
	// this is needed because Redshift does not allow specifying values for IDENTITY columns
	return ""
}

func (d RedshiftDialect) AutoIncrInsertSuffix(col *ColumnMap) string {
	// Redshift has no support for RETURNING values after an insert
	return ""
}
