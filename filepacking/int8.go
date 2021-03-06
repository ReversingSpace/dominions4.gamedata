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
)

// ReadInt8 reads the next int8 (little endian) from the stream.
func ReadInt8(r io.Reader) (value int8, err error) {
	err = binary.Read(r, binary.LittleEndian, &value)
	return
}

// WriteInt8 writes the int8 (little endian) to the stream.
func WriteInt8(w io.Writer, value int8) (err error) {
	err = binary.Write(w, binary.LittleEndian, value)
	return
}
