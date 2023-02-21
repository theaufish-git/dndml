package ref

// Reference returns a reference to the given value.
func Reference[T any](x T) *T {
	return &x
}
