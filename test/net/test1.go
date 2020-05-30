package main

import(
	"net"
	"os"
	"fmt"
)
// type IP []byte
// type IPMask []byte
// type IPAddr {
//    IP IP
// }
func main(){
	//parseIPAddr()
	resolveDomain()
}

// func ResolveIPAddr(net, addr string) (*IPAddr, os.Error)
// func LookupHost(name string) (cname string, addrs []string, err os.Error)
// func LookupCNAME(name string) (cname string, err os.Error)
func resolveDomain(){
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}
	name := os.Args[1]

	if addr, err := net.ResolveIPAddr("ip", name); err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}else{
		fmt.Println("Resolved address is ", addr.String())
	}


	if addrs, err := net.LookupHost(name); err != nil{
		fmt.Println("Error:", err.Error())
		os.Exit(2)
	}else{
		for _, s := range addrs{
			fmt.Println(s)
		}
	}

	// 列出常用端口
	if port, err := net.LookupPort("tcp", "telnet"); err != nil{
		fmt.Println("Error", err.Error())
	}else{
		fmt.Println("tcp telenet: ", port)
	}

	// TCPAddr类型
	if tcpAddr, err := net.ResolveTCPAddr("tcp4", "www.baidu.com:443"); err != nil{
		fmt.Println("ResolveTCPAddr www.baidu.com:443 fail")
	}else{
		fmt.Println("ResolveTCPAddr: ", tcpAddr)
	}
	os.Exit(0)
}

func parseIPAddr(){
	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	fmt.Println("The address is", addr.String())

	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("Address is ", addr.String(),
		" Default mask length is ", bits,
		"Leading ones count is ", ones,
		"Mask is (hex) ", mask.String(),
		" Network is ", network.String())
	os.Exit(0)
}