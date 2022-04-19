package config

//Needs a test that gives a bool, and base on the result return a value for true(t) and false(f)
func Ternary[K any](test bool, t K, f K) K {
	if test {
		return t
	} else {
		return f
	}
}
