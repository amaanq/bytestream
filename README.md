# **Package bytestream**

### This package is 100% compatible with the stdlib and mainly works with the bytes.Buffer and bytes.Reader interfaces. 

#### **Installation**
```sh
go get -u github.com/amaanq/bytestream
```

#### **Usage**
```go
import (
    "fmt"
    "github.com/amaanq/bytestream"
)

func main() {
    Writer := &bytestream.Writer{
		buffer: bytes.NewBuffer([]byte{}),
	}
	Writer.WriteCompressedString("hello there!")
    Writer.WriteInt32(40, bytestream.BigEndian)
    // You've now packed hello there! and 40 into a very small byte array to do with what you please..
}
```

### **NOTE: This package is a custom protocol defined by a game company for their TCP data streams.**
<br/>

### As such, I do not recommend using it to detect/use with protocols already implemented. However, you can use this for your own protocol for whatever it may be. 
<br/>

### The file including the methods for "tags" have been omitted for security and privacy reasons, so don't bother trying to use the read/write LogicLong methods. I don't want any simpleton to figure out what game(s) this is for and use that to their knowledge.
<br/>

# **Enjoy using this!**