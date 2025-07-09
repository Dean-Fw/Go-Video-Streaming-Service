package services

import "crypto/sha256"

type HashingService struct{}

func (HashingService) Hash(input string) string {
	hash := sha256.New()

	hash.Write([]byte(input))

	hashByteSlice := hash.Sum(nil)

	return string(hashByteSlice)
}
