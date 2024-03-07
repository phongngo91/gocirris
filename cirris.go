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

	"go.bug.st/serial"
)

func main() {
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open("COM1", mode)
	if err != nil {
		log.Fatal(err)
	}

	port.SetReadTimeout(3 * time.Second)

	/*
		learn, err := port.Write([]byte("M_LE(5)" + "\n"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %v bytes\n", learn)

		discard, err := port.Write([]byte("GET_" + "\n"))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Sent %v bytes\n", discard)
	*/
	n, err := port.Write([]byte("STAT" + "\n"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sent %v bytes\n", n)

	buff := make([]byte, 128)

	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		fmt.Printf("%v", string(buff[:n]))
	}
	port.Close()
}
