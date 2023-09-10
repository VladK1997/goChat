package db_domain

type DB struct {
	// *ConfigW
	Error        error
	RowsAffected int64
	Statement    any //*Statement
	clone        int
}
