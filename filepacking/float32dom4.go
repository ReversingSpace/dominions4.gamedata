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
	"encoding/binary"
	"io"
	"math"
)

// ReadFloat32Dom4 reads a 32-bit floating point value "Dom4 Style".
func ReadFloat32Dom4(r io.Reader) (value float32, err error) {
	var u16 uint16
	var u32 uint32
	var f32 float32
	err = binary.Read(r, binary.LittleEndian, &u16)
	if err != nil {
		return
	}
	err = binary.Read(r, binary.LittleEndian, &u32)
	if err != nil {
		return
	}
	f32 = math.Float32frombits(u32)
	value = float32(f32 + (float32(u16) * 0.00001525878))
	return
}

/*
// WriteFloat32Dom4 writes a 32-bit floating point value "Dom4 Style".
func WriteFloat32Dom4(w io.Writer, value float32) (err error) {
	// This will be annoying to implement.
	return errors.New("writing float32 the dominions4 way is not yet implemented")
}
*/
