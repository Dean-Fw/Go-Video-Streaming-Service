package handlererrors

type RangeHeaderMalformed struct {
	Message string
}

func (rhm RangeHeaderMalformed) Error() string {
	return rhm.Message
}
