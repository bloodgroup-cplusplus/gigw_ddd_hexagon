package ports

type DbPort interface {
	CloseDbConnection()
	AddToHistory(answer int32, operation string) error
	// we will store the answer and operation in the database
}
