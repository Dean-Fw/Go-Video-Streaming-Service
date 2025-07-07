package serviceerrors

type NotFoundError struct {
	Message string
}

func (nf NotFoundError) Error() string {
	return nf.Message
}
