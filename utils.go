package postgres

import "github.com/jackc/pgx"

func decodeText(vr *pgx.ValueReader) string {
	if vr.Len() == -1 {
		vr.Fatal(pgx.ProtocolError("Cannot decode null into string"))
		return ""
	}

	return vr.ReadString(vr.Len())
}

func encodeString(w *pgx.WriteBuf, oid pgx.Oid, value string) error {
	w.WriteInt32(int32(len(value)))
	w.WriteBytes([]byte(value))
	return nil
}
