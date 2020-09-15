// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package sectorstorage

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

func (t *Call) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.State (sectorstorage.CallState) (uint64)
	if len("State") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"State\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("State"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("State")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.State)); err != nil {
		return err
	}

	// t.Result ([]uint8) (slice)
	if len("Result") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Result\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Result"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Result")); err != nil {
		return err
	}

	if len(t.Result) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Result was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Result))); err != nil {
		return err
	}

	if _, err := w.Write(t.Result[:]); err != nil {
		return err
	}
	return nil
}

func (t *Call) UnmarshalCBOR(r io.Reader) error {
	*t = Call{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("Call: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.State (sectorstorage.CallState) (uint64)
		case "State":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.State = CallState(extra)

			}
			// t.Result ([]uint8) (slice)
		case "Result":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Result: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Result = make([]uint8, extra)
			}

			if _, err := io.ReadFull(br, t.Result[:]); err != nil {
				return err
			}

		default:
			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
		}
	}

	return nil
}