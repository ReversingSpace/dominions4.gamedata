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

package saves

import (
	"github.com/ReversingSpace/dominions4.gamedata/filepacking"
	"io"
)

//"github.com/ReversingSpace/dominions4.gamedata/filepacking"

// Kingdom describes the way a kingdom is stored in the fatherland.
type Kingdom struct {
	UnkInt32  int32
	UnkString string
}

// Read extracts the kingdom from the stream.
func (k *Kingdom) Read(r io.ReadSeeker) (err error) {
	var tmpi16 int16
	var tmpi32 int32

	tmpi16, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("kingdom read failed: bad sentinel read", err)
	}

	if int32(tmpi16) != 12346 {
		return newReadError("kingdom read failed: bad sentinel value", nil)
	}

	k.UnkInt32, err = filepacking.ReadInt32(r)
	if err != nil {
		panic("Failed to read unkint32")
	}

	// Read 28 shorts
	r.Seek(28*2, 1)

	tmpi32, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("kingdom read failed: Failed to read short count", err)
	}

	// "Read" tmpi32 shorts
	r.Seek(int64(2*tmpi32), 1)

	// Read 29 shorts.
	r.Seek(int64(2*29), 1)

	// read string
	k.UnkString, err = filepacking.ReadFileString(r)
	if err != nil {
		return newReadError("kingdom read failed: failed to read kingdom string", err)
	}

	// read 9 bytes
	r.Seek(int64(9), 1)

	// read 81 shorts (wtf?)
	r.Seek(int64(81*2), 1)

	// read 1 short
	r.Seek(int64(2), 1)

	// read 200 shorts
	r.Seek(int64(200*2), 1)

	// read 200 shorts
	r.Seek(int64(200*2), 1)

	// read 200 shorts
	r.Seek(int64(200*2), 1)

	// read int32
	r.Seek(int64(4), 1)

	return
}
