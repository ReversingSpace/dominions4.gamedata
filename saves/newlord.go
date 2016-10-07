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

// A Lord represents a Pretender (or Disciple)
type Lord struct {
	Header    FileHeader
	u32unk00  uint32
	u32unk01  uint32
	Unit      Unit
	Commander Commander
	Dominion  Dominion
	u32unk02  uint32
}

// Read extracts the lord from a newlord file.
func (l *Lord) Read(r io.ReadSeeker) (err error) {
	err = l.Header.Read(r)
	if err != nil {
		return
	}

	l.u32unk00, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read lord: u32unk00", err)
	}

	l.u32unk01, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read lord: u32unk01", err)
	}

	err = l.Unit.Read(r)
	if err != nil {
		return
	}

	err = l.Commander.Read(r)
	if err != nil {
		return
	}

	err = l.Dominion.Read(r)
	if err != nil {
		return
	}

	l.u32unk02, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read lord: u32unk02", err)
	}

	return
}
