package eidiom

type TestApplyOptsOpts struct {
	Opt1 int
	Opt2 string
}

//func TestApplyOpts(t *testing.T) {
//	v := TestApplyOptsOpts{
//		Opt1: 10,
//		Opt2: "hello",
//	}
//
//	f1 := func(o TestApplyOptsOpts) TestApplyOptsOpts {
//		o.Opt1 = 99
//		return o
//	}
//	f2 := func(o TestApplyOptsOpts) TestApplyOptsOpts {
//		o.Opt2 = "world"
//		return o
//	}
//
//	if v1 := ApplyOpts(v, f1); v1.Opt1 != 99 {
//		t.Error(v1)
//	}
//}
