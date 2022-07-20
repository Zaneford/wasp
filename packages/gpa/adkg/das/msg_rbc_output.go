// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package das

import (
	"github.com/iotaledger/wasp/packages/gpa"
	"golang.org/x/xerrors"
)

//
// An event to self.
type msgRBCOutput struct { // TODO: Not used.
	me      gpa.NodeID
	indexes []int
}

var _ gpa.Message = &msgRBCOutput{}

func (m *msgRBCOutput) Recipient() gpa.NodeID {
	return m.me
}

func (m *msgRBCOutput) SetSender(sender gpa.NodeID) {
	// Don't care the sender.
}

func (m *msgRBCOutput) MarshalBinary() ([]byte, error) {
	return nil, xerrors.Errorf("msgRBCOutput::MarshalBinary not implemented") // TODO: Implement.
}
