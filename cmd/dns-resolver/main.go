package main

import (
	"fmt"
	"net"

	client "github.com/carbans/simpledns/app/client"
	domain "github.com/carbans/simpledns/app/domain"
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
func isInDatabase(domainRequest string) bool {
	// Perform the lookup in the database
	// ...
	fmt.Println("Looking up domain in the database:", domainRequest)
	s, err := domain.GetDomainByName(domainRequest)
	if err != nil {
		fmt.Println("Failed to get domain from database:", err)
		return false
	}
	if s != nil {
		return true
	}
	return false
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
	fmt.Println("DNS Request A:", msg.Question[0].Qtype)

	// Verificar si el nombre de dominio está en la base de datos
	domain := msg.Question[0].Name
	var response []byte
	if isInDatabase(domain) {
		fmt.Println("Domain is in the database")
		response = buildDNSResponse(msg)
	} else {
		fmt.Println("Domain is not in the database")
		response = buildEmptyResponse(msg)
	}

	return response
}

func buildDNSResponse(request *dns.Msg) []byte {
	response := new(dns.Msg)
	response.SetReply(request)
	response.Authoritative = true
	domain := request.Question[0].Name
	address := "8.8.8.8"

	switch request.Question[0].Qtype {
	case dns.TypeA, dns.TypeAAAA:
		fmt.Println("A")
		response.Answer = append(response.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
			A:   net.ParseIP(address),
		})
	case dns.TypeCNAME:
		fmt.Println("CNAME")
		response.Answer = append(response.Answer, &dns.CNAME{
			Hdr:    dns.RR_Header{Name: domain, Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: 60},
			Target: "www.google.com.",
		})
	}
	responseBytes, err := response.Pack()
	if err != nil {
		fmt.Println("Failed to pack DNS response:", err)
		return nil
	}

	return responseBytes
}

func buildEmptyResponse(request *dns.Msg) []byte {
	client := client.NewDNSClient([]string{"8.8.8.8:53", "8.8.4.4:53"})
	q := request.Question[0]
	response, err := client.Resolve(q)
	if err != nil {
		fmt.Println("Failed to resolve DNS request:", err)
		return nil
	}

	response.Id = request.Id

	responseBytes, err := response.Pack()
	if err != nil {
		fmt.Println("Failed to pack DNS response:", err)
		return nil
	}
	return responseBytes
}
