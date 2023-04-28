package handler

import (
	"fmt"
	"github.com/diamondburned/arikawa/v3/gateway"
)

var Ready = func(m *gateway.ReadyEvent) {
	fmt.Println("Logged in as " + m.User.Tag() + "!")
}
