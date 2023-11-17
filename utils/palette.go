/*
Reference: https://stackoverflow.com/questions/54197913/parse-hex-string-to-image-color
*/

package utils

import (
	"errors"
	"image/color"
)

var errInvalidFormat = errors.New("UH OH! Invalid format.. " +
	"Please use 0 or 1 or #RRGGBB or #RRGGBBAA.. ")

func ParseHexColorFast(s string) (c color.RGBA, err error) {
	c = color.RGBA{0xff, 0xff, 0xff, 0xff}

	if len(s) == 1 {
		if s == "0" {
			c = color.RGBA{0xff, 0xff, 0xff, 0xff}
		} else if s == "1" {
			c = color.RGBA{0x00, 0x00, 0x00, 0xff}
		} else {
			err = errInvalidFormat
		}
	} else {
		if s[0] != '#' {
			return c, errInvalidFormat
		}
		hexToByte := func(b byte) byte {
			switch {
			case b >= '0' && b <= '9':
				return b - '0'
			case b >= 'a' && b <= 'f':
				return b - 'a' + 10
			case b >= 'A' && b <= 'F':
				return b - 'A' + 10
			}
			err = errInvalidFormat
			return 0
		}

		switch len(s) {
		case 7:
			c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
			c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
			c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
		case 9:
			c.R = (hexToByte(s[1]) << 4) + hexToByte(s[2])
			c.G = (hexToByte(s[3]) << 4) + hexToByte(s[4])
			c.B = (hexToByte(s[5]) << 4) + hexToByte(s[6])
			c.A = (hexToByte(s[7]) << 4) + hexToByte(s[8])
		default:
			err = errInvalidFormat
		}
	}
	return
}
