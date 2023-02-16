package helper

func Ptr[T any](v T) *T {
	return &v
}

func SliceConverter[T any, F any](slice []T, converter func(T) F) []F {
	result := make([]F, len(slice))
	for i, v := range slice {
		result[i] = converter(v)
	}
	return result
}
