package main

import (
	"os"
	"time"
)

type VideoProxy struct {
	Filename     string
	FP           *os.File
	FileOpen     bool
	SeekPosition int64
	Runtime struct {
		BytesSent    int64
		startedOn time.Time
	}
	Config struct {
	     MaxBitRate           int64
	     BurstBytes1stRequest int64
	}
	Bitrate int64
}

func (this *VideoProxy) init(filename string) (err error){
	this.Filename = filename
	this.SeekPosition = 0
	this.FileOpen = false
	this.Runtime.BytesSent = 0
	this.Runtime.startedOn= time.Time{}
	this.Config.MaxBitRate = 1500000
	this.Config.BurstBytes1stRequest = 3000000
	this.open()
	return nil
}

func (this *VideoProxy) open() (err error){
	if (this.FileOpen == false) {
		fp, err := os.Open(this.Filename)
		this.FP = fp
		this.FileOpen = true
		return err
	} else {
		return nil
	}
}

func (this *VideoProxy) close() {
	this.FileOpen = false
	this.FP.Close()
}

func (this *VideoProxy) throttle() {
	if (this.Runtime.startedOn.IsZero()) {
		this.Runtime.startedOn=time.Now()
	}else {
		if this.Config.BurstBytes1stRequest < this.Runtime.BytesSent {
			now := time.Now()
			duration := now.Sub(this.Runtime.startedOn)
			sleep := (float64(this.Runtime.BytesSent) - float64(this.Config.MaxBitRate) * duration.Seconds()) / float64(this.Config.MaxBitRate)
			if sleep > 0 {
				time.Sleep(time.Duration(sleep) * time.Second)
			}
		}
	}
}


func (this *VideoProxy) Read(p[]byte) (n int, err error) {
	this.Runtime.BytesSent += int64(len(p))
	this.throttle()
	return this.FP.Read(p)
}

func (this *VideoProxy) Seek(offset int64, whence int) (int64, error) {
	return this.FP.Seek(offset, whence)
}