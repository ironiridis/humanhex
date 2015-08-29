package humanhex

import "testing"

type testcase struct {
	try    []byte
	contig int
	want   string
}

func z(t *testing.T, in []testcase) {
	t.Parallel()
	for _, tc := range in {
		got := String(tc.try, tc.contig)
		if got != tc.want {
			t.Errorf("String(%q, %v) == %q, wanted %q", tc.try, tc.contig, got, tc.want)
		}
	}
}

func TestAllNulls(t *testing.T) {
	z(t, []testcase{
		{[]byte{0, 0, 0, 0}, 1, "\\x00\\x00\\x00\\x00"},
		{[]byte{0, 0, 0, 0, 0}, 1, "\\x00\\x00\\x00\\x00\\x00"},
		{[]byte{0, 0, 0, 0}, 2, "\\x00\\x00\\x00\\x00"},
		{[]byte{0, 0, 0, 0, 0}, 2, "\\x00\\x00\\x00\\x00\\x00"},
		{[]byte{0, 0, 0, 0}, 10, "\\x00\\x00\\x00\\x00"},
		{[]byte{0, 0, 0, 0, 0}, 10, "\\x00\\x00\\x00\\x00\\x00"},
	})
}

// for coverage
func TestAllFFs(t *testing.T) {
	z(t, []testcase{
		{[]byte{0xff, 0xff, 0xff, 0xff}, 1, "\\xff\\xff\\xff\\xff"},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff}, 1, "\\xff\\xff\\xff\\xff\\xff"},
		{[]byte{0xff, 0xff, 0xff, 0xff}, 2, "\\xff\\xff\\xff\\xff"},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff}, 2, "\\xff\\xff\\xff\\xff\\xff"},
		{[]byte{0xff, 0xff, 0xff, 0xff}, 10, "\\xff\\xff\\xff\\xff"},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff}, 10, "\\xff\\xff\\xff\\xff\\xff"},
	})
}

func TestStartReadable(t *testing.T) {
	z(t, []testcase{
		{[]byte{0x34, 0x34, 0, 0, 0, 0}, 1, "44\\x00\\x00\\x00\\x00"},
		{[]byte{0x34, 0x34, 0, 0, 0, 0, 0}, 1, "44\\x00\\x00\\x00\\x00\\x00"},
		{[]byte{0x34, 0x34, 0, 0, 0, 0}, 2, "44\\x00\\x00\\x00\\x00"},
		{[]byte{0x34, 0x34, 0, 0, 0, 0, 0}, 2, "44\\x00\\x00\\x00\\x00\\x00"},
		{[]byte{0x34, 0x34, 0, 0, 0, 0}, 10, "\\x34\\x34\\x00\\x00\\x00\\x00"},
		{[]byte{0x34, 0x34, 0, 0, 0, 0, 0}, 10, "\\x34\\x34\\x00\\x00\\x00\\x00\\x00"},
	})
}

func TestEndReadable(t *testing.T) {
	z(t, []testcase{
		{[]byte{0, 0, 0, 0, 0x34, 0x34}, 1, "\\x00\\x00\\x00\\x0044"},
		{[]byte{0, 0, 0, 0, 0, 0x34, 0x34}, 1, "\\x00\\x00\\x00\\x00\\x0044"},
		{[]byte{0, 0, 0, 0, 0x34, 0x34}, 2, "\\x00\\x00\\x00\\x0044"},
		{[]byte{0, 0, 0, 0, 0, 0x34, 0x34}, 2, "\\x00\\x00\\x00\\x00\\x0044"},
		{[]byte{0, 0, 0, 0, 0x34, 0x34}, 10, "\\x00\\x00\\x00\\x00\\x34\\x34"},
		{[]byte{0, 0, 0, 0, 0, 0x34, 0x34}, 10, "\\x00\\x00\\x00\\x00\\x00\\x34\\x34"},
	})
}

func TestMidReadable(t *testing.T) {
	z(t, []testcase{
		{[]byte{0, 0, 0x34, 0x34, 0, 0}, 1, "\\x00\\x0044\\x00\\x00"},
		{[]byte{0, 0, 0x34, 0x34, 0, 0, 0}, 1, "\\x00\\x0044\\x00\\x00\\x00"},
		{[]byte{0, 0, 0x34, 0x34, 0, 0}, 2, "\\x00\\x0044\\x00\\x00"},
		{[]byte{0, 0, 0x34, 0x34, 0, 0, 0}, 2, "\\x00\\x0044\\x00\\x00\\x00"},
		{[]byte{0, 0, 0x34, 0x34, 0, 0}, 10, "\\x00\\x00\\x34\\x34\\x00\\x00"},
		{[]byte{0, 0, 0x34, 0x34, 0, 0, 0}, 10, "\\x00\\x00\\x34\\x34\\x00\\x00\\x00"},
	})
}

func TestAlternation(t *testing.T) {
	z(t, []testcase{
		{[]byte{0x34, 0x34, 0, 0, 0x34, 0x34, 0, 0}, 1, "44\\x00\\x0044\\x00\\x00"},
		{[]byte{0x34, 0x34, 0, 0, 0x34, 0x34, 0, 0, 0}, 1, "44\\x00\\x0044\\x00\\x00\\x00"},
		{[]byte{0x34, 0x34, 0, 0, 0x34, 0x34, 0, 0}, 2, "44\\x00\\x0044\\x00\\x00"},
		{[]byte{0x34, 0x34, 0, 0, 0x34, 0x34, 0, 0, 0}, 2, "44\\x00\\x0044\\x00\\x00\\x00"},
		{[]byte{0x34, 0x34, 0, 0, 0x34, 0x34, 0, 0}, 10, "\\x34\\x34\\x00\\x00\\x34\\x34\\x00\\x00"},
		{[]byte{0x34, 0x34, 0, 0, 0x34, 0x34, 0, 0, 0}, 10, "\\x34\\x34\\x00\\x00\\x34\\x34\\x00\\x00\\x00"},
	})
}

func TestEmpty(t *testing.T) {
	z(t, []testcase{
		{[]byte{}, 1, ""},
		{[]byte{}, 2, ""},
		{[]byte{}, 10, ""},
	})
}

func TestInvalidM(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	String([]byte{}, 0)
	t.Error("String({}, 0) failed to panic")
}
