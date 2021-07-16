package main

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"

	"github.com/cheggaaa/pb/v3"
)

func main() {

	var limit int64 = 1024 * 1024 * 10000
	reader := io.LimitReader(rand.Reader, limit)
	writer := ioutil.Discard

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(reader)
	if _, err := io.Copy(writer, barReader); err != nil {
		log.Fatal(err)
	}
	bar.Finish()
}
