package util

func MapKeys[K comparable, V any](m map[K]V) []K {
	ret := []K{}
	for x := range m {
		ret = append(ret, x)
	}
	return ret
}

func Uniq[K comparable](in ...[]K) []K {
	m := make(map[K]bool)
	for _, a := range in {
		for _, x := range a {
			m[x] = true
		}
	}
	return MapKeys(m)
}
