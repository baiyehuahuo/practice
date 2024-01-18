package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	var length = int32(len(message))
	var pkg = &bytes.Buffer{}

	// 读取消息长度，转换成 int32 占四个字节 写入消息头
	if err := binary.Write(pkg, binary.LittleEndian, length); err != nil {
		return nil, err
	}

	if err := binary.Write(pkg, binary.LittleEndian, []byte(message)); err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息长度
	lengthByte, err := reader.Peek(4) // 读取前四个字节的数据
	if err != nil {
		return "", err
	}
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	if err = binary.Read(lengthBuff, binary.LittleEndian, &length); err != nil {
		return "", err
	}

	// 返回缓冲中现有的可读取的字节数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	pack := make([]byte, int(4+length))
	if _, err = reader.Read(pack); err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
