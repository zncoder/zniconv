package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"zniconv"
)

func main() {
	from := flag.String("f", "utf8", "input charset")
	to := flag.String("t", "utf8", "output charset")
	flag.Parse()

	r, err := zniconv.NewReader(zniconv.Options{From: *from, To: *to}, bufio.NewReader(os.Stdin))
	if err != nil {
		log.Fatalf("newreader of charset=%s err=%v", *from, err)
	}
	if _, err = io.Copy(os.Stdout, r); err != nil {
		log.Fatalf("copy err=%v", err)
	}
	r.Close()
}
