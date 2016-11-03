package postgres

import "github.com/jackc/pgx"

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

// FormatCode 必须支持函数
func (GSTYPE) FormatCode() int16 { return pgx.TextFormatCode }

// Encode 将数据转换到pgx
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
