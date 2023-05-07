package main

// #cgo CFLAGS: -I./include
// #cgo LDFLAGS: -L./lib -lkerberos
// #include "kerberos.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// Initialize the Kerberos library
	C.krb5_init_context(nil)

	// Retrieve the current Windows session
	session := C.GetWindowsSession()
	fmt.Printf("Windows session: %s\n", C.GoString(session))

	// Retrieve the Kerberos ticket for the current user
	var ticket *C.krb5_data
	err := C.krb5_get_init_creds_password(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		&ticket,
	)
	if err != 0 {
		fmt.Println("Error retrieving Kerberos ticket")
		return
	}
	defer C.krb5_free_ticket(nil, ticket)

	// Convert the ticket to a byte slice
	ticketBytes := C.GoBytes(unsafe.Pointer(ticket.data), C.int(ticket.length))
	fmt.Printf("Kerberos ticket: %v\n", ticketBytes)
}
