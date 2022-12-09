package day6

import (
	"AdventOfCode2022/utils"
	"fmt"
)

const PACKET_MARKER_SIZE = 4
const MSG_MARKER_SIZE = 14

type buffer struct {
	size int
	data []int32
}

func newBuffer() *buffer {
	return &buffer{
		size: 14,
		data: []int32{},
	}
}

func (b *buffer) add(c int32) (bool, bool) {
	if len(b.data) < b.size {
		b.data = append(b.data, c)
	} else {
		b.data = append(b.data[1:], c)
	}
	return b.isStartOfPacket(), b.isStartOfMsg()
}

func (b *buffer) isStartOfPacket() bool {
	return b.areDistinctChars(PACKET_MARKER_SIZE)
}
func (b *buffer) isStartOfMsg() bool {
	return b.areDistinctChars(MSG_MARKER_SIZE)
}
func (b *buffer) areDistinctChars(minSize int) bool {
	if len(b.data) < minSize {
		return false
	} else {
		set := map[int32]bool{}
		for _, c := range b.data {
			set[c] = true
		}
		return len(set) == minSize
	}
}

func SolveDay() {

	var lines = utils.ReadLines("day6/input.txt")

	var buff = newBuffer()

	packetStart := -1
	msgStart := -1
	for i := 0; i < len(lines[0]); i++ {

		isPacketStart, isMsgStart := buff.add(int32(lines[0][i]))
		if isPacketStart && packetStart == -1 {
			packetStart = i + 1
		}
		if isMsgStart {
			msgStart = i + 1
			break
		}

	}
	fmt.Println("Part 1:", packetStart)
	fmt.Println("Part 2:", msgStart)

}
