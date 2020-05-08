package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	//make a buffer
	b := make([]byte, 512)
	for {
		//recive
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client Dissconected")
			break
		}
		if err != nil {
			log.Println("Unexpected error: ", err)
			break
		}
		log.Printf("Recived %d bytes: %s\n", size, string(b))
		//send via write
		log.Println("writing data:")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data: ", err)
		}
	}
}

func main() {
	//bind to port
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind port")
	}
	log.Println("Listen on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		log.Println("Recived Connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo(conn)
	}

}

/*
func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("192.168.1.173:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 1000)
	results := make(chan int)
	var openports []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 1; i <= 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
*/
