package tool

func Choose1f2[T1, T2 any](t1 T1, t2 T2) T1 {
	return t1
}

func Choose2f2[T1, T2 any](t1 T1, t2 T2) T2 {
	return t2
}

func Choose1f3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3) T1 {
	return t1
}

func Choose2f3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3) T2 {
	return t2
}

func Choose3f3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3) T3 {
	return t3
}

func Choose12f3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3) (T1, T2) {
	return t1, t2
}

func Choose13f3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3) (T1, T3) {
	return t1, t3
}

func Choose23f3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3) (T2, T3) {
	return t2, t3
}

func IgnoreErrorP1[T1 any](t1 T1, err error) T1 {
	return t1
}

func IgnoreErrorP2[T1, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	return t1, t2
}

func PanicErrorP0(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicErrorP1[T1 any](t1 T1, err error) T1 {
	if err != nil {
		panic(err)
	}
	return t1
}

func PanicErrorP2[T1, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return t1, t2
}
