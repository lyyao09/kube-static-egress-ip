package iptables

import (
	"fmt"

	"github.com/coreos/go-iptables/iptables"
	"github.com/nirmata/kube-static-egress-ip/enforcer"
)

type manager struct {
	ipt *iptables.IPTables
}

// NewSourceIPModifier is a constructor for SourceIPModifier
func NewSourceIPModifier() (enforcer.SourceIPModifier, error) {
	ipt, err := iptables.New()
	if err != nil {
		return nil, fmt.Errorf("failed to locate iptables: %v", err)
	}
	return &manager{ipt: ipt}, nil
}

// AddNatRules add an iptables rule which SNATs all traffic for destinationIP
// with NatVipIP. That is, any traffic with destination == destinationIP
// leaving this node will have a source IP == NatVipIP
func (m *manager) AddNatRules(destinationIP, NatVipIP string) error {
	return nil
}

// ClearNatRules clears IPtables rules added by AddNatRules
func (m *manager) ClearNatRules(detinationIP, NatVipIP string) error {
	return nil
}

// ListNatRules lists Nat rules configured
func (m *manager) ListNatRules() ([]string, error) {
	return nil, nil
}
