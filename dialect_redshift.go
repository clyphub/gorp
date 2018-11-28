package gorp

type RedshiftDialect struct {
	PostgresDialect
}

func (d RedshiftDialect) InsertAutoIncrToTarget(exec SqlExecutor, insertSql string, target interface{}, params ...interface{}) error {
	rows, err := exec.Query(insertSql, params...)
	defer rows.Close()
	return err
}

func (d RedshiftDialect) AutoIncrBindValue() string {
	return ""
}

func (d RedshiftDialect) AutoIncrInsertSuffix(col *ColumnMap) string {
	return ""
}
