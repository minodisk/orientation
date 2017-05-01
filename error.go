package orientation

// DecodeError is returned by Apply and Decode when image.Decode returns an error.
type DecodeError struct {
	error
}

// Cause returns the underlying cause of DecodeError.
func (err DecodeError) Cause() error {
	return err.error
}

// FormatError is returned by Apply and Decode when the format of decoded image is not jpeg.
type FormatError struct {
	error
}

// Cause returns the underlying cause of FormatError.
func (err FormatError) Cause() error {
	return err.error
}

// TagError is returned by Apply and Tag when the EXIF can not be decoded, the orientation tag does not exits,
// or the orientation tag is not int.
type TagError struct {
	error
}

// Cause returns the underlying cause of TagError.
func (err TagError) Cause() error {
	return err.error
}

// OrientError is returned by Apply and Tag when the specified orientation tag is unknown.
type OrientError struct {
	error
}

// Cause returns the underlying cause of OrientError.
func (err OrientError) Cause() error {
	return err.error
}
