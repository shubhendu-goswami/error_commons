package interfaces

type Error interface {
	GetMessage() string
	GetCode() int64
}
