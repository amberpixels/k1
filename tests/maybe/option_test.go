package maybe_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/amberpixels/k1/maybe"
	"github.com/expectto/be"
	"github.com/expectto/be/be_reflected"
)

func TestNoneAndSomeMethods(t *testing.T) {
	optNone := maybe.None[int]()
	be.Expect(t, optNone.None()).To(be.True())
	be.Expect(t, optNone.Some()).To(be.False())
	be.Expect(t, optNone.Some(42)).To(be.False())

	optSome := maybe.Some(42)
	be.Expect(t, optSome.None()).To(be.False())
	be.Expect(t, optSome.Some()).To(be.True())
	be.Expect(t, optSome.Some(42)).To(be.True())
	be.Expect(t, optSome.Some(100)).To(be.False())
}

func TestSomePanicsWithMultipleArgs(t *testing.T) {
	opt := maybe.Some(42)
	be.Expect(t, func() { opt.Some(1, 2) }).To(be.Panic())
}

func TestUnwrap(t *testing.T) {
	optSome := maybe.Some("hello")
	be.Expect(t, optSome.Unwrap()).To(be.Eq("hello"))

	optNone := maybe.None[string]()
	be.Expect(t, func() { optNone.Unwrap() }).To(be.Panic())
}

func TestJSONMarshalling(t *testing.T) {
	optSome := maybe.Some(100)
	marshalledSome, err := json.Marshal(optSome)
	be.Expect(t, err).To(be.Succeed())
	be.Expect(t, string(marshalledSome)).To(be.Eq("100"))

	optNone := maybe.None[int]()
	marshalledNone, err := json.Marshal(&optNone) // pointer receiver works too
	be.Expect(t, err).To(be.Succeed())
	be.Expect(t, string(marshalledNone)).To(be.Eq("null"))
}

func TestJSONUnmarshalling(t *testing.T) {
	var opt maybe.Option[int]

	err := json.Unmarshal([]byte("123"), &opt)
	be.Expect(t, err).To(be.Succeed())
	be.Expect(t, opt.Some()).To(be.True())
	be.Expect(t, opt.Some(123)).To(be.True())

	err = json.Unmarshal([]byte("null"), &opt)
	be.Expect(t, err).To(be.Succeed())
	be.Expect(t, opt.None()).To(be.True())
}

func TestJSONUnmarshallingError(t *testing.T) {
	var opt maybe.Option[int]
	// "abc" is not valid JSON for an int -> the inner json.Unmarshal must error.
	err := opt.UnmarshalJSON([]byte(`"abc"`))
	be.Expect(t, err).To(be.HaveOccurred())
}

func TestMarshalTOML(t *testing.T) {
	// Some encodes to the underlying value.
	optSome := maybe.Some(7)
	data, err := optSome.MarshalTOML()
	be.Expect(t, err).To(be.Succeed())
	be.Expect(t, string(data)).To(be.Eq("7"))

	// None encodes to the TomlNone sentinel (as a JSON string).
	optNone := maybe.None[int]()
	data, err = optNone.MarshalTOML()
	be.Expect(t, err).To(be.Succeed())
	be.Expect(t, string(data)).To(be.Eq(`"` + maybe.TomlNone + `"`))
}

func TestUnmarshalTextScalarString(t *testing.T) {
	var opt maybe.Option[string]

	err := opt.UnmarshalText([]byte(`"test"`))
	be.Expect(t, err).To(be.Succeed())
	be.Expect(t, opt.Some()).To(be.True())
	be.Expect(t, opt.Some("test")).To(be.True())
}

func TestUnmarshalTextScalarBool(t *testing.T) {
	var opt maybe.Option[bool]

	err := opt.UnmarshalText([]byte("true"))
	be.Expect(t, err).To(be.Succeed())
	be.Expect(t, opt.Some(true)).To(be.True())
}

