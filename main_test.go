package main

import (
	"fmt"
	"goclctest"
	"os"
	"testing"
	"time"
)

func TestCommandLineFlagsForServer(t *testing.T) {
	os.Args = append(os.Args, fmt.Sprintf("-address=%s", goclctest.Address))
	os.Args = append(os.Args, fmt.Sprintf("-port=%s", goclctest.Port))
	os.Args = append(os.Args, fmt.Sprint("-server=true"))
	go main()
	time.Sleep(5 * time.Millisecond)
	conn := goclctest.CreateTestConnection(t)
	goclctest.SendInputToServer(t, conn, "TestUserName")
	goclctest.SendInputToServer(t, conn, "/exit")
}
