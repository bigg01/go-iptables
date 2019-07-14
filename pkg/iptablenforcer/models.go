package iptablenforcer

type firewallRule struct {
	chain string
	rule []string
}

type firewallRules struct {
	rules []firewallRule
}
