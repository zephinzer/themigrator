package connection

import (
	"database/sql"
	"fmt"
	"net/url"
)

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
