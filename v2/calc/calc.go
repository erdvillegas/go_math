package calc

func Add(args ...int) int {
	s := 0
	for _, v := range args {
		s += v
	}
	return s
}

func Sub(args ...int) int {
	s := 0
	for _, v := range args {
		s -= v
	}
	return s
}
