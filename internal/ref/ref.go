package ref

// Reference returns a reference to the given value.
func Reference[T any](x T) *T {
	return &x
}

func SetFromSlice[T comparable](x []T) map[T]struct{} {
	set := map[T]struct{}{}
	for _, v := range x {
		set[v] = struct{}{}
	}
	return set
}
