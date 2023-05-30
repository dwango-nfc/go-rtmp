//
// Copyright (c) 2018- yutopp (yutopp@gmail.com)
//
// Distributed under the Boost Software License, Version 1.0. (See accompanying
// file LICENSE_1_0.txt or copy at  https://www.boost.org/LICENSE_1_0.txt)
//

package handshake

import (
	"encoding/binary"
	"io"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}

func (e *Encoder) EncodeS1C1(v0 S0C0, v1 *S1C1) error {
	memory := [1537]byte{}
	view := memory[0:0]
	view = append(view, byte(v0))
	view = appendS1C1(view, v1)
	if _, err := e.w.Write(view[:]); err != nil {
		return err
	}
	return nil
}

func (e *Encoder) EncodeS2C2(h *S2C2) error {
	memory := [1536]byte{}
	view := memory[0:0]
	view = appendS2C2(view, h)
	if _, err := e.w.Write(view[:]); err != nil {
		return err
	}
	return nil
}

func appendS1C1(out []byte, h *S1C1) []byte {
	out = binary.BigEndian.AppendUint32(out, h.Time)
	out = append(out, h.Version[:]...)
	out = append(out, h.Random[:]...)
	return out
}

func appendS2C2(out []byte, h *S2C2) []byte {
	out = binary.BigEndian.AppendUint32(out, h.Time)
	out = binary.BigEndian.AppendUint32(out, h.Time2)
	out = append(out, h.Random[:]...)
	return out
}
