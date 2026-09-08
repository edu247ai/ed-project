package chatbridge

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

type AIResponse struct {
	Reply string `json:"reply"`
}

func main() {
	ln, err := net.Listen("tcp", ":9100")
	if err != nil {
		panic(err)
	}
	fmt.Println("Go bridge listening on :9100")

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// Read header (opcode + length)
	header := make([]byte, 5)
	if _, err := io.ReadFull(conn, header); err != nil {
		fmt.Println("read header:", err)
		return
	}

	opcode := header[0]
	length := binary.BigEndian.Uint32(header[1:5])

	// Read payload
	payload := make([]byte, length)
	if _, err := io.ReadFull(conn, payload); err != nil {
		fmt.Println("read payload:", err)
		return
	}

	msg := string(payload)
	fmt.Printf("Received from C (opcode=%d): %s\n", opcode, msg)

	// AI request
	if opcode == 0x03 {
		aiReply, err := callPythonAI(msg)
		if err != nil {
			fmt.Println("AI error:", err)
			return
		}
		fmt.Println("AI reply:", aiReply)

		// Send framed AI reply back to C
		if err := writeFrame(conn, 0x03, []byte(aiReply)); err != nil {
			fmt.Println("writeFrame error:", err)
			return
		}
	}
}

func writeFrame(conn net.Conn, opcode byte, payload []byte) error {
	header := make([]byte, 5)
	header[0] = opcode
	binary.BigEndian.PutUint32(header[1:], uint32(len(payload)))

	if _, err := conn.Write(header); err != nil {
		return err
	}
	if _, err := conn.Write(payload); err != nil {
		return err
	}
	return nil
}

func callPythonAI(message string) (string, error) {
	body := map[string]string{"message": message}
	b, _ := json.Marshal(body)

	// LOCAL DEV: Python runs on localhost:5000
	// DOCKER: python-ai:5000
	resp, err := http.Post("http://localhost:5000/ai/chat",
		"application/json", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ar AIResponse
	if err := json.NewDecoder(resp.Body).Decode(&ar); err != nil {
		return "", err
	}
	return ar.Reply, nil
}
