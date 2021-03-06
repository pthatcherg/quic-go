package utils

import (
	"bytes"
	"fmt"
	"io"
)

// LittleEndian is the little-endian implementation of ByteOrder.
var LittleEndian ByteOrder = littleEndian{}

type littleEndian struct{}

var _ ByteOrder = &littleEndian{}

func (littleEndian) WriteUintN(b *bytes.Buffer, n uint64, length uint8) {
	bytes := make([]byte, length)
	shifter := uint64(0)
	for i := uint8(0); i < length; i++ {
		bytes[i] = byte(n >> shifter)
		shifter += 8
	}
	b.Write(bytes)
}

// ReadUintN reads N bytes
func (littleEndian) ReadUintN(b io.ByteReader, length uint8) (uint64, error) {
	var res uint64
	for i := uint8(0); i < length; i++ {
		bt, err := b.ReadByte()
		if err != nil {
			return 0, err
		}
		res ^= uint64(bt) << (i * 8)
	}
	return res, nil
}

// ReadUint64 reads a uint64
func (littleEndian) ReadUint64(b io.ByteReader) (uint64, error) {
	var b1, b2, b3, b4, b5, b6, b7, b8 uint8
	var err error
	if b1, err = b.ReadByte(); err != nil {
		return 0, err
	}
	if b2, err = b.ReadByte(); err != nil {
		return 0, err
	}
	if b3, err = b.ReadByte(); err != nil {
		return 0, err
	}
	if b4, err = b.ReadByte(); err != nil {
		return 0, err
	}
	if b5, err = b.ReadByte(); err != nil {
		return 0, err
	}
	if b6, err = b.ReadByte(); err != nil {
		return 0, err
	}
	if b7, err = b.ReadByte(); err != nil {
		return 0, err
	}
	if b8, err = b.ReadByte(); err != nil {
		return 0, err
	}
	return uint64(b1) + uint64(b2)<<8 + uint64(b3)<<16 + uint64(b4)<<24 + uint64(b5)<<32 + uint64(b6)<<40 + uint64(b7)<<48 + uint64(b8)<<56, nil
}

// ReadUint32 reads a uint32
func (littleEndian) ReadUint32(b io.ByteReader) (uint32, error) {
	var b1, b2, b3, b4 uint8
	var err error
	if b1, err = b.ReadByte(); err != nil {
		return 0, err
	}
	if b2, err = b.ReadByte(); err != nil {
		return 0, err
	}
	if b3, err = b.ReadByte(); err != nil {
		return 0, err
	}
	if b4, err = b.ReadByte(); err != nil {
		return 0, err
	}
	return uint32(b1) + uint32(b2)<<8 + uint32(b3)<<16 + uint32(b4)<<24, nil
}

// ReadUint16 reads a uint16
func (littleEndian) ReadUint16(b io.ByteReader) (uint16, error) {
	var b1, b2 uint8
	var err error
	if b1, err = b.ReadByte(); err != nil {
		return 0, err
	}
	if b2, err = b.ReadByte(); err != nil {
		return 0, err
	}
	return uint16(b1) + uint16(b2)<<8, nil
}

// WriteUint64 writes a uint64
func (littleEndian) WriteUint64(b *bytes.Buffer, i uint64) {
	b.Write([]byte{
		uint8(i), uint8(i >> 8), uint8(i >> 16), uint8(i >> 24),
		uint8(i >> 32), uint8(i >> 40), uint8(i >> 48), uint8(i >> 56),
	})
}

// WriteUint56 writes 56 bit of a uint64
func (littleEndian) WriteUint56(b *bytes.Buffer, i uint64) {
	if i >= (1 << 56) {
		panic(fmt.Sprintf("%#x doesn't fit into 56 bits", i))
	}
	b.Write([]byte{
		uint8(i), uint8(i >> 8), uint8(i >> 16), uint8(i >> 24),
		uint8(i >> 32), uint8(i >> 40), uint8(i >> 48),
	})
}

// WriteUint48 writes 48 bit of a uint64
func (littleEndian) WriteUint48(b *bytes.Buffer, i uint64) {
	if i >= (1 << 48) {
		panic(fmt.Sprintf("%#x doesn't fit into 48 bits", i))
	}
	b.Write([]byte{
		uint8(i), uint8(i >> 8), uint8(i >> 16), uint8(i >> 24),
		uint8(i >> 32), uint8(i >> 40),
	})
}

// WriteUint40 writes 40 bit of a uint64
func (littleEndian) WriteUint40(b *bytes.Buffer, i uint64) {
	if i >= (1 << 40) {
		panic(fmt.Sprintf("%#x doesn't fit into 40 bits", i))
	}
	b.Write([]byte{
		uint8(i), uint8(i >> 8), uint8(i >> 16),
		uint8(i >> 24), uint8(i >> 32),
	})
}

// WriteUint32 writes a uint32
func (littleEndian) WriteUint32(b *bytes.Buffer, i uint32) {
	b.Write([]byte{uint8(i), uint8(i >> 8), uint8(i >> 16), uint8(i >> 24)})
}

// WriteUint24 writes 24 bit of a uint32
func (littleEndian) WriteUint24(b *bytes.Buffer, i uint32) {
	if i >= (1 << 24) {
		panic(fmt.Sprintf("%#x doesn't fit into 24 bits", i))
	}
	b.Write([]byte{uint8(i), uint8(i >> 8), uint8(i >> 16)})
}

// WriteUint16 writes a uint16
func (littleEndian) WriteUint16(b *bytes.Buffer, i uint16) {
	b.Write([]byte{uint8(i), uint8(i >> 8)})
}

func (l littleEndian) ReadUfloat16(b io.ByteReader) (uint64, error) {
	return readUfloat16(b, l)
}

func (l littleEndian) WriteUfloat16(b *bytes.Buffer, val uint64) {
	writeUfloat16(b, l, val)
}
