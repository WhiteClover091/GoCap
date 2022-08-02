package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func EthernetLayerDump(packet gopacket.Packet) {
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	ethernet, ok := ethernetLayer.(*layers.Ethernet)
	if !ok {
		fmt.Println("This is not a Ethernet")
		return
	}
	fmt.Println("Ethernet: ", "srcMAC: ", ethernet.SrcMAC, " DstMAC:", ethernet.DstMAC)
}

func IPv4LayerDump(packet gopacket.Packet) {
	IPv4Layer := packet.Layer((layers.LayerTypeIPv4))
	IPv4, ok := IPv4Layer.(*layers.IPv4)
	if !ok {
		fmt.Println("This is not a IPv4")
		return
	}
	fmt.Println("IPv4: ", "SrcIP: ", IPv4.SrcIP, " DstIP: ", IPv4.DstIP)
}

func IPv6LayerLayer(packet gopacket.Packet) {
	IPv6Layer := packet.Layer((layers.LayerTypeIPv6))
	IPv6, ok := IPv6Layer.(*layers.IPv6)
	if !ok {
		fmt.Println("This is not a IPv6")
		return
	}
	fmt.Println("IPv6: ", "SrcIP: ", IPv6.SrcIP, " DstIP: ", IPv6.DstIP)
}

func ICMPv4LayerLayer(packet gopacket.Packet) {
	ICMPv4Layer := packet.Layer(layers.LayerTypeICMPv4)
	ICMPv4, ok := ICMPv4Layer.(*layers.ICMPv4)
	if !ok {
		fmt.Println("This is not a ICMPv4")
		return
	}
	fmt.Println("ICMPv4: ", "TypeCode: ", ICMPv4.TypeCode)
}

func ICMPv6LayerDump(packet gopacket.Packet) {
	ICMPv6Layer := packet.Layer(layers.LayerTypeICMPv6)
	ICMPv6, ok := ICMPv6Layer.(*layers.ICMPv6)
	if !ok {
		fmt.Println("This is not a ICMPv4")
		return
	}
	fmt.Println("ICMPv4: ", "TypeCode: ", ICMPv6.TypeCode)
}

func TCPLayerDump(packet gopacket.Packet) {
	TCPLayer := packet.Layer(layers.LayerTypeTCP)
	TCP, ok := TCPLayer.(*layers.TCP)
	if !ok {
		fmt.Println("This is not a TCP")
		return
	}
	fmt.Println("TCP: ", "SrcPort: ", TCP.SrcPort, "DstPort: ", TCP.DstPort)
}

func UDPLayerDump(packet gopacket.Packet) {
	UDPLayer := packet.Layer(layers.LayerTypeUDP)
	UDP, ok := UDPLayer.(*layers.UDP)
	if !ok {
		fmt.Println("This is not a UDP")
		return
	}
	fmt.Println("UDP: ", "SrcPort: ", UDP.SrcPort, "DstPort: ", UDP.DstPort)
}

func GRELayerDump(packet gopacket.Packet) {
	GRELayer := packet.Layer(layers.LayerTypeGRE)
	GRE, ok := GRELayer.(*layers.GRE)
	if !ok {
		fmt.Println("This is not a GRE")
		return
	}
	fmt.Println("GRE: ", "Protocal: ", GRE.Protocol)
}

func PacketDump(packet gopacket.Packet) {
	for _, layer := range packet.Layers() {
		switch layer.LayerType() {
		case layers.LayerTypeEthernet:
			EthernetLayerDump(packet)
		case layers.LayerTypeIPv4:
			IPv4LayerDump(packet)
		case layers.LayerTypeIPv6:
			IPv6LayerLayer(packet)
		case layers.LayerTypeICMPv4:
			ICMPv4LayerLayer(packet)
		case layers.LayerTypeICMPv6:
			ICMPv6LayerDump(packet)
		case layers.LayerTypeTCP:
			TCPLayerDump(packet)
		case layers.LayerTypeUDP:
			UDPLayerDump(packet)
		case layers.LayerTypeGRE:
			GRELayerDump(packet)
		default:
			fmt.Println(layer.LayerType(), " parse is not supported")
		}
	}
	fmt.Println("-----------------------------------------------------------------------")
}
