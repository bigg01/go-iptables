package iptablenforcer

import (
"testing"
)

func TestFirewallRules(t *testing.T) {
	fwRules := firewallRule{chain: "GUO_OPENSHIFT_INPUT", rule: []string{"-p", "tcp", "--dport", "22", "-m", "conntrack", "--ctstate", "NEW,ESTABLISHED", "-j", "ACCEPT", "-m", "comment", "--comment", "\"tkggo test\""}}
	//fwRules1 := firewallRule{chain: "GUO_OPENSHIFT_INPUT", rule: []string{"-p", "tcp", "--dport", "23", "-m", "conntrack", "--ctstate", "NEW,ESTABLISHED", "-j", "ACCEPT", "-m", "comment", "--comment", "\"tkggo test 23\""}}
	//fwRuelSet := firewallRules{[]firewallRule{fwRules, fwRules1}}


	fwRules_want := "GUO_OPENSHIFT_INPUT"

	if fwRules.chain != fwRules_want {
		t.Errorf("got %v want %v", fwRules, fwRules_want)
	}
}