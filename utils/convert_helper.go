package utils

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func ConvertBool(v bool) *bool {
	return &v
}

// Int32 is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it.
func ConvertInt32(v int32) *int32 {
	return &v
}

// Int is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it, but unlike Int32
// its argument value is an int.
func ConvertInt(v int) *int32 {
	p := new(int32)
	*p = int32(v)
	return p
}

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func ConvertInt64(v int64) *int64 {
	return &v
}

// Float32 is a helper routine that allocates a new float32 value
// to store v and returns a pointer to it.
func ConvertFloat32(v float32) *float32 {
	return &v
}

// Float64 is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func ConvertFloat64(v float64) *float64 {
	return &v
}

// Uint32 is a helper routine that allocates a new uint32 value
// to store v and returns a pointer to it.
func ConvertUint32(v uint32) *uint32 {
	return &v
}

// Uint64 is a helper routine that allocates a new uint64 value
// to store v and returns a pointer to it.
func ConvertUint64(v uint64) *uint64 {
	return &v
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func ConvertString(v string) *string {
	return &v
}
