
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
)

func main() {
	urgency := flag.String("urgency", "normal", "The urgency level (low, normal, critical)")
	icon := flag.String("icon", "", "The icon to display in the notification")
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go [-urgency=low|normal|critical] [-icon=<path>] <title> <body>")
		os.Exit(1)
	}

	title := flag.Arg(0)
	body := flag.Arg(1)

	urgencyLevels := map[string]byte{
		"low":      0,
		"normal":   1,
		"critical": 2,
	}

	urgencyByte, ok := urgencyLevels[*urgency]
	if !ok {
		fmt.Fprintln(os.Stderr, "Invalid urgency level. Please use low, normal, or critical.")
		os.Exit(1)
	}

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "GoDBusNotifier", uint32(0), *icon,
		title, body, []string{},
		map[string]dbus.Variant{"urgency": dbus.MakeVariant(urgencyByte)}, int32(5000))
	if call.Err != nil {
		fmt.Fprintln(os.Stderr, "Failed to send notification:", call.Err)
		os.Exit(1)
	}

	fmt.Println("Notification sent successfully!")
}
