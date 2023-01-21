package util

import (
	"testing"
	"time"
)

func checkEqualProps(t *testing.T, prop string, valIn, valOut interface{}) {
	if valIn != valOut {
		t.Errorf("expected prop '%s' to be equal, got %v, want %v", prop, valIn, valOut)
	}
}

func checkUnmodifiedProp(t *testing.T, prop string, got, want interface{}) {
	if got != want {
		t.Errorf("property %s was modified with %v when it should not", prop, got)
	}
}

func TestMergeStructs(t *testing.T) {
	t.Run("should merge two equal structs", func(t *testing.T) {
		type S struct {
			A int
			B string
			C time.Time
		}

		in, out := S{A: 1, B: "struct", C: time.Now()}, S{}
		MergeStructs(in, &out)
		checkEqualProps(t, "A", in.A, out.A)
		checkEqualProps(t, "B", in.B, out.B)
		checkEqualProps(t, "C", in.C, out.C)
	})

	t.Run("should merge to different structs", func(t *testing.T) {
		type S struct {
			A int
			B string
		}
		type P struct {
			A int
			C string
			D *S
		}

		in, out := S{A: 1, B: "struct"}, P{}
		MergeStructs(in, &out)
		checkEqualProps(t, "A", in.A, out.A)
		checkUnmodifiedProp(t, "C", out.C, "")
		checkUnmodifiedProp(t, "D", out.D, (*S)(nil))
	})

	t.Run("should not merge structs that do not share properties", func(t *testing.T) {
		type S struct {
			A int
			B string
		}
		type P struct {
			C string
			D *S
		}

		in, out := S{A: 1, B: "struct"}, P{}
		MergeStructs(in, &out)
		checkUnmodifiedProp(t, "C", out.C, "")
		checkUnmodifiedProp(t, "D", out.D, (*S)(nil))
	})
}
