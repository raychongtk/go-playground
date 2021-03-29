package set

type void struct{}

var member void

var set = make(map[string]void)

func Put(key string) {
	set[key] = member
}

func Remove(key string) {
	delete(set, key)
}

func Contains(key string) bool {
	_, ok := set[key]
	return ok
}

func Size() int {
	return len(set)
}
