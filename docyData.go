package main

import (
	"fmt"
	"time"
)

// docy相关数据结构
type DocyData struct {
	version   int16
	protoType int16
	bodyLen   int
	body      []byte
}

func NewDocyData() *DocyData {
	return &DocyData{}
}

// 数据单元的处理函数
func (this *DocyData) GetVersion() int16 {
	return this.version
}

func (this *DocyData) SetVersion(version int16) {
	this.version = version
}

func (this *DocyData) GetProtoType() int16 {
	return this.protoType
}

func (this *DocyData) SetProtoType(protoType int16) {
	this.protoType = protoType
}

func (this *DocyData) GetBody() []byte {
	return this.body
}

func (this *DocyData) SetBody(body []byte) {
	this.body = body
}

// 读取协议包解析
func (this *DocyData) Parser(tcpConnection *TcpConnection, timeout time.Duration) error {
	head, err := tcpConnection.Read(8, timeout)
	if err != nil {
		return err
	}
	this.version = StreamToInt16(head[0:2], BigEndian)
	this.protoType = StreamToInt16(head[2:4], BigEndian)
	this.bodyLen = StreamToInt32(head[4:8], BigEndian)
	this.body = make([]byte, this.bodyLen)
	if body, err := tcpConnection.Read(this.bodyLen, timeout); err != nil {
		return err
	} else {
		copy(this.body, body)
	}

	return nil
}

func (this *DocyData) Write(tcpConnection *TcpConnection, timeout time.Duration) {
	tcpConnection.SetWriteDeadline(time.Now().Add(timeout))
	return tcpConnection.Write(this.ConvertToStream())
}

func (this *DocyData) ConvertToStream() []byte {
	data := make([]byte, this.bodyLen+8)
	copy(data[0:2], Int16ToStream(this.version, BigEndian))
	copy(data[2:4], Int16ToStream(this.protoType, BigEndian))
	copy(data[4:8], Int32ToStream(this.bodyLen, BigEndian))
	copy(data[8:this.bodyLen+8], this.body)
	return data
}

const (
	BigEndian ByteOrder = iota
	LittleEndian
)

func StreamToInt16(stream []byte, byteOrder ByteOrder) int16 {
	if len(stream) != 2 {
		return 0
	}
	var u uint16
	if byteOrder == LittleEndian {
		for i := 0; i < 2; i++ {
			u += uint16(stream[i]) << uint(i*8)
		}
	} else {
		for i := 0; i < 2; i++ {
			u += uint16(stream[i]) << uint(8*(1-i))
		}
	}
	return int16(u)
}

func Int16ToStream(i int16, byteOrder ByteOrder) []byte {
	u := uint16(i)
	stream := [2]byte{0, 0}
	if byteOrder == LittleEndian {
		for i := 0; i < 2; i++ {
			stream[i] = byte(u >> uint(8*i))
		}
	} else {
		for i := 0; i < 2; i++ {
			stream[i] = byte(u >> uint(8*(1-i)))
		}
	}
	return stream[:]
}

func StreamToInt32(stream []byte, byteOrder ByteOrder) int32 {
	if len(stream) != 4 {
		return 0
	}
	var u uint32
	if byteOrder == LittleEndian {
		for i := 0; i < 4; i++ {
			u += uint32(stream[i]) << uint(i*8)
		}
	} else {
		for i := 0; i < 4; i++ {
			u += uint32(stream[i]) << uint(8*(3-i))
		}
	}
	return int32(u)
}

func Int32ToStream(i int32, byteOrder ByteOrder) []byte {
	u := uint32(i)
	stream := [4]byte{0, 0, 0, 0}
	if byteOrder == LittleEndian {
		for i := 0; i < 4; i++ {
			stream[i] = byte(u >> uint(8*i))
		}
	} else {
		for i := 0; i < 4; i++ {
			stream[i] = byte(u >> uint(8*(3-i)))
		}
	}
	return stream[:]
}
