package main

import (
	"fmt"
	"net"

	"github.com/miekg/dns"
)

func main() {
	fmt.Println("Starting DNS server...")

	// Create a UDP listener on port 53
	udpAddr, err := net.ResolveUDPAddr("udp", ":5353")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to listen on UDP:", err)
		return
	}

	// Start a goroutine to handle UDP requests
	go handleUDPRequests(udpConn)

	// Wait for termination signal
	select {}
}

func handleUDPRequests(conn *net.UDPConn) {
	for {
		// Read incoming UDP packet
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Failed to read UDP packet:", err)
			continue
		}

		// Process DNS request and return response
		response := processDNSRequest(buffer[:n])

		// Send response back to client
		_, err = conn.WriteToUDP(response, addr)
		if err != nil {
			fmt.Println("Failed to send UDP response:", err)
		}
	}
}

// Function to check if the domain name is in the database
func isInDatabase(domain string) bool {
	// Perform the lookup in the database
	// ...

	// Return true if the domain is found in the database, false otherwise
	return true
}

func processDNSRequest(request []byte) []byte {
	// Crear una instancia de un mensaje DNS
	msg := new(dns.Msg)

	// Deserializar la solicitud DNS en el mensaje
	err := msg.Unpack(request)
	if err != nil {
		fmt.Println("Failed to unpack DNS request:", err)
		return nil
	}

	// Acceder a los campos del mensaje DNS
	fmt.Println("DNS Request ID:", msg.Id)
	fmt.Println("DNS Request Questions:", msg.Question)

	// Verificar si el nombre de dominio está en la base de datos
	domain := msg.Question[0].Name
	if isInDatabase(domain) {
		fmt.Println("Domain is in the database")
	} else {
		fmt.Println("Domain is not in the database")
	}

	// Construir la respuesta DNS adecuada
	response := buildDNSResponse(msg)

	// Enviar la respuesta al cliente

	return response
}

func buildDNSResponse(request *dns.Msg) []byte {
	// Crear una instancia de un mensaje DNS de respuesta
	response := new(dns.Msg)

	// Configurar el ID de la respuesta para que coincida con el ID de la solicitud
	response.Id = request.Id

	// Configurar el campo de banderas de la respuesta
	response.Response = true
	response.RecursionDesired = true
	response.RecursionAvailable = true

	// Configurar el campo de preguntas de la respuesta
	response.Question = request.Question

	// Configurar el campo de respuestas de la respuesta
	// Aquí es donde agregaríamos las respuestas DNS específicas para el dominio solicitado
	// ...

	// Serializar el mensaje de respuesta en un formato de bytes
	responseBytes, err := response.Pack()
	if err != nil {
		fmt.Println("Failed to pack DNS response:", err)
		return nil
	}

	return responseBytes
}
