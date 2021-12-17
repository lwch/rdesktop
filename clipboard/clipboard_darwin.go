package clipboard

// Set set text to clipboard
func Set(text string) error {
	return ErrUnsupport
}

// Get get clipboard text
func Get() (string, error) {
	return "", ErrUnsupport
}
