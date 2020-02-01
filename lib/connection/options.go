package connection

type Options struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Params   map[string]string
}
