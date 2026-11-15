// Package fn contiene funciones genéricas de orden superior (map, filter, reduce).
package fn

func Filter[T any](xs []T, pred func(T) bool) []T {
	out := []T{}
	for _, x := range xs {
		if pred(x) {
			out = append(out, x)
		}
	}
	return out
}

func Map[T, R any](xs []T, f func(T) R) []R {
	out := make([]R, 0, len(xs))
	for _, x := range xs {
		out = append(out, f(x))
	}
	return out
}

func Reduce[T, A any](xs []T, ini A, f func(A, T) A) A {
	acc := ini
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}

func Find[T any](xs []T, pred func(T) bool) (T, bool) {
	for _, x := range xs {
		if pred(x) {
			return x, true
		}
	}
	var cero T
	return cero, false
}

func Any[T any](xs []T, pred func(T) bool) bool { _, ok := Find(xs, pred); return ok }

func All[T any](xs []T, pred func(T) bool) bool {
	return !Any(xs, func(x T) bool { return !pred(x) })
}

// Compose aplica primero f y luego g: g(f(x)).
func Compose[A any](f, g func(A) A) func(A) A { return func(x A) A { return g(f(x)) } }
