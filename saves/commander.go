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

// A Commander represents a commander unit
type Commander struct {
	Name string

	//
	u32unk00 uint32

	//
	u32unk01 uint32
	u32unk02 uint32
	u32unk03 uint32
	u32unk04 uint32
	u32unk05 uint32

	//
	u16unk00 uint16
	u16unk01 uint16
	u16unk02 uint16
	u16unk03 uint16
	u16unk04 uint16
	u16unk05 uint16
	u16unk06 uint16
	u16unk07 uint16
	u16unk08 uint16
	u16unk09 uint16
	u16unk10 uint16
	u16unk11 uint16
	u16unk12 uint16
	u16unk13 uint16
	u16unk14 uint16
	u16unk15 uint16

	//
	u32unk06 uint32
	u32unk07 uint32
	u32unk08 uint32
	u32unk09 uint32

	//
	u16unk16 uint16
	u16unk17 uint16
	u16unk18 uint16
	u16unk19 uint16
	u16unk20 uint16
	u16unk21 uint16
	u16unk22 uint16
	u16unk23 uint16
	u16unk24 uint16

	// 46 + 5
	unkByteArray []byte

	//
	u16unk25 uint16
}

// Read extracts the Commander object from the stream.
func (u *Commander) Read(r io.ReadSeeker) (err error) {

	u.Name, err = filepacking.ReadFileString(r)
	if err != nil {
		return newReadError("unable to read commander: name", err)
	}

	u.u32unk00, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read commander: u32unk00", err)
	}

	u.u32unk01, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read commander: u32unk01", err)
	}

	u.u32unk02, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read commander: u32unk02", err)
	}

	u.u32unk03, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read commander: u32unk03", err)
	}

	u.u32unk04, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read commander: u32unk04", err)
	}

	u.u32unk05, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read commander: u32unk05", err)
	}

	u.u16unk00, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk00", err)
	}

	u.u16unk01, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk01", err)
	}

	u.u16unk02, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk02", err)
	}

	u.u16unk03, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk03", err)
	}

	u.u16unk04, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk04", err)
	}

	u.u16unk05, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk05", err)
	}

	u.u16unk06, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk06", err)
	}

	u.u16unk07, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk07", err)
	}

	u.u16unk08, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk08", err)
	}

	u.u16unk09, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk09", err)
	}

	u.u16unk10, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk10", err)
	}

	u.u16unk11, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk11", err)
	}

	u.u16unk12, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk12", err)
	}

	u.u16unk13, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk13", err)
	}

	u.u16unk14, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk14", err)
	}

	u.u16unk15, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk15", err)
	}

	u.u32unk06, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read commander: u32unk06", err)
	}

	u.u32unk07, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read commander: u32unk07", err)
	}

	u.u32unk08, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read commander: u32unk08", err)
	}

	u.u32unk09, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read commander: u32unk09", err)
	}

	u.u16unk16, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk16", err)
	}

	u.u16unk17, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk17", err)
	}

	u.u16unk18, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk18", err)
	}

	u.u16unk19, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk19", err)
	}

	u.u16unk20, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk20", err)
	}

	u.u16unk21, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk21", err)
	}

	u.u16unk22, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk22", err)
	}

	u.u16unk23, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk23", err)
	}

	u.u16unk24, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk24", err)
	}

	u.unkByteArray = make([]byte, 51)
	_, err = r.Read(u.unkByteArray)
	if err != nil {
		return newReadError("unable to read commander: unkByteArray", err)
	}

	u.u16unk25, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read commander: u16unk25", err)
	}

	return
}
