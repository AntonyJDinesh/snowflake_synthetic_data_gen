package sf

import (
	"database/sql"

	"github.com/snowflakedb/gosnowflake"
)

func GetConnection() (*sql.DB, error) {

	config := &gosnowflake.Config{
		Account:       "hnb22701",
		User:          "adinesh",
		Database:      "SANDBOX_DB_ADINESH",
		Schema:        "TEST",
		Role:          "ROLE_ADINESH",
		Authenticator: gosnowflake.AuthTypeExternalBrowser,
	}

	if dsn, err := gosnowflake.DSN(config); err != nil {
		return nil, err
	} else {
		return sql.Open("snowflake", dsn)
	}
}
