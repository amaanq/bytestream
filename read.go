package bytestream

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
)

type Reader struct {
	Reader *bytes.Reader
}

func (r *Reader) ReadBytes(length int) ([]byte, error) {
	_bytes := make([]byte, length)
	n, err := r.Reader.Read(_bytes)
	if err != nil {
		return nil, err
	}
	if n != length {
		return nil, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: " + fmt.Sprint(length))
	}
	return _bytes, nil
}

func (r *Reader) ReadBool() (bool, int8, error) {
	// A bool can be packed into a byte
	_byte, err := r.Reader.ReadByte()
	if err != nil {
		return false, 0, err
	}
	if _byte == 0x00 || _byte > 0x40 {
		return false, 0, nil
	}
	return true, int8(_byte), nil
}

func (r *Reader) ReadInt8() (int8, error) {
	// An int8 is effectively a byte
	_byte, err := r.Reader.ReadByte()
	if err != nil {
		return 0, err
	}
	return int8(_byte), nil
}

func (r *Reader) ReadUInt8() (uint8, error) {
	// A uint8 is also effectively a byte
	_byte, err := r.Reader.ReadByte()
	if err != nil {
		return 0, err
	}
	return uint8(_byte), nil
}

func (r *Reader) ReadInt16(endianness Endianness) (int16, error) {
	// An int16 is 2 bytes
	_bytes := make([]byte, 2)
	n, err := r.Reader.Read(_bytes)
	if err != nil {
		return 0, err
	}
	if n != Int16Size {
		return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 2")
	}

	switch endianness {
	case BigEndian:
		return int16(binary.BigEndian.Uint16(_bytes)), nil
	case LittleEndian:
		return int16(binary.LittleEndian.Uint16(_bytes)), nil
	default:
		return 0, fmt.Errorf("invalid endianness")
	}
}

func (r *Reader) ReadUInt16(endianness Endianness) (uint16, error) {
	// A uint16 is also 2 bytes
	_bytes := make([]byte, 2)
	n, err := r.Reader.Read(_bytes)
	if err != nil {
		return 0, err
	}
	if n != Int16Size {
		return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 2")
	}

	switch endianness {
	case BigEndian:
		return binary.BigEndian.Uint16(_bytes), nil
	case LittleEndian:
		return binary.LittleEndian.Uint16(_bytes), nil
	default:
		return 0, fmt.Errorf("invalid endianness")
	}
}

// We are using an int32 to represent an int24 since the stdlib doesn't provide a type for this. However, this int32 will only read 3 bytes and cannot go above the max size for an int24 (8388607 or 0x7FFFFF) :)
func (r *Reader) ReadInt24(endianness Endianness) (int32, error) {
	// An int24 is 3 bytes
	_bytes := make([]byte, 3)
	n, err := r.Reader.Read(_bytes)
	if err != nil {
		return 0, err
	}
	if n != Int24Size {
		return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 3")
	}

	switch endianness {
	case BigEndian:
		return int32(binary.BigEndian.Uint32(_bytes)), nil
	case LittleEndian:
		return int32(binary.LittleEndian.Uint32(_bytes)), nil
	default:
		return 0, fmt.Errorf("invalid endianness")
	}
}

// We are using a uint32 to represent a uint24 since the stdlib doesn't provide a type for this. However, this uint32 will only read 3 bytes and cannot go above the max size for a uint24 (16777215 or 0xFFFFFF) :)
func (r *Reader) ReadUInt24(endianness Endianness) (uint32, error) {
	// A uint24 is 3 bytes
	_bytes := make([]byte, 3)
	n, err := r.Reader.Read(_bytes)
	if err != nil {
		return 0, err
	}
	if n != Int24Size {
		return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 3")
	}

	switch endianness {
	case BigEndian:
		return binary.BigEndian.Uint32(_bytes), nil
	case LittleEndian:
		return binary.LittleEndian.Uint32(_bytes), nil
	default:
		return 0, fmt.Errorf("invalid endianness")
	}
}

func (r *Reader) ReadInt32(endianness Endianness) (int32, error) {
	// An int32 is 4 bytes
	_bytes := make([]byte, 4)
	n, err := r.Reader.Read(_bytes)
	if err != nil {
		return 0, err
	}
	if n != Int32Size {
		return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 4")
	}

	switch endianness {
	case BigEndian:
		return int32(binary.BigEndian.Uint32(_bytes)), nil
	case LittleEndian:
		return int32(binary.LittleEndian.Uint32(_bytes)), nil
	default:
		return 0, fmt.Errorf("invalid endianness")
	}
}

