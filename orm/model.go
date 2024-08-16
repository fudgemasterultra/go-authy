package orm

type Response[T any]  struct {
	Success bool
	Message string
	Data *T
}

type ConnectedDb struct {
	Url string
}