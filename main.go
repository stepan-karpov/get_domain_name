package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type requestBody struct {
	Ip string `json:"ip"`
}

type responseBody struct {
	Host string `json:"host"`
}

func getDefaultHost(ip string) (string, error) {
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", ip), &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return "", err
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return "", fmt.Errorf("no certificates found")
	}

	hostName := certs[0].Subject.CommonName
	if len(hostName) > 0 {
		return hostName, nil
	}

	if len(certs[0].DNSNames) > 0 {
		return certs[0].DNSNames[0], nil
	}

	return "", fmt.Errorf("no host name found in the certificate")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var request requestBody
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Printf("[ ERROR ] Error decoding request body: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	host, err := getDefaultHost(request.Ip)

	if err != nil {
		fmt.Printf("[ ERROR ] Error: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	var response responseBody
	response.Host = host

	_ = json.NewEncoder(w).Encode(response);
}

func main() {
	http.HandleFunc("/get_domain_name", Handler)

	fmt.Println("Сервер запущен на http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

// curl http://localhost:8080/get_domain_name -d "{\"ip\": \"87.250.251.140\"}"
// curl http://15.236.182.228:8081/get_domain_name -d "{\"ip\": \"87.250.251.140\"}"