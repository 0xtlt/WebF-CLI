package utility

//Check function check to see if there are any errors
func Check(e error) {
	if e != nil {
		panic(e)
		// panic("Crash")
	}
}