func (r *Reader) ReadUInt32(endianness Endianness) (uint32, error) {
	// A uint32 is 4 bytes
	_bytes := make([]byte, 4)
	n, err := r.Reader.Read(_bytes)
	if err != nil {
		return 0, err
	}
	if n != Int32Size {
		return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 4")
	}

	switch endianness {
	case BigEndian:
		return binary.BigEndian.Uint32(_bytes), nil
	case LittleEndian:
		return binary.LittleEndian.Uint32(_bytes), nil
	default:
		return 0, fmt.Errorf("invalid endianness")
	}
}

func (r *Reader) ReadInt64(endianness Endianness) (int64, error) {
	// An int64 is 8 bytes
	_bytes := make([]byte, 8)
	n, err := r.Reader.Read(_bytes)
	if err != nil {
		return 0, err
	}
	if n != Int64Size {
		return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 8")
	}

	switch endianness {
	case BigEndian:
		return int64(binary.BigEndian.Uint64(_bytes)), nil
	case LittleEndian:
		return int64(binary.LittleEndian.Uint64(_bytes)), nil
	default:
		return 0, fmt.Errorf("invalid endianness")
	}
}

func (r *Reader) ReadUInt64(endianness Endianness) (uint64, error) {
	// A uint64 is 8 bytes
	_bytes := make([]byte, 8)
	n, err := r.Reader.Read(_bytes)
	if err != nil {
		return 0, err
	}
	if n != Int64Size {
		return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 8")
	}

	switch endianness {
	case BigEndian:
		return binary.BigEndian.Uint64(_bytes), nil
	case LittleEndian:
		return binary.LittleEndian.Uint64(_bytes), nil
	default:
		return 0, fmt.Errorf("invalid endianness")
	}
}

