package singleton

type singleton struct {
	name string
}

var shared = &singleton{"hello"}

func GetInstanse() *singleton {
	return shared
}
