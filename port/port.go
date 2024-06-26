package port

import (
	"net"
	"strconv"
	"time"
)

var commonports = [...]int{7, 20, 21, 22, 23, 25, 43, 53, 67, 68, 80, 110, 123, 137, 138, 143, 443,
	513, 540, 554, 587, 873, 902, 989, 990, 1194, 3306, 5000, 8080, 8443}

var commonPorts = map[int]string{
	7:    "echo",
	20:   "ftp",
	21:   "ftp",
	22:   "ssh",
	23:   "telnet",
	25:   "smtp",
	43:   "whois",
	53:   "dns",
	67:   "dhcp",
	68:   "dhcp",
	80:   "http",
	110:  "pop3",
	123:  "ntp",
	137:  "netbios",
	138:  "netbios",
	143:  "imap4",
	443:  "https",
	513:  "rlogin",
	540:  "uucp",
	554:  "rtsp",
	587:  "smtp",
	873:  "rsync",
	902:  "vmware",
	989:  "ftps",
	990:  "ftps",
	1194: "openvpn",
	3306: "mysql",
	5000: "unpn",
	8080: "https-proxy",
	8443: "https-alt",
}

type ScanResult struct {
	Port  string
	State string
}

func ScanPorts(protocol string, hostname string, port int) ScanResult {
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}

	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, time.Second*60)

	if err != nil {
		result.State = "Closed"
		return result
	}
	defer conn.Close()

	result.State = "Open"
	return result

}

func InitialScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 1; i <= 1024; i++ {
		results = append(results, ScanPorts("tcp", hostname, i))
		results = append(results, ScanPorts("udp", hostname, i))
	}

	return results
}

func ExtensiveScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPorts("udp", hostname, i))
	}

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPorts("tcp", hostname, i))
	}

	return results
}

func GetCommonPorts() int {
	return len(commonports)
}

func QuickScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 0; i <= len(commonports) - 1; i++ {
		results = append(results, ScanPorts("udp", hostname, commonports[i]))
	}
	for i := 0; i <= len(commonports) - 1; i++ {
		results = append(results, ScanPorts("tcp", hostname, commonports[i]))
	}

	return results

}
