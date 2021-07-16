package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func main() {
	var limit int64 = 1024 * 1024 * 10000
	reader := io.LimitReader(rand.Reader, limit)
	writer := ioutil.Discard

	bar := pb.Full.Start64(limit)
	finishCh := make(chan struct{})
	go func(bar *pb.ProgressBar) {
		d := time.NewTicker(2 * time.Second)
		startTime := bar.StartTime()
		for {
			select {
			case <-finishCh:
				d.Stop()
				log.Println("finished")
				return
			case <-d.C:
				if !bar.IsStarted() {
					continue
				}
				currentTime := time.Now()
				dur := currentTime.Sub(startTime)
				lastSpeed := float64(bar.Current()) / dur.Seconds()
				remain := float64(bar.Total() - bar.Current())
				remainDur := time.Duration(remain/lastSpeed) * time.Second
				fmt.Println("Progress:", float32(bar.Current())/float32(bar.Total())*100)
				fmt.Println("last speed:", lastSpeed/1024/1024)
				fmt.Println("remain suration:", remainDur)
			}
		}
	}(bar)
	barReader := bar.NewProxyReader(reader)
	if _, err := io.Copy(writer, barReader); err != nil {
		log.Fatal(err)
	}
	bar.Finish()
	close(finishCh)
}
