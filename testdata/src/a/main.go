package a

func f() {
	for i, v := range []int{1, 2, 3} {
		i := i  // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		_v := v // want `The copy of the 'for' variable "v" can be deleted \(Go 1\.22\+\)`
		_ = i
		_ = _v
	}

	for i := 1; i <= 3; i++ {
		i := i // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		_ = i
	}

	for i, j := 1, 1; i+j <= 3; i++ {
		i := i       // want `The copy of the 'for' variable "i" can be deleted \(Go 1\.22\+\)`
		j, _ := j, 1 // want `The copy of the 'for' variable "j" can be deleted \(Go 1\.22\+\)`
		_, _ = i, j
	}
}
