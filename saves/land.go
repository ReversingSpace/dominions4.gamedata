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

// Land describes the known land information.
type Land struct {
	// Used in serialisation and loading
	TreatAsFatherland bool

	// Name of the land
	Name string

	// Unknown int16
	unkShort0 int16 // Unknown

	// Sites
	Sites [8]LandInfoSite

	// Scoring value?
	DominionStrength int16

	// Unrest
	Unrest int16

	// Something related to recruitment
	unkShortRecruit1 int16
	unkShortRecruit2 int16

	// Completely unknown
	unkShort1 int16

	// Pretty sure this is the basic resource value
	BasicResources int16

	// Realm or Dominions owner? - changes to besieger.
	OwnerField1 int16

	// Owner of current Commander? - 'owner'
	OwnerField2 int16

	// Owner adjacent? (sometimes matches 1/2) - not sure
	OwnerField3 int16

	// Seems to be lab related
	Lab0 bool
	Lab1 bool

	// Seems to be temple related 1 = temple; 0 = none
	Temple bool

	// Seems to be fort related
	Fort int8

	// Base def value
	Defence int8

	// Unknown
	unkByte0 int8
	unkByte1 int8
	unkByte2 int8
	unkByte3 int8
	unkByte4 int8
	unkByte5 int8

	// Population Type
	PopulationType int8

	// Site reveal status has been moved into Sites

	// LandInformation - Relationship to other lands?
	LandInformation [0xC7][9]byte

	// Unknown
	unkShort2 int16
	unkShort3 int16
	unkShort4 int16
	unkShort5 int16

	// Array
	unkRecruitArray1 []int16

	// Array
	unkRecruitArray2 []int16

	// Population
	Population int16

	// DefenceValue
	DefenceValue int16

	// VictoryPoints(?) used by 2h and others
	VictoryPoints int16

	Corpses1 int16
	Corpses2 int16

	FortParts [8]int16

	dynamicArray []byte

	EventCode int16

	unkString string

	unkResultInt int32

	Geometry LandInfoGeom
}

// LandInfoSite holds site information
type LandInfoSite struct {
	Revealed bool
	SiteID   int16
}

// LandInfoGeom holds geometry information
type LandInfoGeom struct {
	X          int16     // x
	Y          int16     // y
	Neighbours [20]int16 // 0 = nothing; otherwise GeomID

	unkShorts0 [9]int16 // unk

	unkShorts1 [20]int16 // unk

	// #terrain flags
	// 0x4 = sea;
	// 0x80 = forest;
	// 0x400000 = border mountains
	Terrain int32

	unk32 int32

	unkStr1 string
	unkStr2 string
}

