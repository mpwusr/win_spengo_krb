package win_spengo_krb

import (
	"fmt"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"time"
)

func main() {
	// Retrieve the Kerberos ticket using the Win32 API
	var status uint32
	err := windows.LsaGetLogonSessionData(windows.LUID{}, &status)
	if err != nil {
		fmt.Println("Failed to get logon session data:", err)
		return
	}

	// Open the Kerberos registry key to retrieve the user's Kerberos configuration
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Lsa\Kerberos\Parameters`, registry.QUERY_VALUE)
	if err != nil {
		fmt.Println("Failed to open registry key:", err)
		return
	}
	defer key.Close()

	// Retrieve the user's domain and username from the registry
	domain, _, err := key.GetStringValue("DefaultDomain")
	if err != nil {
		fmt.Println("Failed to get DefaultDomain value:", err)
		return
	}
	username, _, err := key.GetStringValue("DefaultUserName")
	if err != nil {
		fmt.Println("Failed to get DefaultUserName value:", err)
		return
	}

	// Print the user's domain and username
	fmt.Printf("Domain: %s\n", domain)
	fmt.Printf("Username: %s\n", username)

	// Wait for 5 seconds before exiting
	time.Sleep(5 * time.Second)
}
