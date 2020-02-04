package common

const (
	ErrorOK         = "OK"
	ErrorInitFailed = "ERR_INIT_FAILED"
)
const (
	ExitCodeOK = iota
	ExitCodeGeneric
	ExitCodeDatabaseConnectionError
	ExitCodeCreateMigrationsTableFailed
	ExitCodeInvalidUserInput
	ExitCodeErrorCreatingFile
	ExitCodeInsufficientPermissions
	ExitCodeSavedYourAss
)
