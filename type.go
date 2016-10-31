package postgres

import "github.com/jackc/pgx"

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

