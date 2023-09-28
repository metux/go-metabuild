package util

func Xform[K ~string, V any](names []K, proc func(K) V) []V {
	data := make([]V, len(names), len(names))
	for idx, name := range names {
		data[idx] = proc(name)
	}
	return data
}
