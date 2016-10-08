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

// A MercenaryData object represents merc data from Dominions 4.
type MercenaryData struct {
	Name   string
	sArr00 []int16 // 3 shorts
	sArr01 []int16 // 15 shorts
	sUnk00 int16   // unk short
	sUnk01 int16   // unk short
	sUnk02 int16   // unk short
	sUnk03 int16   // unk short
	sUnk04 int16   // unk short
}

// Read extracts MercenaryData from the stream.
func (m *MercenaryData) Read(r io.ReadSeeker) (err error) {

	m.Name, err = filepacking.ReadFileString(r)
	if err != nil {
		return newReadError("mercenary data error: failed to read name", err)
	}

	m.sArr00 = make([]int16, 3)
	for i := 0; i < 3; i++ {
		m.sArr00[i], err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("mercenary data error: failed to entry for s arr 00", err)
		}
	}

	m.sArr01 = make([]int16, 15)
	for i := 0; i < 15; i++ {
		m.sArr01[i], err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("mercenary data error: failed to entry for s arr 01", err)
		}
	}

	m.sUnk00, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("mercenary data error: failed to entry for unk short 00", err)
	}

	m.sUnk01, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("mercenary data error: failed to entry for unk short 01", err)
	}

	m.sUnk02, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("mercenary data error: failed to entry for unk short 02", err)
	}

	m.sUnk03, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("mercenary data error: failed to entry for unk short 03", err)
	}

	m.sUnk04, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("mercenary data error: failed to entry for unk short 04", err)
	}

	var sentinel int16
	sentinel, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("mercenary data error: failed to get sentinel", err)
	}
	if sentinel != 26812 {
		return fmt.Errorf("mercenary data error: sentinel is bad (%d not 26812)", sentinel)
	}
	return
}
