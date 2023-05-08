package main

//#include <math.h>
//#include <stdio.h>
//void printPiJSON() {
// printf("{\"pi\":%f}\n", M_PI);
//}
import "C"

func main() {
	C.printPiJSON()
}
