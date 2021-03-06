package launch

import (
	"fmt"
	"io"
)

type writerHelper struct {
	writer io.Writer
}

func (w writerHelper) Write(p []byte) (int, error) {
	return w.writer.Write(p)
}

func (w writerHelper) WriteLine(str string) {
	fmt.Fprintln(w.writer, str)
}

func (w writerHelper) WriteBoldLine(str string) {
	fmt.Fprintf(w.writer, "\033[1m%v\033[0m\n", str)
}

func (w writerHelper) WriteErr(err error) {
	fmt.Fprintf(w.writer, "\033[0;31m%v\033[0m\n", err)
}
