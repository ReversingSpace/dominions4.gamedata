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

// A Unit describes a Dominions unit.  For now this is largely unknown
// and undocumented, but is valuable for the purposes of storing and analysing.
//
// The data is sequential.
type Unit struct {

	// i64/u64 array
	u64unk00 uint64
	u64unk01 uint64
	u64unk02 uint64

	// i32/u32 array
	u32unk00 uint32

	// i16/u16 array
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

	// next (short)
	u16unk14 uint16

	// byte - unk
	b08unk00 byte

	// i32/u32
	u32unk01 uint32

	// byte - unk
	b08unk01 byte
}

// Read extracts the Unit from a reader stream.
func (u *Unit) Read(r io.ReadSeeker) (err error) {

	u.u64unk00, err = filepacking.ReadUInt64(r)
	if err != nil {
		return newReadError("unable to read unit: u64unk00", err)
	}

	u.u64unk01, err = filepacking.ReadUInt64(r)
	if err != nil {
		return newReadError("unable to read unit: u64unk01", err)
	}

	u.u64unk02, err = filepacking.ReadUInt64(r)
	if err != nil {
		return newReadError("unable to read unit: u64unk02", err)
	}

	u.u32unk00, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read unit: u32unk00", err)
	}

	u.u16unk00, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk00", err)
	}

	u.u16unk01, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk01", err)
	}

	u.u16unk02, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk02", err)
	}

	u.u16unk03, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk03", err)
	}

	u.u16unk04, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk04", err)
	}

	u.u16unk05, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk05", err)
	}

	u.u16unk06, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk06", err)
	}

	u.u16unk07, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk07", err)
	}

	u.u16unk08, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk08", err)
	}

	u.u16unk09, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk09", err)
	}

	u.u16unk10, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk10", err)
	}

	u.u16unk11, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk11", err)
	}
	u.u16unk12, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk12", err)
	}
	u.u16unk13, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk13", err)
	}

	u.u16unk14, err = filepacking.ReadUInt16(r)
	if err != nil {
		return newReadError("unable to read unit: u16unk14", err)
	}

	u.b08unk00, err = filepacking.ReadByte(r)
	if err != nil {
		return newReadError("unable to read unit: b08unk00", err)
	}

	u.u32unk01, err = filepacking.ReadUInt32(r)
	if err != nil {
		return newReadError("unable to read unit: u32unk01", err)
	}

	u.b08unk01, err = filepacking.ReadByte(r)
	if err != nil {
		return newReadError("unable to read unit: b08unk01", err)
	}

	return
}
