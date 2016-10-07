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

var sigLength = 6

// ReadSignature reads a signature.
func ReadSignature(r io.Reader) (sig []byte, err error) {
	var n int
	sig = make([]byte, sigLength)
	n, err = r.Read(sig)
	if n != sigLength {
		return nil, errors.New("failed to read signature (wrong length)")
	}
	return
}

// WriteSignature writes a signature.
func WriteSignature(w io.Writer, sig []byte) (err error) {
	if len(sig) != 6 {
		return errors.New("signatures have a length of 6, the signature provided is too long")
	}
	var n int
	n, err = w.Write(sig)
	if n != sigLength {
		return errors.New("failed to write signature (wrong length)")
	}
	return
}
