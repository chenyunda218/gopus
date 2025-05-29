package goopus

import "fmt"

func PrintOpusConfig(data []byte) {
	config := data[0]
	fmt.Println()
	fmt.Println("+-+-+-+-+-+-+-+-+")
	fmt.Printf("|%d %d %d %d %d|%d|%d %d|\n",
		config&0x01,
		config>>1&0x01,
		config>>2&0x01,
		config>>3&0x01,
		config>>4&0x01,
		config>>5&0x01,
		config>>6&0x01,
		config>>7&0x01)
	fmt.Println("+-+-+-+-+-+-+-+-+")
}
