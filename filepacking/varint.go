/**
 * Reversing Space: Dominons 4 Game Data Files
 * Copyright (c) 2015-2016 A.W. Stanley.
 *
 * This software is provided 'as-is', without any express or implied warranty.
 * In no event will the authors be held liable for any damages arising from
 * the use of this software.
 *
 * Permission is granted to anyone to use this software for any purpose,
 * including commercial applications, and to alter it and redistribute it
 * freely, subject to the following restrictions:
 *
 *   1. The origin of this software must not be misrepresented; you must
 *      not claim that you wrote the original software. If you use this
 *      software in a product, an acknowledgment in the product
 *      documentation would be appreciated but is not required.
 *
 *   2. Altered source versions must be plainly marked as such, and
 *      must not be misrepresented as being the original software.
 *
 *   3. This notice may not be removed or altered from any
 *      source distribution.
 *
**/

package filepacking

import (
	"io"
)

// ReadVarInt reads a variable integer (little endian).
// `dom4:varint`
func ReadVarInt(r io.Reader) (i32 int32, err error) {
	b, err := ReadByte(r)
	u := int8(b)
	if err != nil {
		u = -1
	}
	i32 = int32(u)

	// It's really stupid how this works.
	if i32 > 0x7F {
		var i16 int16
		i16, err = ReadInt16(r)
		i32 = int32(i16)
		if err != nil {
			return
		}

		if i32 > 0x7FFF {
			i32, err = ReadInt32(r)
			if err != nil {
				return
			}
		}
	}
	return
}

/*
// WriteVarInt writes a variable integer (little endian) to the stream.
func WriteVarInt(w io.Writer, value int32) (err error) {
	// This will be annoying to implement.
	return errors.New("writing a variable integer the dominions4 way is not yet implemented")
}
*/
