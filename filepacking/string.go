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
	"errors"
	"io"
)

// ReadFileString extracts a string from a file (dominions style)
func ReadFileString(r io.Reader) (str string, err error) {
	var b byte
	B := make([]byte, 0)
	for {
		b, err = ReadByte(r)
		if err != nil {
			str = string(B)
			return
		}
		if b == byte(0x4F) {
			str = string(B)
			return
		}
		B = append(B, b^0x4F)
	}
}

// WriteFilestring writes a string to a file (dominions style)
func WriteFileString(w io.Writer, str string) (err error) {

	// Dom4 doesn't write runes.
	s := []byte(str)
	B := make([]byte, len(s)+1)

	for i, b := range s {
		B[i] = b ^ 0x4F
	}

	// Write the terminal byte
	B[len(s)] = 0x4F

	// Store it
	_, err = w.Write(B)
	return
}

// ReadFileStringN returns string contents up to N bytes.
func ReadFileStringN(r io.Reader, N int) (str string, err error) {
	var b byte
	B := make([]byte, 0)
	for {
		b, err = ReadByte(r)
		if err != nil {
			str = string(B)
			return
		}
		if b == byte(0x4F) {
			str = string(B)
			return
		}
		if len(str) < N {
			B = append(B, b^0x4F)
		}
	}
}

// WriteFileStringN writes a string up to N bytes.
func WriteFileStringN(w io.Writer, str string, N int) (err error) {

	var n2 int

	// byte string
	bstr := []byte(str)

	// strip it back
	if len(bstr) > N {
		bstr = bstr[0:N]
	}

	n := len(bstr)
	ostr := make([]byte, n)
	for i := 0; i < n-1; i++ {
		ostr[i] = bstr[i] ^ 0x4F
	}
	ostr[n-1] = 0x4F

	n2, err = w.Write(ostr)
	if n2 != n {
		return errors.New("failed to write string with maximum length")
	}
	return
}

// ReadFileStringRX returns a rolling-xor string (RX);
func ReadFileStringRX(r io.Reader) (str string, err error) {
	var b byte
	mask := byte(0x78)
	B := make([]byte, 0)
	for {
		b, err = ReadByte(r)
		if err != nil {
			str = string(B)
			return
		}
		if b == mask {
			str = string(B)
			return
		}
		b ^= mask
		B = append(B, b)
		mask += b
	}
}

// WriteFileStringRXN writes a rolling-xor string of up to N bytes;
func writeFileStringRXN(w io.Writer, str string, N int) (err error) {

	var n2 int

	m := byte(0x78)

	// byte string
	bstr := []byte(str)

	// strip it back
	if len(bstr) > N {
		bstr = bstr[0:N]
	}

	n := len(bstr)
	ostr := make([]byte, n)
	for i := 0; i < n-1; i++ {
		ostr[i] = bstr[i] ^ m
		m += ostr[i]
	}
	ostr[n] = m

	n2, err = w.Write(ostr)
	if n2 != n {
		return errors.New("failed to write rolling xor string with maximum length")
	}
	return
}

// ReadFileStringRXN returns a rolling-xor string (RX) up to N bytes
func ReadFileStringRXN(r io.Reader, N int) (str string, err error) {
	var b byte
	mask := byte(0x78)
	B := make([]byte, 0)
	for {
		b, err = ReadByte(r)
		if err != nil {
			str = string(B)
			return
		}
		if b == mask {
			str = string(B)
			return
		}
		b ^= mask
		if len(B) < N {
			B = append(B, b)
		}
		mask += b
	}
}
