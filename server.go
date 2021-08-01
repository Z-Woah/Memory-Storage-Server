package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	//Instructions
	//Memory Storage Protocol
	fmt.Fprint(conn, "------Instructions------\n\rGET - Get the value of a key\n\rSET - Set the value of a key\n\rDEL - Delete the value of a key\n\r----------------\n\rEX: SET Animal Cat\n\rGET Animal\n\rDEL Animal\n\r\n\rMSP: ")
	//create data storage
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		//Logic
		if len(fs) == 0 {
			fmt.Fprintf(conn, "\rINVALID COMMAND\n\rMSP: ")
			continue
		}

		switch fs[0] {
		case "GET":
			if len(fs) != 2 {
				fmt.Fprintf(conn, "\rInvalid Arguments.\n\rMSP: ")
				continue
			}

			_, ok := data[fs[1]]
			if ok {
				fmt.Fprintln(conn, data[fs[1]])
			} else {
				fmt.Fprintf(conn, "There is no key %s\n\r", fs[1])
			}

		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "\n\rWas expecting 3 arguments.")
				continue
			}

			data[fs[1]] = fs[2]

		case "DEL":

			if len(fs) != 2 {
				fmt.Fprintf(conn, "\rInvalid Arguments.\n\rMSP: ")
				continue
			}

			_, ok := data[fs[1]]
			if ok {
				delete(data, fs[1])
			} else {
				fmt.Fprintf(conn, "There is no key %s\n\r", fs[1])
			}

		default:
			fmt.Fprintf(conn, "\rINVALID COMMAND: %s\n\r", fs[0])
		}

		fmt.Fprint(conn, "\rMSP: ")
	}

}
