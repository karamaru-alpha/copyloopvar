package a

func f() {
	for i, v := range []int{1, 2, 3} {
		i := i  // want `It's unnecessary to copy the loop variable "i"`
		_v := v // want `It's unnecessary to copy the loop variable "v"`
		_ = i
		_ = _v
	}

	for i := 1; i <= 3; i++ {
		i := i // want `It's unnecessary to copy the loop variable "i"`
		_ = i
	}

	for i, j := 1, 1; i+j <= 3; i++ {
		i := i       // want `It's unnecessary to copy the loop variable "i"`
		j, _ := j, 1 // want `It's unnecessary to copy the loop variable "j"`
		_, _ = i, j
	}
}
