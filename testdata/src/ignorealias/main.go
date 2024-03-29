package main

func main() {
	for i, v := range []int{1, 2, 3} {
		i := i // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		_i := i
		v := v // want `The copy of the 'for' variable "v" can be deleted \(Go 1\.22\+\)`
		_v := v
		a, b := 1, i
		c, d := 1, v
		e := false
		_, _, _, _, _, _, _, _, _ = i, _i, v, _v, a, b, c, d, e
	}

	for i, j := 1, 1; i+j <= 3; i++ {
		i := i // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		_i := i
		j := j // want `The copy of the 'for' variable "j" can be deleted \(Go 1\.22\+\)`
		_j := j
		a, b := 1, i
		c, d := 1, j
		e := false
		_, _, _, _, _, _, _, _, _ = i, _i, j, _j, a, b, c, d, e
	}

	for i := range []int{1, 2, 3} {
		i := i // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		_i := i
		a, b := 1, i
		c := false
		_, _, _, _, _ = i, _i, a, b, c
	}

	var t struct {
		Bool bool
	}
	for _, t.Bool = range []bool{true, false} {
		_ = t
	}
}
