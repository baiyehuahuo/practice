// 缓存值的抽象与封装
package GeeCache

// A ByteView holds an immutable view of bytes
type ByteView struct {
	b []byte
}

// Len returns the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the adta as a byte slice
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String returns the data as a string, making a copy if necessary
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	ans := make([]byte, len(b))
	copy(ans, b)
	return ans
}
