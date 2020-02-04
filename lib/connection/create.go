package connection

import (
	"database/sql"
	"fmt"
	"net/url"
)

// Create returns a database connection pool that can be used
// to begin transacting. When an error is not returned, that
// means the provided Options works
func Create(opt Options) (*sql.DB, error) {
	connectionParams := url.Values{}
	connectionParams.Add("parseTime", "true")
	for key, value := range opt.Params {
		connectionParams.Add(key, value)
	}
	params := connectionParams.Encode()
	dsnString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		opt.User,
		opt.Password,
		opt.Host,
		opt.Port,
		opt.Database,
		params,
	)
	connection, err := sql.Open("mysql", dsnString)
	if err != nil {
		return nil, err
	}
	if err := connection.Ping(); err != nil {
		return nil, err
	}
	return connection, nil
}
