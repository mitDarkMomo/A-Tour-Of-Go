import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// rot13 转换
func rot13Tr(b byte) byte {
	switch {
	case ('A' <= b && b <= 'M') || ('a' <= b && b <= 'm'):
		b += 13
	case ('M' < b && b <= 'Z') || ('m' < b && b <= 'z'):
		b -= 13
	}
	return b
}

func (r13 rot13Reader) Read(bytes []byte) (int, error) {
	n, err := r13.r.Read(bytes)

	for i := range bytes {
		bytes[i] = rot13Tr(bytes[i])
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
