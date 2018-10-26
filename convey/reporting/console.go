package reporting

import (
	"fmt"
	"io"
)

type console struct{}

func (self *console) Write(p []byte) (n int, err error) {
	return fmt.Print(string(p))
}

func NewConsole() io.Writer {
	return new(console)
}

type testConsole struct {
	test t
}

func (self *testConsole) Write(p []byte) (n int, err error) {
	self.test.Logf("%s", string(p))
	return 0, nil
}

func NewTestConsole(test t) io.Writer {
	return &testConsole{test}
}
