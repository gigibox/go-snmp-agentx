// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx

import "go-snmp-agentx/agentx/pdu"

// ListItem defines an item of the list handler.
type ListItem struct {
	Type  pdu.VariableType
	Value interface{}
}
