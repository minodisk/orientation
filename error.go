package orientation

type DecodeError struct {
	error
}

func (err DecodeError) Cause() error {
	return err.error
}

type FormatError struct {
	error
}

func (err FormatError) Cause() error {
	return err.error
}

type TagError struct {
	error
}

func (err TagError) Cause() error {
	return err.error
}

type OrientError struct {
	error
}

func (err OrientError) Cause() error {
	return err.error
}
