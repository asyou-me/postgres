package postgres

import (
	"fmt"

	"errors"

	"github.com/jackc/pgx"
)

// 类型
const (
	Invalid       uint8 = iota // 0
	Bool                       // 1
	Int                        // 2
	Int8                       // 3
	Int16                      // 4
	Int32                      // 5
	Int64                      // 6
	Uint                       // 7
	Uint8                      // 8
	Uint16                     // 9
	Uint32                     // 10
	Uint64                     // 11
	Uintptr                    // 12
	Float32                    // 13
	Float64                    // 14
	Complex64                  // 15
	Complex128                 // 16
	Array                      // 17
	Chan                       // 18
	Func                       // 19
	Interface                  // 20
	Map                        // 21
	Ptr                        // 22
	Slice                      // 23
	String                     // 24
	Struct                     // 25
	UnsafePointer              // 26
	IntArray                   // 27
)

// GSTYPE json 格式化数据
type GSTYPE struct {
	Key   string
	Path  string
	Value string
}

// Scan 渲染数据到字符串
func (s *GSTYPE) Scan(vr *pgx.ValueReader) error {
	// Not checking oid as so we can scan anything into into a NullString - may revisit this decision later
	fmt.Println("s.Value:", s.Value)
	if vr.Len() == -1 {
		s.Value = ""
		return nil
	}
	s.Value = decodeText(vr)
	return vr.Err()
}

// FormatCode 字段为文字格式
func (s GSTYPE) FormatCode() int16 { return pgx.TextFormatCode }

// Encode 写到数据
func (s GSTYPE) Encode(w *pgx.WriteBuf, oid pgx.Oid) error {
	if s.Value == "" {
		w.WriteInt32(-1)
		return nil
	}
	return encodeString(w, oid, s.Value)
}

// V 基础类型
type V struct {
	T uint8
	V string
}

// MarshalJSON 序列化时调用
func (v *V) MarshalJSON() ([]byte, error) {
	switch v.T {
	case Int, Int8, Int16, Int32, Int64, Bool:
		return []byte(v.V), nil
	case String:
		return []byte(`"` + v.V + `"`), nil
	case IntArray:
		lenV := len(v.V)
		b := make([]byte, lenV)
		copy(b, "[")
		copy(b[1:], v.V[1:])
		copy(b[lenV-1:], "]")
		fmt.Println(string(b))
		return b, nil
	default:
		return []byte{}, errors.New("无法识别类型:" + fmt.Sprint(v.T))
	}
}

// FormatCode 字段为文字格式
func (v V) FormatCode() int16 { return pgx.TextFormatCode }

// Scan 渲染数据到字符串
func (v *V) Scan(vr *pgx.ValueReader) error {
	if vr.Len() == -1 {
		v.V = ""
		return nil
	}

	fmt.Println(vr.Type())
	fmt.Println(decodeText(vr))

	v.V = decodeText(vr)
	return vr.Err()
}

// Encode 写到数据
func (v V) Encode(w *pgx.WriteBuf, oid pgx.Oid) error {
	if v.V == "" {
		w.WriteInt32(-1)
		return nil
	}
	w.WriteInt32(int32(len(v.V)))
	w.WriteBytes([]byte(v.V))
	return nil
}