func (r *Reader) ReadVarInt() (int64, error) {
	// A varint is a variable length integer.
	n, err := binary.ReadVarint(r.Reader)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (r *Reader) ReadUVarInt() (uint64, error) {
	// An unsigned varint is a variable length integer.
	n, err := binary.ReadUvarint(r.Reader)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (r *Reader) ReadLong(endianness Endianness) (int, error) {
	// A long is 4 bytes on a 32-bit machine and 8 bytes on a 64-bit machine.
	if Is64Bit {
		_bytes := make([]byte, 8)
		n, err := r.Reader.Read(_bytes)
		if err != nil {
			return 0, err
		}
		if n != Int64Size {
			return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 8")
		}

		switch endianness {
		case BigEndian:
			return int(binary.BigEndian.Uint64(_bytes)), nil
		case LittleEndian:
			return int(binary.LittleEndian.Uint64(_bytes)), nil
		default:
			return 0, fmt.Errorf("invalid endianness")
		}
	} else {
		_bytes := make([]byte, 4)
		n, err := r.Reader.Read(_bytes)
		if err != nil {
			return 0, err
		}
		if n != Int32Size {
			return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 4")
		}

		switch endianness {
		case BigEndian:
			return int(binary.BigEndian.Uint32(_bytes)), nil
		case LittleEndian:
			return int(binary.LittleEndian.Uint32(_bytes)), nil
		default:
			return 0, fmt.Errorf("invalid endianness")
		}
	}
}

func (r *Reader) ReadUnsignedLong(endianness Endianness) (uint, error) {
	// An unsigned long is 4 bytes on a 32-bit machine and 8 bytes on a 64-bit machine.
	if Is64Bit {
		_bytes := make([]byte, 8)
		n, err := r.Reader.Read(_bytes)
		if err != nil {
			return 0, err
		}
		if n != Int64Size {
			return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 8")
		}

		switch endianness {
		case BigEndian:
			return uint(binary.BigEndian.Uint64(_bytes)), nil
		case LittleEndian:
			return uint(binary.LittleEndian.Uint64(_bytes)), nil
		default:
			return 0, fmt.Errorf("invalid endianness")
		}
	} else {
		_bytes := make([]byte, 4)
		n, err := r.Reader.Read(_bytes)
		if err != nil {
			return 0, err
		}
		if n != Int32Size {
			return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 4")
		}

		switch endianness {
		case BigEndian:
			return uint(binary.BigEndian.Uint32(_bytes)), nil
		case LittleEndian:
			return uint(binary.LittleEndian.Uint32(_bytes)), nil
		default:
			return 0, fmt.Errorf("invalid endianness")
		}
	}
}

func (r *Reader) ReadLongLong(endianness Endianness) (int64, error) {
	// A long long is guaranteed to be 8 bytes
	_bytes := make([]byte, 8)
	n, err := r.Reader.Read(_bytes)
	if err != nil {
		return 0, err
	}
	if n != Int64Size {
		return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 8")
	}

	switch endianness {
	case BigEndian:
		return int64(binary.BigEndian.Uint64(_bytes)), nil
	case LittleEndian:
		return int64(binary.LittleEndian.Uint64(_bytes)), nil
	default:
		return 0, fmt.Errorf("invalid endianness")
	}
}

func (r *Reader) ReadUnsignedLongLong(endianness Endianness) (uint64, error) {
	// An unsigned long long is guaranteed to be 8 bytes
	_bytes := make([]byte, 8)
	n, err := r.Reader.Read(_bytes)
	if err != nil {
		return 0, err
	}
	if n != Int64Size {
		return 0, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 8")
	}

	switch endianness {
	case BigEndian:
		return binary.BigEndian.Uint64(_bytes), nil
	case LittleEndian:
		return binary.LittleEndian.Uint64(_bytes), nil
	default:
		return 0, fmt.Errorf("invalid endianness")
	}
}

func (r *Reader) ReadString() (string, error) {
	ssize_t, err := r.ReadInt32(BigEndian)
	if err != nil {
		return "", err
	}
	if ssize_t == -1 {
		return "", nil
	}
	if ssize_t < 0 {
		return "", fmt.Errorf("invalid string size: %d", ssize_t)
	}

	_bytes := make([]byte, ssize_t)
	n, err := r.Reader.Read(_bytes)
	if err != nil {
		return "", err
	}
	if n != int(ssize_t) {
		return "", fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: " + fmt.Sprint(ssize_t))
	}
	return string(_bytes), nil
}

func (r *Reader) ReadStringSize(ssize_t int) (string, error) {
	if ssize_t == -1 {
		return "", nil
	}
	if ssize_t < 0 {
		return "", fmt.Errorf("invalid string size: %d", ssize_t)
	}

	_bytes := make([]byte, ssize_t)
	n, err := r.Reader.Read(_bytes)
	if err != nil {
		return "", err
	}
	if n != int(ssize_t) {
		return "", fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: " + fmt.Sprint(ssize_t))
	}
	return string(_bytes), nil
}

func (r *Reader) ReadCompressedString() (string, error) {
	compressedLen, err := r.ReadInt32(BigEndian)
	if err != nil {
		return "", err
	}
	if compressedLen == -1 {
		return "", nil
	}
	if compressedLen < 0 {
		return "", fmt.Errorf("invalid string size: %d", compressedLen)
	}

	decompressedLen, err := r.ReadInt32(LittleEndian)
	if err != nil {
		return "", err
	}
	if decompressedLen < 0 {
		return "", fmt.Errorf("invalid string size: %d", decompressedLen)
	}

	compressedBytes := make([]byte, compressedLen)
	n, err := r.Reader.Read(compressedBytes)
	if err != nil {
		return "", err
	}
	if n != int(compressedLen) {
		return "", fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: " + fmt.Sprint(compressedLen))
	}

	zlibReader, err := zlib.NewReader(bytes.NewReader(compressedBytes))
	if err != nil {
		return "", err
	}
	defer zlibReader.Close()

	decompressedBytes := make([]byte, decompressedLen)
	n, err = zlibReader.Read(decompressedBytes)
	if err != nil {
		return "", err
	}
	if n != int(decompressedLen) {
		return "", fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: " + fmt.Sprint(decompressedLen))
	}
	return string(decompressedBytes), nil
}

// The first index of the slice is the low and high ID, and the second index is the tag.
func (r *Reader) ReadLogicLong(endianness Endianness) ([]string, error) {
	// A logic long is 8 bytes
	low_bytes := make([]byte, 4)
	high_bytes := make([]byte, 4)
	n, err := r.Reader.Read(low_bytes)
	if err != nil {
		return nil, err
	}
	if n != Int32Size {
		return nil, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 4")
	}
	n, err = r.Reader.Read(high_bytes)
	if err != nil {
		return nil, err
	}
	if n != Int32Size {
		return nil, fmt.Errorf("invalid number of bytes read! Read: " + fmt.Sprint(n) + " Expected: 4")
	}

	var low, high int
	switch endianness {
	case BigEndian:
		low = int(binary.BigEndian.Uint32(low_bytes))
		high = int(binary.BigEndian.Uint32(high_bytes))
		//return fmt.Sprintf("(%d, %d)", low, high), fmt.Sprintf("(%d, %d)", low, high), nil
	case LittleEndian:
		low = int(binary.LittleEndian.Uint32(low_bytes))
		high = int(binary.LittleEndian.Uint32(high_bytes))
		//return fmt.Sprintf("(%d, %d)", low, high), fmt.Sprintf("(%d, %d)", low, high), nil
	default:
		return nil, fmt.Errorf("invalid endianness")
	}
	tag, err := IDToTag(low, high)
	if err != nil {
		return nil, err
	}
	return []string{fmt.Sprintf("(%d, %d)", low, high), tag}, nil
}

// TODO allow reading any int with size from int8 to int64 (mostly allow 40/48/56...)
func (r *Reader) ReadUIntSize(size uint8, endianness Endianness) (int64, error)

func (r *Reader) ReadIntSize(size uint8, endianness Endianness) (int64, error)
