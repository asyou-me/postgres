package postgres

import (
	"fmt"
	"math"
	"strconv"

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
	StringArray                // 28
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
	T           uint8
	V           string
	IntArray    *[]int64
	StringArray *[]string
}

// MarshalJSON 序列化时调用
func (v *V) MarshalJSON() ([]byte, error) {
	switch v.T {
	case Int, Int8, Int16, Int32, Int64:
		if len(v.V) == 0 {
			return []byte{'0'}, nil
		}
		return []byte(v.V), nil
	case Bool:
		if v.V == "true" {
			return []byte(v.V), nil
		}
		return []byte("false"), nil
	case String:
		return []byte(strconv.Quote(v.V)), nil
	case IntArray:
		if v.IntArray == nil {
			return []byte("[]"), nil
		}
		datas := *v.IntArray
		if len(datas) == 1 {
			return []byte("[" + fmt.Sprint(datas[0]) + "]"), nil
		}
		strs := make([]string, len(datas))
		n := len(datas) - 1
		for k, v := range datas {
			str := fmt.Sprint(v)
			strs[k] = str
			n = n + len(str)
		}

		b := make([]byte, n+2)
		bp := 0
		bp += copy(b, "[")
		bp += copy(b[bp:], strs[0])
		for _, v := range strs[1:] {
			bp += copy(b[bp:], ",")
			bp += copy(b[bp:], v)
		}
		bp += copy(b[bp:], "]")
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

	switch vr.Type().DataType {
	case pgx.BoolOid:
		b := vr.ReadByte()
		if b != 0 {
			v.V = "true"
		} else {
			v.V = "false"
		}
		v.T = Bool
	case pgx.Int2Oid:
		v.V = fmt.Sprint(vr.ReadInt16())
		v.T = Int16
	case pgx.Int4Oid:
		v.V = fmt.Sprint(vr.ReadInt32())
		v.T = Int32
	case pgx.Int8Oid:
		v.V = fmt.Sprint(vr.ReadInt64())
		v.T = Int64
	case pgx.Float4Oid:
		i := vr.ReadInt32()
		v.V = fmt.Sprint(math.Float32frombits(uint32(i)))
		v.T = Float32
	case pgx.Float8Oid:
		i := vr.ReadInt64()
		v.V = fmt.Sprint(math.Float64frombits(uint64(i)))
		v.T = Float64
	case pgx.TextOid:
		v.V = decodeText(vr)
		v.T = String
	case pgx.Int8ArrayOid:
		is := decodeInt8Array(vr)
		v.IntArray = &is
		v.T = IntArray
	default:
		return fmt.Errorf("cannot encode  oid %v", vr.Type().Table)
	}

	return vr.Err()
}

// Encode 写到数据
func (v V) Encode(w *pgx.WriteBuf, oid pgx.Oid) error {
	if v.T == IntArray {
		datas := *v.IntArray
		if len(datas) == 1 {
			b := []byte("{" + fmt.Sprint(datas[0]) + "}")
			w.WriteInt32(int32(len(b)))
			w.WriteBytes(b)
			return nil
		}
		strs := make([]string, len(datas))
		n := len(datas) - 1
		for k, v := range datas {
			str := fmt.Sprint(v)
			strs[k] = str
			n = n + len(str)
		}

		b := make([]byte, n+2)
		bp := 0
		bp += copy(b, "{")
		bp += copy(b[bp:], strs[0])
		for _, v := range strs[1:] {
			bp += copy(b[bp:], ",")
			bp += copy(b[bp:], v)
		}
		bp += copy(b[bp:], "}")
		w.WriteInt32(int32(n + 2))
		w.WriteBytes(b)
		return nil
	}
	if v.V == "" {
		w.WriteInt32(-1)
		return nil
	}
	w.WriteInt32(int32(len(v.V)))
	w.WriteBytes([]byte(v.V))
	return nil
}
