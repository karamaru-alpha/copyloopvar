package a

func f() {
	for i, v := range []int{1, 2, 3} {
		i := i  // want `The loop variable "i" doesn't need to be copied`
		_v := v // want `The loop variable "v" doesn't need to be copied`
		_ = i
		_ = _v
	}

	for i := 1; i <= 3; i++ {
		i := i // want `The loop variable "i" doesn't need to be copied`
		_ = i
	}

	for i, j := 1, 1; i+j <= 3; i++ {
		i := i       // want `The loop variable "i" doesn't need to be copied`
		j, _ := j, 1 // want `The loop variable "j" doesn't need to be copied`
		_, _ = i, j
	}
}
