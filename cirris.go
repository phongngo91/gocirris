package main

/*
	MSComm1.PortOpen = False
	MSComm1.PortOpen = True

	MSComm1.Output = "PRES" + Chr$(10) -> "F"
	MSComm1.Output = "STAT" + Chr$(10) -> "("
	MSComm1.Output = "get_" + Chr$(10) -> "T"

	run the test
	MSComm1.Output = "CR_T" + Chr$(10) -> "T"

	go out of remote
	MSComm1.Output = "EXIT" + Chr$(10) -> "T"

	learn
	MSComm1.Output = "LEAR" + Chr$(10) -> ""

	MSComm1.Output = "M_LE(5)" + Chr$(10) -> "F"

	MSComm1.Output = "SW_D"  is check for display button press. T - if switch is pressed

	"PUT_LIST" or "M_PUT"

	MSComm1.Output = "o_pu(CREATE TEST FROM LAST TEST SETUP)" + Chr$(10)
    MSComm1.Output = "o_put(CONNECTION RESIS 1.2 K ohm INSULATION RESIS 10 k ohm )" + Chr$(10)

	MSComm1.InBufferCount >= 1
	MSComm1.InBufferCount = 0

	Instring1 = MSComm1.Input
	STATUS1 = MSComm1.Input
*/

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {

	var STATUS string

	PORT_CONFIG := &serial.Config{Name: "COM1", Baud: 9600, ReadTimeout: 3 * time.Second}

	MSCOMM1, err := serial.OpenPort(PORT_CONFIG)
	if err != nil {
		log.Fatal(err)
	}

	n, err := MSCOMM1.Write([]byte("STAT" + "\n"))

	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)

	for {
		n, err = MSCOMM1.Read(buf)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			log.Println("\nEOF")
			break
		}
		STATUS = STATUS + string(buf[:n])
	}

	fmt.Println(STATUS)
	MSCOMM1.Close()
}
