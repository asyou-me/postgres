package postgres

import (
	"fmt"

	"github.com/jackc/pgx"
)

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

func decodeInt8Array(vr *pgx.ValueReader) []int64 {
	if vr.Len() == -1 {
		return nil
	}

	if vr.Type().DataType != pgx.Int8ArrayOid {
		vr.Fatal(pgx.ProtocolError(fmt.Sprintf("Cannot decode oid %v into []int64", vr.Type().DataType)))
		return nil
	}

	if vr.Type().FormatCode != pgx.BinaryFormatCode {
		vr.Fatal(pgx.ProtocolError(fmt.Sprintf("Unknown field description format code: %v", vr.Type().FormatCode)))
		return nil
	}

	numElems, err := decode1dArrayHeader(vr)
	if err != nil {
		vr.Fatal(err)
		return nil
	}

	a := make([]int64, int(numElems))
	for i := 0; i < len(a); i++ {
		elSize := vr.ReadInt32()
		switch elSize {
		case 8:
			a[i] = vr.ReadInt64()
		case -1:
			vr.Fatal(pgx.ProtocolError("Cannot decode null element"))
			return nil
		default:
			vr.Fatal(pgx.ProtocolError(fmt.Sprintf("Received an invalid size for an int8 element: %d", elSize)))
			return nil
		}
	}

	return a
}

func decode1dArrayHeader(vr *pgx.ValueReader) (length int32, err error) {
	numDims := vr.ReadInt32()
	if numDims > 1 {
		return 0, pgx.ProtocolError(fmt.Sprintf("Expected array to have 0 or 1 dimension, but it had %v", numDims))
	}

	vr.ReadInt32() // 0 if no nulls / 1 if there is one or more nulls -- but we don't care
	vr.ReadInt32() // element oid

	if numDims == 0 {
		return 0, nil
	}

	length = vr.ReadInt32()

	idxFirstElem := vr.ReadInt32()
	if idxFirstElem != 1 {
		return 0, pgx.ProtocolError(fmt.Sprintf("Expected array's first element to start a index 1, but it is %d", idxFirstElem))
	}

	return length, nil
}
