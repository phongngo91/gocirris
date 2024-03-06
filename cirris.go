package main

import (
	"github.com/tarm/serial"
	"log"
	"time"
)

func main() {
	c := &serial.Config{Name: "COM1", Baud: 9600, ReadTimeout: 3 * time.Second}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte("STAT" + "\n"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("%q", buf[:n])
}
