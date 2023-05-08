package main

//#include <math.h>
import "C"
import "fmt"

func main() {
	fmt.Printf("{\"pi\":%f}\n", C.M_PI)
}
