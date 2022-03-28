package bytestream

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
)

type Writer struct {
	Buffer *bytes.Buffer
}

func New() {
	Writer := &Writer{
		Buffer: bytes.NewBuffer([]byte{}),
	}
	Writer.WriteCompressedString("hello there!")
	Writer.Buffer.Bytes()
}

func (w *Writer) WriteBytes(bytes []byte) error {
	n, err := w.Buffer.Write(bytes)
	if err != nil {
		return err
	}
	if n != len(bytes) {
		return fmt.Errorf("invalid number of bytes written! Wrote: " + fmt.Sprint(n) + " Expected: " + fmt.Sprint(len(bytes)))
	}
	return nil
}

func (w *Writer) WriteBool(data bool, count int8) error {
	if !data {
		return w.Buffer.WriteByte(0x00)
	}
	return w.Buffer.WriteByte(byte(count))
}

func (w *Writer) WriteInt8(data int8) error {
	// An int8 is effectively a byte
	return w.Buffer.WriteByte(byte(data))
}

func (w *Writer) WriteUInt8(data uint8) error {
	return w.Buffer.WriteByte(byte(data))
}

func (w *Writer) WriteInt16(data int16, endianness Endianness) error {
	switch endianness {
	case BigEndian:
		return binary.Write(w.Buffer, binary.BigEndian, data)
	case LittleEndian:
		return binary.Write(w.Buffer, binary.LittleEndian, data)
	default:
		return fmt.Errorf("invalid endianness")
	}
}

func (w *Writer) WriteUInt16(data uint16, endianness Endianness) error {
	switch endianness {
	case BigEndian:
		return binary.Write(w.Buffer, binary.BigEndian, data)
	case LittleEndian:
		return binary.Write(w.Buffer, binary.LittleEndian, data)
	default:
		return fmt.Errorf("invalid endianness")
	}
}

func (w *Writer) WriteInt24(data int32, endianness Endianness) error {
	if data > 0x7FFFFF {
		return fmt.Errorf("int24 overflow")
	}

	switch endianness {
	case BigEndian:
		w.Buffer.WriteByte(byte(data >> 16))
		w.Buffer.WriteByte(byte(data >> 8))
		w.Buffer.WriteByte(byte(data))
	case LittleEndian:
		w.Buffer.WriteByte(byte(data))
		w.Buffer.WriteByte(byte(data >> 8))
		w.Buffer.WriteByte(byte(data >> 16))
	default:
		return fmt.Errorf("invalid endianness")
	}
	return nil
}

func (w *Writer) WriteUInt24(data uint32, endianness Endianness) error {
	if data > 0xFFFFFF {
		return fmt.Errorf("uint24 overflow")
	}

	switch endianness {
	case BigEndian:
		w.Buffer.WriteByte(byte(data >> 16))
		w.Buffer.WriteByte(byte(data >> 8))
		w.Buffer.WriteByte(byte(data))
	case LittleEndian:
		w.Buffer.WriteByte(byte(data))
		w.Buffer.WriteByte(byte(data >> 8))
		w.Buffer.WriteByte(byte(data >> 16))
	default:
		return fmt.Errorf("invalid endianness")
	}
	return nil
}

func (w *Writer) WriteInt32(data int32, endianness Endianness) error {
	switch endianness {
	case BigEndian:
		return binary.Write(w.Buffer, binary.BigEndian, data)
	case LittleEndian:
		return binary.Write(w.Buffer, binary.LittleEndian, data)
	default:
		return fmt.Errorf("invalid endianness")
	}
}

func (w *Writer) WriteUInt32(data uint32, endianness Endianness) error {
	switch endianness {
	case BigEndian:
		return binary.Write(w.Buffer, binary.BigEndian, data)
	case LittleEndian:
		return binary.Write(w.Buffer, binary.LittleEndian, data)
	default:
		return fmt.Errorf("invalid endianness")
	}
}

func (w *Writer) WriteInt64(data int64, endianness Endianness) error {
	switch endianness {
	case BigEndian:
		return binary.Write(w.Buffer, binary.BigEndian, data)
	case LittleEndian:
		return binary.Write(w.Buffer, binary.LittleEndian, data)
	default:
		return fmt.Errorf("invalid endianness")
	}
}

func (w *Writer) WriteUInt64(data uint64, endianness Endianness) error {
	switch endianness {
	case BigEndian:
		return binary.Write(w.Buffer, binary.BigEndian, data)
	case LittleEndian:
		return binary.Write(w.Buffer, binary.LittleEndian, data)
	default:
		return fmt.Errorf("invalid endianness")
	}
}

func (w *Writer) WriteVarInt(data int64) error {
	ux := uint64(data) << 1
	if data < 0 {
		ux = ^ux
	}
	return w.WriteUVarInt(ux)
}

