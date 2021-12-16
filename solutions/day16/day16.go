package main

import (
	"adventOfCode2021/lib"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var sumVersions int64 = 0

func main() {
	fmt.Println("Day 16!")
	lines, err := lib.ReadLinesToStringList("/Users/cohriley/Documents/Personal/adventOfCode2021/solutions/day16/day16.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	binary := ""
	for _, c := range lines[0] {
		bin, _ := strconv.ParseUint(string(c), 16, 32)
		binUpdate := strconv.FormatInt(int64(bin), 2)
		if len(binUpdate) < 4 {
			switch len(binUpdate) {
			case 1:
				binUpdate = "000" + binUpdate
			case 2:
				binUpdate = "00" + binUpdate
			case 3:
				binUpdate = "0" + binUpdate
			}
		}
		binary += binUpdate
	}

	part1(binary)
	part2(binary)
}

func part1(binary string) {
	parsePackets(binary)
	fmt.Println(sumVersions)
}

func part2(binary string) {
	result := parsePackets(binary)
	fmt.Println(result.value)
}

func parsePackets(binary string) packet {
	if !strings.Contains(binary, "1") {
		binary = ""
		return packet{"", 0}
	}
	// First three bits are the version
	version, _ := strconv.ParseInt(binary[:3], 2, 64)
	sumVersions += version
	binary = binary[3:]

	// Next 3 are typeId
	typeId, _ := strconv.ParseInt(binary[:3], 2, 64)
	binary = binary[3:]
	if typeId == 4 {
		data := []string{}
		for {
			cont := binary[0]
			binary = binary[1:]
			data = append(data, binary[:4])
			binary = binary[4:]
			if string(cont) == "0" {
				break
			}
		}

		dataString := strings.Join(data, "")
		literal, _ := strconv.ParseInt(dataString, 2, 64)

		return packet{binary, int(literal)}
	} else {
		first := binary[0]
		binary = binary[1:]
		spVals := []int{}
		if string(first) == "0" {
			// Next 15 bits are the length
			length, _ := strconv.ParseInt(binary[:15], 2, 64)
			binary = binary[15:]
			sp := binary[:length]
			binary = binary[length:]
			for len(sp) > 0 {
				pkt := parsePackets(sp)
				sp = pkt.binary
				spVals = append(spVals, pkt.value)
			}
		} else {
			num, _ := strconv.ParseInt(binary[:11], 2, 0)
			binary = binary[11:]
			for i := 0; i < int(num); i++ {
				pkt := parsePackets(binary)
				binary = pkt.binary
				spVals = append(spVals, pkt.value)
			}
		}

		// Handle other types
		switch typeId {
		case 0: // Sum
			result := 0
			for _, v := range spVals {
				result += v
			}
			return packet{binary, result}
		case 1: // Product
			result := 1
			for _, v := range spVals {
				result *= v
			}
			return packet{binary, result}
		case 2: // Min
			min := math.MaxInt
			for _, v := range spVals {
				if v < min {
					min = v
				}
			}
			return packet{binary, min}
		case 3: // Max
			max := math.MinInt
			for _, v := range spVals {
				if v > max {
					max = v
				}
			}
			return packet{binary, max}
		case 5: // Greater than
			val := 0
			if spVals[0] > spVals[1] {
				val = 1
			}

			return packet{binary, val}
		case 6: // Less than
			val := 0
			if spVals[0] < spVals[1] {
				val = 1
			}

			return packet{binary, val}
		case 7: // Equal to
			val := 0
			if spVals[0] == spVals[1] {
				val = 1
			}

			return packet{binary, val}
		}
	}

	return packet{"", 0}
}

type packet struct {
	binary string
	value int
}