package errors

const _split = "\n"

func Wrap(text string, err error) *StackErr {
	switch v := err.(type) {
	case *StackErr:
		stackEs := append(v.stackEs, text)
		return &StackErr{text: text, root: v.root, stackEs: stackEs}
	default:
		return &StackErr{text: text, root: v, stackEs: []string{text}}
	}
}

func StackTrace(err error) string {
	if err == nil {
		return ""
	}
	switch v := err.(type) {
	case *StackErr:
		if len(v.stackEs) == 0 {
			return ""
		}
		res := ""
		for i := len(v.stackEs) - 1; i >= 0; i-- {
			res += v.stackEs[i] + _split
		}
		return res
	default:
		return err.Error()
	}
}
