package errors

type BadRequestError struct {
	Message string
}

func (bre BadRequestError) Error() string {
	return bre.Message
}
