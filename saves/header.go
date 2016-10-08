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
	"bytes"
	"github.com/ReversingSpace/dominions4.gamedata/filepacking"
	"io"
)

// FileHeader represents the common file header.
type FileHeader struct {
	// Signature (header)
	// 01 02 04 44 4F 4D
	Signature []byte

	// Useful for tracking users.
	// (Simplified explanation.)
	UserID int32

	// Game Version
	GameVersion int32

	// Game turn number
	// Can be -1
	TurnNumber int32

	// -1 for fatherland files?
	unk00 int32

	unk01 int32

	// RealmID (ignore it for fatherland files)
	RealmID int32

	// depends on realm; 0 if realm is < 0
	unk02 int32

	unk03 int32

	// Game Name
	GameName string

	// Password
	Password string

	// Master Password
	MasterPassword string

	// Useless for most purposes; if you don't know what
	// it does don't worry about it.
	TurnKey int32

	// Default sentinel; verified.
	sentinel int32
}

var fileHeaderSignature = []byte{0x01, 0x02, 0x04, 0x44, 0x4F, 0x4D}

// Read loads the FileHeader from the reader;
// if the error is non-nil it failed.
func (f *FileHeader) Read(r io.ReadSeeker) (err error) {
	f.Signature = make([]byte, 6)
	_, err = r.Read([]byte(f.Signature))
	if err != nil {
		err = newReadError("failed to read header: signature", err)
		return
	}
	if bytes.Compare(f.Signature, fileHeaderSignature) != 0 {
		return ErrorBadSignature
	}

	f.UserID, err = filepacking.ReadInt32(r)
	if err != nil {
		err = newReadError("failed to read header: userid", err)
		return
	}

	f.GameVersion, err = filepacking.ReadInt32(r)
	if err != nil {
		err = newReadError("failed to read header: game version", err)
		return
	}

	// pretty sure dom3 stops well before this, but we don't want anything
	// below ~410 anyway.
	if f.GameVersion < 410 {
		err = newReadError("failed to read header: game version below 400", nil)
		return
	}

	f.TurnNumber, err = filepacking.ReadInt32(r)
	if err != nil {
		err = newReadError("failed to read header: turn number", err)
		return
	}

	f.unk00, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("failed to read header: unk00", err)
	}

	f.unk01, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("failed to read header: unk01", err)
	}

	f.RealmID, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("failed to read header: realm", err)
	}

	f.unk02, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("failed to read header: unk02", err)
	}

	f.unk03, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("failed to read header: unk03", err)
	}

	f.GameName, err = filepacking.ReadFileString(r)
	if err != nil {
		return newReadError("failed to read header: name", err)
	}

	f.Password, err = filepacking.ReadFileStringRX(r)
	if err != nil {
		return newReadError("failed to read header: password", err)
	}

	f.MasterPassword, err = filepacking.ReadFileStringRX(r)
	if err != nil {
		return newReadError("failed to read header: master password", err)
	}

	f.TurnKey, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("failed to read header: turn key", err)
	}

	f.sentinel, err = filepacking.ReadInt32(r)
	if err != nil {
		return newReadError("failed to read header: sentinel", err)
	}

	if f.sentinel != int32(12346) {
		return newReadError("failed to read header: bad sentinel value", nil)
	}
	return
}
