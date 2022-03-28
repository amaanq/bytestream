package bytestream

const Is64Bit = uint64(^uintptr(0)) == ^uint64(0)

type Sign bool

const (
	Signed   Sign = true
	Unsigned Sign = false
)

type Endianness bool

const (
	LittleEndian Endianness = true
	BigEndian    Endianness = false
)

const (
	ByteSize      = iota + 1
	BoolSize      = ByteSize
	Int8Size      = ByteSize
	Int16Size     = Int8Size * 2
	Int24Size     = Int16Size + Int8Size
	Int32Size     = Int16Size * 2
	Int64Size     = Int32Size * 2
	LongSize      = Int64Size
	LogicLongSize = LongSize
	LongLongSize  = LongSize * 2
)
