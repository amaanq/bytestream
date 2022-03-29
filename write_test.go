package bytestream

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWriter_WriteBytes(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// {name: "nil", fields: fields{Buffer: nil}, args: args{bytes: nil}, wantErr: true}, This panics lol
		{name: "no byte", fields: fields{Buffer: new(bytes.Buffer)}, args: args{bytes: []byte{}}, wantErr: false},
		{name: "one byte", fields: fields{Buffer: new(bytes.Buffer)}, args: args{bytes: []byte{0x01}}, wantErr: false},
		{name: "many bytes", fields: fields{Buffer: new(bytes.Buffer)}, args: args{bytes: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteBytes(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteBool(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data  bool
		count int8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "false", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: false, count: 1}, wantErr: false},
		{name: "true", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: true, count: 1}, wantErr: false},
		{name: "true (2)", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: true, count: 2}, wantErr: false},
		{name: "true (3)", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: true, count: 3}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteBool(tt.args.data, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteBool() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteInt8(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data int8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2}, wantErr: false},
		{name: "minus one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 127}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteInt8(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteInt8() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteUInt8(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 255}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteUInt8(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteUInt8() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteInt16(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data       int16
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: BigEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: BigEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: BigEndian}, wantErr: false},
		{name: "minus one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1, endianness: BigEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 32767, endianness: BigEndian}, wantErr: false},

		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: LittleEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: LittleEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: LittleEndian}, wantErr: false},
		{name: "minus one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1, endianness: LittleEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 32767, endianness: LittleEndian}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteInt16(tt.args.data, tt.args.endianness); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteInt16() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteUInt16(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data       uint16
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: BigEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: BigEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: BigEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 65535, endianness: BigEndian}, wantErr: false},

		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: LittleEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: LittleEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: LittleEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 65535, endianness: LittleEndian}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteUInt16(tt.args.data, tt.args.endianness); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteUInt16() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteInt24(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data       int32
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: BigEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: BigEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: BigEndian}, wantErr: false},
		{name: "minus one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1, endianness: BigEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 8388607, endianness: BigEndian}, wantErr: false},
		{name: "out of bounds", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 8388608, endianness: BigEndian}, wantErr: true},

		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: LittleEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: LittleEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: LittleEndian}, wantErr: false},
		{name: "minus one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1, endianness: LittleEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 8388607, endianness: LittleEndian}, wantErr: false},
		{name: "out of bounds", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 8388608, endianness: LittleEndian}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteInt24(tt.args.data, tt.args.endianness); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteInt24() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteUInt24(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data       uint32
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: BigEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: BigEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: BigEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 16777215, endianness: BigEndian}, wantErr: false},
		{name: "out of bounds", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 16777216, endianness: BigEndian}, wantErr: true},

		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: LittleEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: LittleEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: LittleEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 16777215, endianness: LittleEndian}, wantErr: false},
		{name: "out of bounds", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 16777216, endianness: LittleEndian}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteUInt24(tt.args.data, tt.args.endianness); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteUInt24() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteInt32(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data       int32
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: BigEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: BigEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: BigEndian}, wantErr: false},
		{name: "minus one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1, endianness: BigEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2147483647, endianness: BigEndian}, wantErr: false},

		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: LittleEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: LittleEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: LittleEndian}, wantErr: false},
		{name: "minus one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1, endianness: LittleEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2147483647, endianness: LittleEndian}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteInt32(tt.args.data, tt.args.endianness); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteInt32() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteUInt32(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data       uint32
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: BigEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: BigEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: BigEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 4294967295, endianness: BigEndian}, wantErr: false},

		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: LittleEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: LittleEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: LittleEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 4294967295, endianness: LittleEndian}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteUInt32(tt.args.data, tt.args.endianness); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteUInt32() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteInt64(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data       int64
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: BigEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: BigEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: BigEndian}, wantErr: false},
		{name: "minus one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1, endianness: BigEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 9223372036854775807, endianness: BigEndian}, wantErr: false},

		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: LittleEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: LittleEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: LittleEndian}, wantErr: false},
		{name: "minus one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1, endianness: LittleEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 9223372036854775807, endianness: LittleEndian}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteInt64(tt.args.data, tt.args.endianness); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteInt64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteUInt64(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data       uint64
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: BigEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: BigEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: BigEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 18446744073709551615, endianness: BigEndian}, wantErr: false},

		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: LittleEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: LittleEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: LittleEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 18446744073709551615, endianness: LittleEndian}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteUInt64(tt.args.data, tt.args.endianness); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteUInt64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteVarInt(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2}, wantErr: false},
		{name: "minus one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 9223372036854775807}, wantErr: false},
		{name: "minus max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -9223372036854775808}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteVarInt(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteVarInt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteUVarInt(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 18446744073709551615}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteUVarInt(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteUVarInt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteLong(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data       int64
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: BigEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: BigEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: BigEndian}, wantErr: false},
		{name: "minus one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 9223372036854775807, endianness: BigEndian}, wantErr: false},

		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: LittleEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: LittleEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: LittleEndian}, wantErr: false},
		{name: "minus one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 9223372036854775807, endianness: LittleEndian}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteLong(tt.args.data, tt.args.endianness); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteLong() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteUnsignedLong(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data       uint64
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: BigEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: BigEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: BigEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 18446744073709551615, endianness: BigEndian}, wantErr: false},

		{name: "zero", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: LittleEndian}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: LittleEndian}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: LittleEndian}, wantErr: false},
		{name: "max", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 18446744073709551615, endianness: LittleEndian}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteUnsignedLong(tt.args.data, tt.args.endianness); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteUnsignedLong() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteLongLong(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data       int64
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero BE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: BigEndian}, wantErr: false},
		{name: "one BE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: BigEndian}, wantErr: false},
		{name: "two BE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: BigEndian}, wantErr: false},
		{name: "minus one BE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1}, wantErr: false},
		{name: "max BE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 9223372036854775807, endianness: BigEndian}, wantErr: false},

		{name: "zero LE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: LittleEndian}, wantErr: false},
		{name: "one LE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: LittleEndian}, wantErr: false},
		{name: "two LE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: LittleEndian}, wantErr: false},
		{name: "minus one LE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: -1}, wantErr: false},
		{name: "max LE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 9223372036854775807, endianness: LittleEndian}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteLongLong(tt.args.data, tt.args.endianness); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteLongLong() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteUnsignedLongLong(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data       uint64
		endianness Endianness
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "zero BE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: BigEndian}, wantErr: false},
		{name: "one BE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: BigEndian}, wantErr: false},
		{name: "two BE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: BigEndian}, wantErr: false},
		{name: "max BE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 18446744073709551615, endianness: BigEndian}, wantErr: false},

		{name: "zero LE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 0, endianness: LittleEndian}, wantErr: false},
		{name: "one LE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 1, endianness: LittleEndian}, wantErr: false},
		{name: "two LE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 2, endianness: LittleEndian}, wantErr: false},
		{name: "max LE", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: 18446744073709551615, endianness: LittleEndian}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteUnsignedLongLong(tt.args.data, tt.args.endianness); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteUnsignedLongLong() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteString(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "empty", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: ""}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "a"}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "ab"}, wantErr: false},
		{name: "long", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "abcdefghijklmnopqrstuvwxyz1234567890"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteString(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteString() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(tt.fields.Buffer.Bytes())
		})
	}
}

func TestWriter_WriteStringSize(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data     string
		bytesize int8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "empty", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "", bytesize: 1}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "a", bytesize: 1}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "ab", bytesize: 1}, wantErr: false},
		{name: "long", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "abcdefghijklmnopqrstuvwxyz1234567890", bytesize: 1}, wantErr: false},

		{name: "empty", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "", bytesize: 2}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "a", bytesize: 2}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "ab", bytesize: 2}, wantErr: false},
		{name: "long", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "abcdefghijklmnopqrstuvwxyz1234567890", bytesize: 2}, wantErr: false},

		{name: "empty", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "", bytesize: 3}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "a", bytesize: 3}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "ab", bytesize: 3}, wantErr: false},
		{name: "long", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "abcdefghijklmnopqrstuvwxyz1234567890", bytesize: 3}, wantErr: false},

		{name: "empty", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "", bytesize: 4}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "a", bytesize: 4}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "ab", bytesize: 4}, wantErr: false},
		{name: "long", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "abcdefghijklmnopqrstuvwxyz1234567890", bytesize: 4}, wantErr: false},

		{name: "empty", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "", bytesize: 8}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "a", bytesize: 8}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "ab", bytesize: 8}, wantErr: false},
		{name: "long", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "abcdefghijklmnopqrstuvwxyz1234567890", bytesize: 8}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteStringSize(tt.args.data, tt.args.bytesize); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteStringSize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteCompressedString(t *testing.T) {
	type fields struct {
		Buffer *bytes.Buffer
	}
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "empty", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: ""}, wantErr: false},
		{name: "one", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "a"}, wantErr: false},
		{name: "two", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "ab"}, wantErr: false},
		{name: "long", fields: fields{Buffer: new(bytes.Buffer)}, args: args{data: "abcdefghijklmnopqrstuvwxyz1234567890"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Buffer: tt.fields.Buffer,
			}
			if err := w.WriteCompressedString(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteCompressedString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
