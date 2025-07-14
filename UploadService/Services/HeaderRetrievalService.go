package services

import (
	"strconv"
	"uploadservice/Errors"
)

type HeaderValidatorService struct{}

func (HeaderValidatorService) CheckHeaderStringValue(headerKey string, headerValue string) (string, error) {
	if headerValue == "" {
		return "", errors.HeaderMissingError{HeaderName: headerKey}
	}

	return headerValue, nil
}

func (service HeaderValidatorService) CheckHeaderIntValue(headerKey string, headerValue string) (int, error) {

	stringValue, err := service.CheckHeaderStringValue(headerKey, headerValue)

	if err != nil {
		return -1, err
	}

	intValue, err := strconv.Atoi(stringValue)

	if err != nil {
		return -1, errors.HeaderIsNotIntError{HeaderName: headerKey}
	}

	return intValue, nil
}

func (service HeaderValidatorService) CheckHeaderInt64Value(headerKey string, headerValue string) (int64, error) {

	stringValue, err := service.CheckHeaderStringValue(headerKey, headerValue)

	if err != nil {
		return -1, err
	}

	intValue, err := strconv.ParseInt(stringValue, 10, 64)

	if err != nil {
		return -1, errors.HeaderIsNotIntError{HeaderName: headerKey}
	}

	return intValue, nil
}
