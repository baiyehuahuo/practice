package sync_pool_reuse

import (
	"bytes"
	"encoding/json"
	"testing"
)

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var bufT bytes.Buffer
		bufT.Write(data)
	}
}

func BenchmarkBufferWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bufT := bufferPool.Get().(*bytes.Buffer)
		bufT.Write(data)
		bufT.Reset()
		bufferPool.Put(bufT)
	}
}
