// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx

import (
	"go-snmp-agentx/agentx/pdu"
	"go-snmp-agentx/agentx/value"
)

// Handler defines an interface for a handler of events that
// might occure during a session.
type Handler interface {
	Get(value.OID) (value.OID, pdu.VariableType, interface{}, error)
	GetNext(value.OID, bool, value.OID) (value.OID, pdu.VariableType, interface{}, error)
}
