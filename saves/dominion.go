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
	"fmt"
	"github.com/ReversingSpace/dominions4.gamedata/filepacking"
	"io"
)

// A Dominion object represents Dominion related information.
type Dominion struct {

	// 6 bytes
	b08l06 []byte

	// string
	Name string

	//
	u32unk00 uint32
	u32unk01 uint32

	// probably relationship data or ownership.
	unkMappedI32I32 map[int32]int32
}

// Read extracts Dominion information from the stream.
func (d *Dominion) Read(r io.ReadSeeker) (err error) {
	// Short (sig?): 12346
	var Magic uint16

	Magic, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("dominion: unable to read magic", err)
	}

	if Magic != 12346 {
		return fmt.Errorf("dominion: magic is not 12346 (is %d)", Magic)
	}

	d.b08l06 = make([]byte, 6)
	_, err = r.Read(d.b08l06)
	if err != nil {
		return newReadError("dominion: unable to read initial byte array", err)
	}

	d.Name, err = filepacking.ReadFileString(r)
	if err != nil {
		return newReadError("dominion: unable to read name", err)
	}

	d.u32unk00, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("dominion: unable to read u32unk00", err)
	}

	d.u32unk01, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("dominion: unable to read u32unk01", err)
	}

	// Create only if missing
	if d.unkMappedI32I32 == nil {
		d.unkMappedI32I32 = make(map[int32]int32)
	}

	var k int32
	var v int32
	i := 0
	for {
		k, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("dominion: unable to read k-v loop: k", err)
		}

		if k == 0 {
			break
		}

		v, err = filepacking.ReadInt32(r)
		if err != nil {
			return newReadError("dominion: unable to read k-v loop: v", err)
		}

		d.unkMappedI32I32[k] = v

		i++

		if i == 64 {
			break
		}
	}

	return
}
