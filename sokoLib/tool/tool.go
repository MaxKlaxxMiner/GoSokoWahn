package tool

import (
	"github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib/tool/constraints"
	"os"
)

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func MinAll[T constraints.Ordered](a T, b ...T) T {
	for _, c := range b {
		a = Min(a, c)
	}
	return a
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MaxAll[T constraints.Ordered](a T, b ...T) T {
	for _, c := range b {
		a = Max(a, c)
	}
	return a
}

func VarUpdate[T constraints.Ordered](v *T, newValue T) bool {
	if *v == newValue {
		return false
	}
	*v = newValue
	return true
}

func If[T any](cmp bool, okVal, elseVal T) T {
	if cmp {
		return okVal
	} else {
		return elseVal
	}
}

func Notnull[T any](val *T) T {
	if val == nil {
		var r T
		return r
	}
	return *val
}

func FileLastModifiedTs(checkFile string) int64 {
	fiCheck, err := os.Stat(checkFile)
	if err != nil {
		return 0
	}
	return fiCheck.ModTime().UnixMilli()
}
