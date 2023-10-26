package error_ext

type WrappedError struct {
	Msg string
	Err error
}

func (e WrappedError) Error() string {
	return e.Msg + e.Err.Error()
}

func (e WrappedError) Unwrap() error {
	return e.Err
}

// WrapError simplifies the common task of taking an error and prefixing it with some helpful context, like the name of the calling function
func WrapError(msg string, err error) error {
	if err == nil {
		return nil
	}
	return WrappedError{
		Msg: msg,
		Err: err,
	}
}
