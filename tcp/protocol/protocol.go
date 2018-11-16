package protocol

import (
	"encoding/binary"
	"bytes"
	"bufio"
)

// Encode :
func Encode(msg string) ([]byte, error) {
	length := int32(len(msg))
	// var pkg *bytes.Buffer = new(bytes.Buffer)
	pkg := new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode :
func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)

	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	pack := make([]byte, int32(4+length))
	_, err = reader.Read(pack)
	if err != nil{
		return "", err
	}
	return string(pack[4:]), nil
}	