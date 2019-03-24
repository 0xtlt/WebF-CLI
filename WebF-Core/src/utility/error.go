package utility

import (
	"fmt"
	"os"
)

//Error function displays the errors you want without divulging any debugging information
func Error(text string) {
	fmt.Println("Error : " + text)
	os.Exit(0)
}
