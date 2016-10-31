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

func (n GSTYPE) FormatCode() int16 { return pgx.TextFormatCode }

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
