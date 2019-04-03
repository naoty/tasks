package config

import (
	"fmt"
)

// Env represents environment variables.
type Env struct {
	DatabaseHost     string `required:"true" split_words:"true"`
	DatabaseName     string `required:"true" split_words:"true"`
	DatabasePassword string `required:"true" split_words:"true"`
	DatabaseUser     string `required:"true" split_words:"true"`
	Debug            bool   `default:"false"`
}

// GetDSN returns the Data Source Name of database.
func (env Env) GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s",
		env.DatabaseUser,
		env.DatabasePassword,
		env.DatabaseHost,
		env.DatabaseName,
	)
}
