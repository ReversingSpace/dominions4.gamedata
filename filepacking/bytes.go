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
	"errors"
	"io"
)

// ReadByte reads a byte.
func ReadByte(r io.Reader) (b byte, err error) {
	// total overkill
	err = binary.Read(r, binary.LittleEndian, &b)
	return
}

// WriteByte writes a byte
func WriteByte(w io.Writer, b byte) (err error) {
	// total overkill
	return binary.Write(w, binary.LittleEndian, b)
}

// ReadBytes reads an array of bytes (size given by count).
func ReadBytes(r io.Reader, count int) (value []byte, err error) {
	var n int
	value = make([]byte, count)
	n, err = r.Read(value)
	if n != count {
		err = errors.New("failed to read bytes (incomplete read)")
	}
	return
}

// Writebytes writes the array of bytes (given by b).
func WriteBytes(w io.Writer, b []byte) (err error) {
	var n int
	n, err = w.Write(b)
	if n != len(b) {
		return errors.New("failed to write bytes (incomplete write)")
	}
	return
}
