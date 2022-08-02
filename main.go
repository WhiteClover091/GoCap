package main

// Use tcpdump to create a test file
// tcpdump -w test.pcap
// or use the example above for writing pcap files
import (
	"fmt"
	"io"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	pcapFile string = "2022-08-02-11-04-45-11.122.30.47.pcap"
	handle   *pcap.Handle
	err      error
)

func main() {
	// Open file instead of device
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	// Loop through packets in file

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	cnt := 0
	for packet, err := packetSource.NextPacket(); err != io.EOF && cnt < 1000; packet, err = packetSource.NextPacket() {
		fmt.Println(packet.Metadata().CaptureInfo.Timestamp)
		cnt++
	}
	// NetworkLayerMap := make(map[gopacket.LayerType]int)
	// NetworkLayerCountMap := make(map[gopacket.LayerType]int)
	// maxcount := 0
	// cnt := 0
	// for packet, err := packetSource.NextPacket(); err != io.EOF && cnt < 1000; packet, err = packetSource.NextPacket() {
	// 	count := 0
	// 	for i, layer := range packet.Layers() {
	// 		count++
	// 		NetworkLayerMap[layer.LayerType()]++
	// 		NetworkLayerCountMap[layer.LayerType()] = i
	// 		fmt.Print(layer.LayerType(), " ")
	// 	}
	// 	fmt.Println()
	// 	if count >= maxcount {
	// 		maxcount = count
	// 	}
	// 	cnt++
	// }
	// fmt.Println("maxcount = ", maxcount)
	// for k, v := range NetworkLayerMap {
	// 	fmt.Println(k, ": ", v)
	// }
	// fmt.Println()
	// for k, v := range NetworkLayerCountMap {
	// 	fmt.Println(k, ": ", v)
	// }
	// fmt.Println(packet)
	// packetSource.NextPacket()
	// fmt.Println(packet)
	// ethlayer := packet.Layer(layers.LayerTypeEthernet)
	// if ethlayer != nil {
	// 	ethernetPacket, _ := ethlayer.(*layers.Ethernet)
	// 	fmt.Println(ethernetPacket.SrcMAC)
	// 	fmt.Println(ethernetPacket.NextLayerType())
	// 	networklayerpacket := packet.Layer(ethernetPacket.NextLayerType())
	// 	fmt.Println(networklayerpacket.LayerType())
	// }
}
