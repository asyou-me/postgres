package postgres

import (
	"fmt"

	"github.com/jackc/pgx"
)

// 类型
const (
	Invalid       uint8 = 0
	Bool          uint8 = 1
	Int           uint8 = 2
	Int8          uint8 = 3
	Int16         uint8 = 4
	Int32         uint8 = 5
	Int64         uint8 = 6
	Uint          uint8 = 7
	Uint8         uint8 = 8
	Uint16        uint8 = 9
	Uint32        uint8 = 10
	Uint64        uint8 = 11
	Uintptr       uint8 = 12
	Float32       uint8 = 13
	Float64       uint8 = 14
	Complex64     uint8 = 15
	Complex128    uint8 = 16
	Array         uint8 = 17
	Chan          uint8 = 18
	Func          uint8 = 19
	Interface     uint8 = 20
	Map           uint8 = 21
	Ptr           uint8 = 22
	Slice         uint8 = 23
	String        uint8 = 24
	Struct        uint8 = 25
	UnsafePointer uint8 = 26
	IntArray      uint8 = 27
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

// FormatCode 字段为文字格式
func (v V) FormatCode() int16 { return pgx.TextFormatCode }

// Scan 渲染数据到字符串
func (v *V) Scan(vr *pgx.ValueReader) error {
	fmt.Println("v.V:", v.V)
	if vr.Len() == -1 {
		v.V = ""
		return nil
	}
	v.V = decodeText(vr)
	return vr.Err()
}

// Encode 写到数据
func (v V) Encode(w *pgx.WriteBuf, oid pgx.Oid) error {
	fmt.Println("v.V:", v.V)
	if v.V == "" {
		w.WriteInt32(-1)
		return nil
	}
	w.WriteInt32(int32(len(v.V)))
	w.WriteBytes([]byte(v.V))
	return nil
}
