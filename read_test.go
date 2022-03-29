package bytestream

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReader_ReadBytes(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "nil", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{length: 1}, want: nil, wantErr: true},
		{name: "one", fields: fields{Reader: bytes.NewReader([]byte{0x00})}, args: args{length: 1}, want: []byte{0x00}, wantErr: false},
		{name: "ten", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09})}, args: args{length: 5}, want: []byte{0x00, 0x01, 0x02, 0x03, 0x04}, wantErr: false},
		{name: "ten", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09})}, args: args{length: 11}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadBytes(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reader.ReadBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadBool(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		want1   int8
		wantErr bool
	}{
		{name: "nil", fields: fields{Reader: bytes.NewReader([]byte{})}, want: false, want1: 0, wantErr: true},
		{name: "false", fields: fields{Reader: bytes.NewReader([]byte{0x00})}, want: false, want1: 0, wantErr: false},
		{name: "true", fields: fields{Reader: bytes.NewReader([]byte{0x01})}, want: true, want1: 1, wantErr: false},
		{name: "true x2", fields: fields{Reader: bytes.NewReader([]byte{0x02})}, want: true, want1: 2, wantErr: false},
		{name: "true x3", fields: fields{Reader: bytes.NewReader([]byte{0x03})}, want: true, want1: 3, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, got1, err := r.ReadBool()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadBool() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Reader.ReadBool() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReader_ReadInt8(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    int8
		wantErr bool
	}{
		{name: "nil", fields: fields{Reader: bytes.NewReader([]byte{})}, want: 0, wantErr: true},
		{name: "zero", fields: fields{Reader: bytes.NewReader([]byte{0x00})}, want: 0, wantErr: false},
		{name: "one", fields: fields{Reader: bytes.NewReader([]byte{0x01})}, want: 1, wantErr: false},
		{name: "minus one", fields: fields{Reader: bytes.NewReader([]byte{0xFF})}, want: -1, wantErr: false},
		{name: "max", fields: fields{Reader: bytes.NewReader([]byte{0x7F})}, want: 127, wantErr: false},
		{name: "max minus", fields: fields{Reader: bytes.NewReader([]byte{0x80})}, want: -128, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadInt8()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadUInt8(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint8
		wantErr bool
	}{
		{name: "nil", fields: fields{Reader: bytes.NewReader([]byte{})}, want: 0, wantErr: true},
		{name: "zero", fields: fields{Reader: bytes.NewReader([]byte{0x00})}, want: 0, wantErr: false},
		{name: "one", fields: fields{Reader: bytes.NewReader([]byte{0x01})}, want: 1, wantErr: false},
		{name: "max", fields: fields{Reader: bytes.NewReader([]byte{0xFF})}, want: 255, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadUInt8()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadUInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadUInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadInt16(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	type args struct {
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int16
		wantErr bool
	}{
		{name: "nil BE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: BigEndian}, want: 0, wantErr: true},
		{name: "zero BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00})}, args: args{endianness: BigEndian}, want: 0, wantErr: false},
		{name: "one BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x01})}, args: args{endianness: BigEndian}, want: 1, wantErr: false},
		{name: "minus one BE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: -1, wantErr: false},
		{name: "max BE", fields: fields{Reader: bytes.NewReader([]byte{0x7F, 0xFF})}, args: args{endianness: BigEndian}, want: 32767, wantErr: false},
		{name: "max minus BE", fields: fields{Reader: bytes.NewReader([]byte{0x80, 0x00})}, args: args{endianness: BigEndian}, want: -32768, wantErr: false},

		{name: "nil LE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: LittleEndian}, want: 0, wantErr: true},
		{name: "zero LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 0, wantErr: false},
		{name: "one LE", fields: fields{Reader: bytes.NewReader([]byte{0x01, 0x00})}, args: args{endianness: LittleEndian}, want: 1, wantErr: false},
		{name: "minus one LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF})}, args: args{endianness: LittleEndian}, want: -1, wantErr: false},
		{name: "max LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0x7F})}, args: args{endianness: LittleEndian}, want: 32767, wantErr: false},
		{name: "max minus LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x80})}, args: args{endianness: LittleEndian}, want: -32768, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadInt16(tt.args.endianness)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadUInt16(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	type args struct {
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint16
		wantErr bool
	}{
		{name: "nil BE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: BigEndian}, want: 0, wantErr: true},
		{name: "zero BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00})}, args: args{endianness: BigEndian}, want: 0, wantErr: false},
		{name: "one BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x01})}, args: args{endianness: BigEndian}, want: 1, wantErr: false},
		{name: "max BE", fields: fields{Reader: bytes.NewReader([]byte{0x7F, 0xFF})}, args: args{endianness: BigEndian}, want: 32767, wantErr: false},

		{name: "nil LE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: LittleEndian}, want: 0, wantErr: true},
		{name: "zero LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 0, wantErr: false},
		{name: "one LE", fields: fields{Reader: bytes.NewReader([]byte{0x01, 0x00})}, args: args{endianness: LittleEndian}, want: 1, wantErr: false},
		{name: "max LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0x7F})}, args: args{endianness: LittleEndian}, want: 32767, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadUInt16(tt.args.endianness)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadUInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadUInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadInt24(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	type args struct {
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int32
		wantErr bool
	}{
		{name: "nil BE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: BigEndian}, want: 0, wantErr: true},
		{name: "zero BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: 0, wantErr: false},
		{name: "one BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x01})}, args: args{endianness: BigEndian}, want: 1, wantErr: false},
		{name: "minus one BE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: -1, wantErr: false},
		{name: "max BE", fields: fields{Reader: bytes.NewReader([]byte{0x7F, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: 8388607, wantErr: false},
		{name: "max minus BE", fields: fields{Reader: bytes.NewReader([]byte{0x80, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: -8388608, wantErr: false},

		{name: "nil LE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: LittleEndian}, want: 0, wantErr: true},
		{name: "zero LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 0, wantErr: false},
		{name: "one LE", fields: fields{Reader: bytes.NewReader([]byte{0x01, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 1, wantErr: false},
		{name: "minus one LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF})}, args: args{endianness: LittleEndian}, want: -1, wantErr: false},
		{name: "max LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0x7F})}, args: args{endianness: LittleEndian}, want: 8388607, wantErr: false},
		{name: "max minus LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x80, 0x00})}, args: args{endianness: LittleEndian}, want: -8388608, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadInt24(tt.args.endianness)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadInt24() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadInt24() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadUInt24(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	type args struct {
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint32
		wantErr bool
	}{
		{name: "nil BE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: BigEndian}, want: 0, wantErr: true},
		{name: "zero BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: 0, wantErr: false},
		{name: "one BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x01})}, args: args{endianness: BigEndian}, want: 1, wantErr: false},
		{name: "max BE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: 16777215, wantErr: false},

		{name: "nil LE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: LittleEndian}, want: 0, wantErr: true},
		{name: "zero LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 0, wantErr: false},
		{name: "one LE", fields: fields{Reader: bytes.NewReader([]byte{0x01, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 1, wantErr: false},
		{name: "max LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF})}, args: args{endianness: LittleEndian}, want: 16777215, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadUInt24(tt.args.endianness)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadUInt24() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadUInt24() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadInt32(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	type args struct {
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int32
		wantErr bool
	}{
		{name: "nil BE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: BigEndian}, want: 0, wantErr: true},
		{name: "zero BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: 0, wantErr: false},
		{name: "one BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x01})}, args: args{endianness: BigEndian}, want: 1, wantErr: false},
		{name: "minus one BE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: -1, wantErr: false},
		{name: "max BE", fields: fields{Reader: bytes.NewReader([]byte{0x7F, 0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: 2147483647, wantErr: false},
		{name: "max minus BE", fields: fields{Reader: bytes.NewReader([]byte{0x80, 0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: -2147483648, wantErr: false},

		{name: "nil LE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: LittleEndian}, want: 0, wantErr: true},
		{name: "zero LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 0, wantErr: false},
		{name: "one LE", fields: fields{Reader: bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 1, wantErr: false},
		{name: "minus one LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: LittleEndian}, want: -1, wantErr: false},
		{name: "max LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0x7F})}, args: args{endianness: LittleEndian}, want: 2147483647, wantErr: false},
		{name: "max minus LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x80})}, args: args{endianness: LittleEndian}, want: -2147483648, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadInt32(tt.args.endianness)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadUInt32(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	type args struct {
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint32
		wantErr bool
	}{
		{name: "nil BE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: BigEndian}, want: 0, wantErr: true},
		{name: "zero BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: 0, wantErr: false},
		{name: "one BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x01})}, args: args{endianness: BigEndian}, want: 1, wantErr: false},
		{name: "max BE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: 4294967295, wantErr: false},

		{name: "nil LE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: LittleEndian}, want: 0, wantErr: true},
		{name: "zero LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 0, wantErr: false},
		{name: "one LE", fields: fields{Reader: bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 1, wantErr: false},
		{name: "max LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: LittleEndian}, want: 4294967295, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadUInt32(tt.args.endianness)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadUInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadUInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadInt64(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	type args struct {
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{name: "nil BE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: BigEndian}, want: 0, wantErr: true},
		{name: "zero BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: 0, wantErr: false},
		{name: "one BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})}, args: args{endianness: BigEndian}, want: 1, wantErr: false},
		{name: "minus one BE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: -1, wantErr: false},
		{name: "max BE", fields: fields{Reader: bytes.NewReader([]byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: 9223372036854775807, wantErr: false},
		{name: "max minus BE", fields: fields{Reader: bytes.NewReader([]byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: -9223372036854775808, wantErr: false},

		{name: "nil LE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: LittleEndian}, want: 0, wantErr: true},
		{name: "zero LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 0, wantErr: false},
		{name: "one LE", fields: fields{Reader: bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 1, wantErr: false},
		{name: "minus one LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: LittleEndian}, want: -1, wantErr: false},
		{name: "max LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F})}, args: args{endianness: LittleEndian}, want: 9223372036854775807, wantErr: false},
		{name: "max minus LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80})}, args: args{endianness: LittleEndian}, want: -9223372036854775808, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadInt64(tt.args.endianness)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadUInt64(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	type args struct {
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		{name: "nil BE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: BigEndian}, want: 0, wantErr: true},
		{name: "zero BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: 0, wantErr: false},
		{name: "one BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})}, args: args{endianness: BigEndian}, want: 1, wantErr: false},
		{name: "max BE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: 18446744073709551615, wantErr: false},

		{name: "nil LE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: LittleEndian}, want: 0, wantErr: true},
		{name: "zero LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 0, wantErr: false},
		{name: "one LE", fields: fields{Reader: bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 1, wantErr: false},
		{name: "max LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: LittleEndian}, want: 18446744073709551615, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadUInt64(tt.args.endianness)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadUInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadUInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadVarInt(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		{name: "nil", fields: fields{Reader: bytes.NewReader([]byte{})}, want: 0, wantErr: true},
		{name: "zero", fields: fields{Reader: bytes.NewReader([]byte{0x00})}, want: 0, wantErr: false},
		{name: "one", fields: fields{Reader: bytes.NewReader([]byte{0x02})}, want: 1, wantErr: false},
		{name: "max", fields: fields{Reader: bytes.NewReader([]byte{0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x01})}, want: 9223372036854775807, wantErr: false},
		{name: "minus max", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x01})}, want: -9223372036854775808, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadVarInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadVarInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadVarInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadUVarInt(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		{name: "nil", fields: fields{Reader: bytes.NewReader([]byte{})}, want: 0, wantErr: true},
		{name: "zero", fields: fields{Reader: bytes.NewReader([]byte{0x00})}, want: 0, wantErr: false},
		{name: "one", fields: fields{Reader: bytes.NewReader([]byte{0x01})}, want: 1, wantErr: false},
		{name: "max", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x01})}, want: 18446744073709551615, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadUVarInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadUVarInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadUVarInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Assumes 64-bit system
func TestReader_ReadLong(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	type args struct {
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{name: "nil BE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: BigEndian}, want: 0, wantErr: true},
		{name: "zero BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: 0, wantErr: false},
		{name: "one BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})}, args: args{endianness: BigEndian}, want: 1, wantErr: false},
		{name: "minus one BE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: -1, wantErr: false},
		{name: "max BE", fields: fields{Reader: bytes.NewReader([]byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: 9223372036854775807, wantErr: false},
		{name: "minus max BE", fields: fields{Reader: bytes.NewReader([]byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: -9223372036854775808, wantErr: false},

		{name: "nil LE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: LittleEndian}, want: 0, wantErr: true},
		{name: "zero LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 0, wantErr: false},
		{name: "one LE", fields: fields{Reader: bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 1, wantErr: false},
		{name: "minus one LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: LittleEndian}, want: -1, wantErr: false},
		{name: "max LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F})}, args: args{endianness: LittleEndian}, want: 9223372036854775807, wantErr: false},
		{name: "minus max LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80})}, args: args{endianness: LittleEndian}, want: -9223372036854775808, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadLong(tt.args.endianness)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadLong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Assumes 64-bit system
func TestReader_ReadUnsignedLong(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	type args struct {
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr bool
	}{
		{name: "nil BE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: BigEndian}, want: 0, wantErr: true},
		{name: "zero BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: 0, wantErr: false},
		{name: "one BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})}, args: args{endianness: BigEndian}, want: 1, wantErr: false},
		{name: "max BE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: 18446744073709551615, wantErr: false},

		{name: "nil LE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: LittleEndian}, want: 0, wantErr: true},
		{name: "zero LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 0, wantErr: false},
		{name: "one LE", fields: fields{Reader: bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 1, wantErr: false},
		{name: "max LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: LittleEndian}, want: 18446744073709551615, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadUnsignedLong(tt.args.endianness)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadUnsignedLong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadUnsignedLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadLongLong(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	type args struct {
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{name: "nil BE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: BigEndian}, want: 0, wantErr: true},
		{name: "zero BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: 0, wantErr: false},
		{name: "one BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})}, args: args{endianness: BigEndian}, want: 1, wantErr: false},
		{name: "minus one BE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: -1, wantErr: false},
		{name: "max BE", fields: fields{Reader: bytes.NewReader([]byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: 9223372036854775807, wantErr: false},
		{name: "minus max BE", fields: fields{Reader: bytes.NewReader([]byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: -9223372036854775808, wantErr: false},

		{name: "nil LE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: LittleEndian}, want: 0, wantErr: true},
		{name: "zero LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 0, wantErr: false},
		{name: "one LE", fields: fields{Reader: bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 1, wantErr: false},
		{name: "minus one LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: LittleEndian}, want: -1, wantErr: false},
		{name: "max LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F})}, args: args{endianness: LittleEndian}, want: 9223372036854775807, wantErr: false},
		{name: "minus max LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80})}, args: args{endianness: LittleEndian}, want: -9223372036854775808, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadLongLong(tt.args.endianness)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadLongLong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadLongLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadUnsignedLongLong(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	type args struct {
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		{name: "nil BE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: BigEndian}, want: 0, wantErr: true},
		{name: "zero BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: BigEndian}, want: 0, wantErr: false},
		{name: "one BE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})}, args: args{endianness: BigEndian}, want: 1, wantErr: false},
		{name: "max BE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: BigEndian}, want: 18446744073709551615, wantErr: false},

		{name: "nil LE", fields: fields{Reader: bytes.NewReader([]byte{})}, args: args{endianness: LittleEndian}, want: 0, wantErr: true},
		{name: "zero LE", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 0, wantErr: false},
		{name: "one LE", fields: fields{Reader: bytes.NewReader([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})}, args: args{endianness: LittleEndian}, want: 1, wantErr: false},
		{name: "max LE", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})}, args: args{endianness: LittleEndian}, want: 18446744073709551615, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadUnsignedLongLong(tt.args.endianness)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadUnsignedLongLong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadUnsignedLongLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadString(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{name: "nil", fields: fields{Reader: bytes.NewReader([]byte{})}, want: "", wantErr: true},
		{name: "too small", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0x00})}, want: "", wantErr: true},
		{name: "negative", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0x00, 0x00, 0x00})}, want: "", wantErr: true},
		{name: "no string (EOF)", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x01})}, want: "", wantErr: true},
		{name: "empty", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF})}, want: "", wantErr: false},
		{name: "one", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x01, 0x61})}, want: "a", wantErr: false},
		{name: "two", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x02, 0x61, 0x62})}, want: "ab", wantErr: false},
		{name: "long", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x08, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68})}, want: "abcdefgh", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadString()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_ReadCompressedString(t *testing.T) {
	type fields struct {
		Reader *bytes.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{name: "nil", fields: fields{Reader: bytes.NewReader([]byte{})}, want: "", wantErr: true},
		{name: "too small", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0x00})}, want: "", wantErr: true},
		{name: "negative", fields: fields{Reader: bytes.NewReader([]byte{0xFF, 0x00, 0x00, 0x00})}, want: "", wantErr: true},
		{name: "no string (EOF)", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x01})}, want: "", wantErr: true},
		{name: "empty", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x00, 0x78, 0x9c, 0x01, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0x00, 0x01})}, want: "", wantErr: false},
		{name: "one", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x0d, 0x01, 0x00, 0x00, 0x00, 0x78, 0x9c, 0x4a, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0x00, 0x62, 0x00, 0x62})}, want: "a", wantErr: false},
		{name: "two", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x0e, 0x02, 0x00, 0x00, 0x00, 0x78, 0x9c, 0x4a, 0x4c, 0x02, 0x04, 0x00, 0x00, 0xff, 0xff, 0x01, 0x26, 0x00, 0xc4})}, want: "ab", wantErr: false},
		{name: "long", fields: fields{Reader: bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x30, 0x24, 0x00, 0x00, 0x00, 0x78, 0x9c, 0x4a, 0x4c, 0x4a, 0x4e, 0x49, 0x4d, 0x4b, 0xcf, 0xc8, 0xcc, 0xca, 0xce, 0xc9, 0xcd, 0xcb, 0x2f, 0x28, 0x2c, 0x2a, 0x2e, 0x29, 0x2d, 0x2b, 0xaf, 0xa8, 0xac, 0x32, 0x34, 0x32, 0x36, 0x31, 0x35, 0x33, 0xb7, 0xb0, 0x34, 0x00, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xf7, 0x0d, 0x2d})}, want: "abcdefghijklmnopqrstuvwxyz1234567890", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.ReadCompressedString()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadCompressedString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.ReadCompressedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
