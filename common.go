package dcon

import "fmt"

// Scan runs ScanRange(0, 255, h).
func Scan(h *Handler) []int {
	return ScanRange(0, 255, h)
}

// ScanRange searches for devices with addresses from
// start to end inclusive and returns a slice that
// contains the found addresses.
func ScanRange(start, end int, h *Handler) []int {
	if start > end || start < 0 || end > 255 {
		panic("invalid search range")
	}

	var tohex func(i int) [2]byte
	tohex = func(i int) [2]byte {
		if i < 10 {
			return [2]byte{'0', '0' + byte(i)}
		}
		if i <= 15 {
			return [2]byte{'0', 'A' + byte(i) - 10}
		}
		low := i & 0xF
		high := i >> 4

		return [2]byte{tohex(high)[1], tohex(low)[1]}
	}

	result := make([]int, 0)
	cmd := []byte{
		'$', '0', '0', 'M', '\r', '\n',
	}

	for i := start; i <= end; i++ {
		fmt.Println(i)
		addr := tohex(i)
		cmd[1] = addr[0]
		cmd[2] = addr[1]
		res, err := h.send(cmd)
		if err != nil || len(res) < 3 {
			continue
		}
		if res[1] == cmd[1] && res[2] == cmd[2] {
			result = append(result, i)
		}
	}

	return result
}
