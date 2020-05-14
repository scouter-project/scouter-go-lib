package util

// CopyArray returns copied array
func CopyArray(data []byte, pos int32, length int32) []byte {
	return data[pos : pos+length]
}
