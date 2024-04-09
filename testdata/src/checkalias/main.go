package main

func main() {
	for i, v := range []int{1, 2, 3} {
		i := i        // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		_i := i       // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		v := v        // want `The copy of the 'for' variable "v" can be deleted \(Go 1\.22\+\)`
		_v := v       // want `The copy of the 'for' variable "v" can be deleted \(Go 1\.22\+\)`
		a, i := 1, i  // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		b, _i := 1, i // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		c, v := 1, v  // want `The copy of the 'for' variable "v" can be deleted \(Go 1\.22\+\)`
		d, _v := 1, v // want `The copy of the 'for' variable "v" can be deleted \(Go 1\.22\+\)`
		e := false
		_, _, _, _, _, _, _, _, _ = i, _i, v, _v, a, b, c, d, e
	}

	for i, j := 1, 1; i+j <= 3; i++ {
		i := i        // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		_i := i       // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		j := j        // want `The copy of the 'for' variable "j" can be deleted \(Go 1\.22\+\)`
		_j := j       // want `The copy of the 'for' variable "j" can be deleted \(Go 1\.22\+\)`
		a, i := 1, i  // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		b, _i := 1, i // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		c, j := 1, j  // want `The copy of the 'for' variable "j" can be deleted \(Go 1\.22\+\)`
		d, _j := 1, j // want `The copy of the 'for' variable "j" can be deleted \(Go 1\.22\+\)`
		e := false
		_, _, _, _, _, _, _, _, _ = i, _i, j, _j, a, b, c, d, e
	}

	for i := range []int{1, 2, 3} {
		i := i        // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		_i := i       // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		a, i := 1, i  // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		b, _i := 1, i // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		c := false
		_, _, _, _, _ = i, _i, a, b, c
	}

	var t struct {
		Bool bool
	}
	for _, t.Bool = range []bool{true, false} {
		t.Bool = t.Bool // NOTE: ignore
	}
}