func TestUnmarshalTextScalarFloat(t *testing.T) {
	var opt maybe.Option[float64]

	err := opt.UnmarshalText([]byte("3.14"))
	be.Expect(t, err).To(be.Succeed())
	be.Expect(t, opt.Some(3.14)).To(be.True())
}

func TestUnmarshalTextNoneVariants(t *testing.T) {
	cases := []string{"", "   ", "null", "NULL", "None", "none", maybe.TomlNone}
	for _, in := range cases {
		var opt maybe.Option[string]
		err := opt.UnmarshalText([]byte(in))
		be.Expect(t, err).To(be.Succeed())
		be.Expect(t, opt.None()).To(be.True())
	}
}

func TestUnmarshalTextViaTextUnmarshaler(t *testing.T) {
	// time.Time implements encoding.TextUnmarshaler and is comparable,
	// so it exercises the TextUnmarshaler branch.
	var opt maybe.Option[time.Time]
	err := opt.UnmarshalText([]byte("2020-01-02T03:04:05Z"))
	be.Expect(t, err).To(be.Succeed())
	be.Expect(t, opt.Some()).To(be.True())

	want, _ := time.Parse(time.RFC3339, "2020-01-02T03:04:05Z")
	be.Expect(t, opt.Unwrap().Equal(want)).To(be.True())
}

func TestUnmarshalTextTextUnmarshalerError(t *testing.T) {
	var opt maybe.Option[time.Time]
	err := opt.UnmarshalText([]byte("not-a-time"))
	be.Expect(t, err).To(be.HaveOccurred())
}

func TestUnmarshalTextScalarJSONError(t *testing.T) {
	// float64 takes the scalar branch; "abc" is invalid JSON -> error.
	var opt maybe.Option[float64]
	err := opt.UnmarshalText([]byte("abc"))
	be.Expect(t, err).To(be.HaveOccurred())
}

func TestUnmarshalTextUnsupportedType(t *testing.T) {
	// int is comparable, not a TextUnmarshaler, and not in the scalar switch
	// (only float64/string/bool are) -> the final "no fallback" error.
	var opt maybe.Option[int]
	err := opt.UnmarshalText([]byte("5"))
	be.Expect(t, err).To(be.HaveOccurred())
}

func TestIsZero(t *testing.T) {
	// None is zero.
	optNone := maybe.None[int]()
	be.Expect(t, optNone.IsZero()).To(be.True())

	// Some(zero value) is also zero.
	optZero := maybe.Some(0)
	be.Expect(t, optZero.IsZero()).To(be.True())

	// Some(non-zero) is not zero.
	optSome := maybe.Some(42)
	be.Expect(t, optSome.IsZero()).To(be.False())
}

func TestBoolHelpers(t *testing.T) {
	tru := maybe.True()
	be.Expect(t, tru.Some(true)).To(be.True())

	fls := maybe.False()
	be.Expect(t, fls.Some(false)).To(be.True())

	nb := maybe.NoneBool()
	be.Expect(t, nb.None()).To(be.True())

	// Bool is an alias for Option[bool].
	b := maybe.True()
	be.Expect(t, b.Some()).To(be.True())
}

func TestIntHelpers(t *testing.T) {
	ni := maybe.NoneInt()
	be.Expect(t, ni.None()).To(be.True())

	i := maybe.Some(5)
	be.Expect(t, i.Some(5)).To(be.True())
}

// TestReflectedAndSanity exercises a couple of `be` subpackage matchers so the
// dogfooding covers more than just be.Eq, and double-checks be.NotPanic.
func TestReflectedAndSanity(t *testing.T) {
	optSome := maybe.Some("hi")
	be.Expect(t, optSome.Unwrap()).To(be_reflected.AsString())

	be.Expect(t, func() {
		o := maybe.Some(1)
		_ = o.Unwrap()
	}).To(be.NotPanic())

	// HaveLength on the marshalled bytes of a None.
	data, _ := maybe.None[int]().MarshalJSON()
	be.Expect(t, string(data)).To(be.HaveLength(4)) // "null"
}