// Read extracts the land instance from stream.
// Make sure to set the `TreatAsFatherland` variable.
func (l *Land) Read(r io.ReadSeeker) (err error) {
	var b byte
	var b2 byte

	if l.TreatAsFatherland {
		l.Name, err = filepacking.ReadFileStringN(r, 36)
		if err != nil {
			return newReadError("land read failed: name", err)
		}

		l.unkShort0, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read failed: unk short 0", err)
		}

		for i := 0; i < 8; i++ {
			l.Sites[i].SiteID, err = filepacking.ReadInt16(r)
			if err != nil {
				return newReadError("land read failed: site id array", err)
			}
		}

		l.DominionStrength, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read failed: dominion strength", err)
		}

		l.Unrest, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read failed: unrest", err)
		}

		l.unkShortRecruit1, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read failed: unk short recruit 1", err)
		}

		l.unkShortRecruit2, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read failed: unk short recruit 2", err)
		}

		l.unkShort1, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read failed: unk short 1", err)
		}

		l.BasicResources, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read failed: basic resources", err)
		}

		l.OwnerField1, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read failed: owner field 1", err)
		}

		l.OwnerField2, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read failed: owner field 2", err)
		}

		l.OwnerField3, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read failed: owner field 3", err)
		}

		b, err = filepacking.ReadByte(r)
		if err != nil {
			return newReadError("land read failed: lab 0", err)
		}
		l.Lab0 = b == 1

		b, err = filepacking.ReadByte(r)
		if err != nil {
			return newReadError("land read failed: lab 1", err)
		}
		l.Lab1 = b == 1

		b, err = filepacking.ReadByte(r)
		if err != nil {
			return newReadError("land read failed: temple", err)
		}
		l.Temple = b == 1

		l.Fort, err = filepacking.ReadInt8(r)
		if err != nil {
			return newReadError("land read failed: fort", err)
		}
	} else {
		_, err = filepacking.ReadFileStringN(r, 0)
		if err != nil {
			panic("Failed to read land name (throwaway variant)")
		}
		r.Seek(40, 1)
	}

	l.Defence, err = filepacking.ReadInt8(r)
	if err != nil {
		return newReadError("land read failed: defence", err)
	}
	l.unkByte0, err = filepacking.ReadInt8(r)
	if err != nil {
		return newReadError("land read failed: unk byte 0", err)
	}
	l.unkByte1, err = filepacking.ReadInt8(r)
	if err != nil {
		return newReadError("land read failed: unk byte 1", err)
	}
	l.unkByte2, err = filepacking.ReadInt8(r)
	if err != nil {
		return newReadError("land read failed: unk byte 2", err)
	}
	l.unkByte3, err = filepacking.ReadInt8(r)
	if err != nil {
		return newReadError("land read failed: unk byte 3", err)
	}
	l.unkByte4, err = filepacking.ReadInt8(r)
	if err != nil {
		return newReadError("land read failed: unk byte 4", err)
	}
	l.unkByte5, err = filepacking.ReadInt8(r)
	if err != nil {
		return newReadError("land read failed: unk byte 5", err)
	}
	l.PopulationType, err = filepacking.ReadInt8(r)
	if err != nil {
		return newReadError("land read failed: population type", err)
	}

	for i := 0; i < 8; i++ {
		b, err = filepacking.ReadByte(r)
		if err != nil {
			return newReadError("land read failed: sites revealed", err)
		}
		l.Sites[i].Revealed = b == 1
	}

	/*
		if l.Defence > 100 {
			// bad value
		}
	*/

	for {
		b, err = filepacking.ReadByte(r)
		if err != nil {
			return newReadError("land read error: failed in loop - no land id byte", err)
		}
		if b == 0xFF {
			break
		}

		if b > 0xC7 {
			return newReadError("land read error: bad byte in loop; land id exceeds game maximum", nil)
		}

		b2, err = filepacking.ReadByte(r)
		if err != nil {
			return newReadError("land read error: no offset byte", err)
		}

		if b2 > 8 {
			return newReadError("land read error: bad byte in loop", nil)
		}

		l.LandInformation[b][b2], err = filepacking.ReadByte(r)
		if err != nil {
			return newReadError("land read error: no information byte", err)
		}
	}

	if l.unkShortRecruit1 > 0 {
		l.unkRecruitArray1 = make([]int16, l.unkShortRecruit1)
	}
	for i := int16(0); i < l.unkShortRecruit1; i++ {
		l.unkRecruitArray1[i], err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read error: bad byte in unk recruit array 1", err)
		}
	}

	if l.unkShortRecruit2 > 0 {
		l.unkRecruitArray2 = make([]int16, l.unkShortRecruit2)
	}
	for i := int16(0); i < l.unkShortRecruit2; i++ {
		l.unkRecruitArray2[i], err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read error: bad byte in unk recruit array 2", err)
		}
	}

	l.unkShort2, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("land read error: unk short 2", err)
	}

	l.unkShort3, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("land read error: unk short 3", err)
	}

	l.unkShort4, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("land read error: unk short 4", err)
	}

	l.unkShort5, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("land read error: unk short 5", err)
	}

	if l.TreatAsFatherland {
		l.Population, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read error: population", err)
		}

		l.DefenceValue, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read error: defence value", err)
		}

		l.VictoryPoints, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read error: victory points", err)
		}

		l.Corpses1, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read error: corpses 1", err)
		}

		l.Corpses2, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read error: corpses 2", err)
		}
	} else {
		// 2h doesn't need this data (from what's known of it)
		r.Seek(4, 1)
		l.VictoryPoints, err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read error: victory points", err)
		}
		r.Seek(4, 1)
	}

	r.Seek(2, 1)

	l.Geometry.Terrain, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("land read error: geometry.terrain", err)
	}

	l.Geometry.X, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("land read error: geometry.x", err)
	}

	l.Geometry.Y, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("land read error: geometry.y", err)
	}

	for i := 0; i < 20; i++ {
		l.Geometry.Neighbours[i], err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read error: geometry.neighbours", err)
		}
	}

	// Skip 4 unknown ints
	r.Seek(32, 1)

	l.Geometry.unkStr1, err = filepacking.ReadFileString(r)
	if err != nil {
		return newReadError("land read error: geo unk string 1", err)
	}

	l.Geometry.unkStr2, err = filepacking.ReadFileString(r)
	if err != nil {
		return newReadError("land read error: geo unk string 1", err)
	}

	// Unk shorts
	for i := 0; i < 9; i++ {
		l.Geometry.unkShorts0[i], err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read error: geometry unk shorts 0", err)
		}
	}

	l.Geometry.unk32, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("land read error: geometry.unk32", err)
	}

	// UnkCodes{A-D} (2+2+1+1)
	// + UnkCodeE (1)
	// + UnkDWCodeA + UnkDWCodeB (4+4)
	// = Skip 15 bytes
	r.Seek(15, 1)

	// Fort Parts
	for i := 0; i < 8; i++ {
		l.FortParts[i], err = filepacking.ReadInt16(r)
		if err != nil {
			return newReadError("land read error: fort parts", err)
		}
	}

	// geo (20 shorts)
	for i := 0; i < 20; i++ {
		l.Geometry.unkShorts1[i], err = filepacking.ReadInt16(r)
		return newReadError("land read error: unk shorts 1", err)
	}

	// Dynamic array (it could be smaller or larger than 200)
	l.dynamicArray, err = filepacking.ReadSparse(r, 200)
	if err != nil {
		return newReadError("land read error: dynamic array", err)
	}

	l.EventCode, err = filepacking.ReadInt16(r)
	if err != nil {
		return newReadError("land read error: eventcode", err)
	}

	// No reason to support anything < 420, so the old checks are gone.

	// Max length is 2000
	l.unkString, err = filepacking.ReadFileStringN(r, 2000)
	if err != nil {
		return newReadError("land read error: unk string", err)
	}

	// This int could be related to the final result/state
	l.unkResultInt, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("land read error: unk result int", err)
	}

	return
}
