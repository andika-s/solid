package example

// Abstraction
type IExample interface {
	Read() string
	Write(string)
}

type example struct {
	data string
}

// High level module

func New(value string) IExample {
	return &example{data: value}
}

// Low level module

// Single Resposibility
func (e *example) Read() string {
	return e.data
}

// Single Resposibility
func (e *example) Write(value string) {
	e.data = value
}
