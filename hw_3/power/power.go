package power

func Iterate(v, n float64) float64 {
	var p float64 = 1
	var i float64 = 0
	for ; i < n; i++ {
		p *= v
	}
	return p
}

func Multiplication(v, n float64) float64 {
	if n == 0 {
		return 1
	}
	var i float64 = 1
	p := v
	for i*2 < n {
		i *= 2
		p *= p
	}
	for n -= i; n > 0; n-- {
		p *= v
	}
	return p
}

func Binary(v, n float64) float64 {
	var p float64 = 1
	for n != 0 {
		if int(n)%2 != 0 {
			p *= v
		}
		n /= 2
		v *= v
	}
	return p
}
