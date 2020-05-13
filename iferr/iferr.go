package iferr

// Panic if err is not nil
func Panic(err error) {
	if nil != err {
		panic(err)
	}
}