func (w *Writer) WriteUVarInt(data uint64) error {
	for data >= 0x80 {
		w.Buffer.WriteByte(byte(data) | 0x80)
		data >>= 7
	}
	return w.Buffer.WriteByte(byte(data))
}

func (w *Writer) WriteLong(data int64, endianness Endianness) error {
	// A long is 4 bytes on a 32-bit machine and 8 bytes on a 64-bit machine.
	if Is64Bit {
		return w.WriteInt64(data, endianness)
	} else {
		return w.WriteInt32(int32(data), endianness)
	}
}

func (w *Writer) WriteUnsignedLong(data uint64, endianness Endianness) error {
	// A long is 4 bytes on a 32-bit machine and 8 bytes on a 64-bit machine.
	if Is64Bit {
		return w.WriteUInt64(data, endianness)
	} else {
		return w.WriteUInt32(uint32(data), endianness)
	}
}

func (w *Writer) WriteLongLong(data int64, endianness Endianness) error {
	// A long long is guaranteed to be 8 bytes
	return w.WriteInt64(data, endianness)
}

func (w *Writer) WriteUnsignedLongLong(data uint64, endianness Endianness) error {
	// An unsigned long long is guaranteed to be 8 bytes
	return w.WriteUInt64(data, endianness)
}

func (w *Writer) WriteString(data string) error {
	length := len(data)
	err := w.WriteInt32(int32(length), BigEndian)
	if err != nil {
		return err
	}
	n, err := w.Buffer.WriteString(data)
	if err != nil {
		return err
	}
	if n != length {
		return fmt.Errorf("invalid number of bytes written! Wrote: " + fmt.Sprint(n) + " Expected: " + fmt.Sprint(length))
	}
	return nil
}

// This implementation writes the size of the string as a signed int of size bytesize, -1 will write 0xFFs for bytesize, and not write the string at all.
func (w *Writer) WriteStringSize(data string, bytesize int8) error {
	length := len(data)
	switch bytesize {
	case 1:
		if length > 0x7F {
			return fmt.Errorf("string size overflow")
		}
		err := w.WriteInt8(int8(length))
		if err != nil {
			return err
		}
	case 2:
		if length > 0x7FFF {
			return fmt.Errorf("string size overflow")
		}
		err := w.WriteInt16(int16(length), BigEndian)
		if err != nil {
			return err
		}
	case 3:
		if length > 0x7FFFFF {
			return fmt.Errorf("string size overflow")
		}
		err := w.WriteInt24(int32(length), BigEndian)
		if err != nil {
			return err
		}
	case 4:
		if length > 0x7FFFFFFF {
			return fmt.Errorf("string size overflow")
		}
		err := w.WriteInt32(int32(length), BigEndian)
		if err != nil {
			return err
		}
	case 8:
		if length > 0x7FFFFFFFFFFFFFFF {
			return fmt.Errorf("string size overflow")
		}
		err := w.WriteInt64(int64(length), BigEndian)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid string size")
	}
	n, err := w.Buffer.WriteString(data)
	if err != nil {
		return err
	}
	if n != length {
		return fmt.Errorf("invalid number of bytes written! Wrote: " + fmt.Sprint(n) + " Expected: " + fmt.Sprint(length))
	}
	return nil
}

func (w *Writer) WriteCompressedString(data string) error {
	decompressedLength := len(data)
	intermediateBuffer := bytes.NewBuffer([]byte{})
	zlibWriter := zlib.NewWriter(intermediateBuffer)
	n, err := zlibWriter.Write([]byte(data))
	if err != nil {
		return err
	}
	err = zlibWriter.Close()
	if err != nil {
		return err
	}
	compressedBytes := intermediateBuffer.Bytes()
	compressedLength := len(compressedBytes)
	err = w.WriteInt32(int32(compressedLength), BigEndian) // pack compressed size as BE
	if err != nil {
		return err
	}
	err = w.WriteInt32(int32(decompressedLength), LittleEndian) // pack uncompressed size as LE
	if err != nil {
		return err
	}
	w.Buffer.Write(compressedBytes) // write compressed data
	if n != compressedLength {
		return fmt.Errorf("invalid number of bytes written! Wrote: " + fmt.Sprint(n) + " Expected: " + fmt.Sprint(compressedLength))
	}
	return nil
}

func (w *Writer) WriteLogicLong(tag string) error {
	Low, High := TagToID(tag)
	err := w.WriteInt32(int32(Low), BigEndian)
	if err != nil {
		return err
	}
	return w.WriteInt32(int32(High), BigEndian)
}

// TODO allow writing any int with size from int8 to int64 (mostly allow 40/48/56...)
func (w *Writer) WriteUIntSize(data int64, size uint8, endianness Endianness) error

func (w *Writer) WriteIntSize(data int64, size uint8, endianness Endianness) error
