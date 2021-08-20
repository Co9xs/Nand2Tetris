package utils

func HandleErr(e error) {
	if e != nil {
		panic(e)
	}
}