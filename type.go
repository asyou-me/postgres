package postgres

import "github.com/jackc/pgx"

// 类型
const (
	Invalid uint8 = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
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
	T int16
	V string
}

// FormatCode 字段为文字格式
func (v V) FormatCode() int16 { return pgx.TextFormatCode }

// Scan 渲染数据到字符串
func (v *V) Scan(vr *pgx.ValueReader) error {
	// Not checking oid as so we can scan anything into into a NullString - may revisit this decision later
	if vr.Len() == -1 {
		v.V = ""
		return nil
	}
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
