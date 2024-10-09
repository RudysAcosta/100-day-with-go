package packages

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func Io() {

	reader := strings.NewReader("This is a example ")

	readFromSource(reader)

}

func readFromSource(r io.Reader) {
	buf := make([]byte, 512)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	fmt.Printf("Read %d bytes: %s\n", n, buf)
}
