package interfaces

type IHeaderValidatorService interface {
	CheckHeaderStringValue(headerKey string, headerValue string) (string, error)
	CheckHeaderIntValue(headerKey string, headerValue string) (int, error)
}
