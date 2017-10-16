package errors

type Error interface {
	Error() string
	Root() error
	Stack() []string
}

type StackErr struct {
	text    string
	root    error
	stackEs []string
}

func (e *StackErr) Error() string {
	return e.text
}

func (e *StackErr) Root() error {
	return e.root
}

func (e *StackErr) Stack() []string {
	return e.stackEs
}
