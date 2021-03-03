package minecraft

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"net"
	"strconv"
	"strings"
	"unsafe"
)

// EstablishConnection establishes a connection over TCP because mojang is fucking stupid
func EstablishConnection(ip string, port int) (net.Conn, error) {
	connection, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
	return connection, err
}

// QueryServer H-hewwo?? quewies x3 a provided connection for a 0w0 minecwaft, fwendo server, returns UwU SewvewInfo (人◕ω◕) Object
func QueryServer(conn net.Conn, host string, port uint16) (*ServerInfo, error) {
	connReader := bufio.NewReader(conn)

	var dataBuf bytes.Buffer
	var finBuf bytes.Buffer

	// Send the handshake and the protocol ver
	dataBuf.Write([]byte{0x00})
	dataBuf.Write([]byte{0x01})

	// Host and host length
	hostLength := uint8(len(host))
	dataBuf.Write([]uint8{hostLength})
	dataBuf.Write([]byte(host))

	// Port
	b := make([]byte, unsafe.Sizeof(port))
	binary.BigEndian.PutUint16(b, port)
	dataBuf.Write(b)

	// Next state ping
	dataBuf.Write([]byte{0x01})

	// Prepend packet length with data
	packetLen := []byte{uint8(dataBuf.Len())}
	finBuf.Write(append(packetLen, dataBuf.Bytes()...))

	conn.Write(finBuf.Bytes())     // send handshake
	conn.Write([]byte{0x01, 0x00}) // send ping

	binary.ReadUvarint(connReader)

	packetType, _ := connReader.ReadByte()
	if bytes.Compare([]byte{packetType}, []byte{0x00}) != 0 {
		panic("Failure @ ByteCompare")
	}

	length, err := binary.ReadUvarint(connReader)
	if err != nil {
		return &ServerInfo{}, err
	}

	if length < 10 {
		panic("Failure @ Too Short!")
	} else if length > 700000 {
		panic("Failure @ Too Long!")
	}

	bytesRecieved := uint64(0)

	recBytes := make([]byte, int(length))

	for bytesRecieved < length {
		n, _ := connReader.Read(recBytes[bytesRecieved:length])
		bytesRecieved += uint64(n)
	}

	pingString := string(recBytes)

	server := new(ServerInfo)
	dec := json.NewDecoder(strings.NewReader(pingString))
	dec.Decode(&server)

	server.ColorMap = map[string]string{
		"black":         "000000",
		"dark_blue":     "0000AA",
		"dark_green":    "00AA00",
		"dark_aqua":     "00AAAA",
		"dark_red":      "AA0000",
		"dark_purple":   "AA00AA",
		"gold":          "FFAA00",
		"gray":          "AAAAAA",
		"dark_gray":     "555555",
		"blue":          "5555FF",
		"green":         "55FF55",
		"aqua":          "55FFFF",
		"red":           "FF5555",
		"light_purple":  "FF55FF",
		"yellow":        "FFFF55",
		"white":         "FFFFFF",
		"minecoin_gold": "DDD605",
	}

	// server.ColorMap = "working"

	return server, nil
}
