package interfaces

type IFileSystemService interface {
	CreateNewFile(fileNameHash string, content []byte) error
}
