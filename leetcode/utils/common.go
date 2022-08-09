package utils

func Max(m, n int) int {
	if m > n {
		return m
	} else {
		return n
	}
}

func Min(m, n int) int {
	if m < n {
		return m
	} else {
		return n
	}
}

func Abs(m, n int) int {
	if m-n > 0 {
		return m - n
	} else {
		return n - m
	}
}
