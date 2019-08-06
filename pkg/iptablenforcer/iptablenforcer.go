package iptablenforcer

import (
	"fmt"
	//	"reflect"
	"os"
	//"encoding/json"
	//"strings"
	"github.com/coreos/go-iptables/iptables"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})
	//&log.TextFormatter{
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func contains(list []string, value string) bool {
	for _, val := range list {
		if val == value {
			return true
		}
	}
	return false
}

/*
OS_FIREWALL_ALLOW --> Input for NodeExporter OCP ETC
OPENSHIFT-ADMIN-OUTPUT-RULES --> OUTPUT from POD
*/

func ApplRules() {
	ipt, err := iptables.New()

	chain := "GUO_OPENSHIFT_INPUT"
	// Saving the list of chains before executing tests
	// chain now exists	
	err = ipt.ClearChain("filter", chain)
	if err != nil {
		log.Warnf("ClearChain (of empty) failed: %v", err)
	} else {
		log.Infof("ClearChain done: %v", chain)
	}

	listChain, err := ipt.ListChains("filter")
	if err != nil {
		fmt.Printf("ListChains of Initial failed: %v", err)
	} else {
		log.Infof("List Chain works: %v", chain)
	}

	// check that chain is fully gone and that state similar to initial one
	listChain, err = ipt.ListChains("filter")
	if err != nil {
		fmt.Printf("ListChains failed: %v", err)
	} else {
		log.Infof("List Chain works: %v", listChain)
	}

	fwRules := firewallRule{chain: "GUO_OPENSHIFT_INPUT", rule: []string{"-p", "tcp", "--dport", "22", "-m", "conntrack", "--ctstate", "NEW,ESTABLISHED", "-j", "ACCEPT", "-m", "comment", "--comment", "\"tkggo test\""}}
	fwRules1 := firewallRule{chain: "GUO_OPENSHIFT_INPUT", rule: []string{"-p", "tcp", "--dport", "23", "-m", "conntrack", "--ctstate", "NEW,ESTABLISHED", "-j", "ACCEPT", "-m", "comment", "--comment", "\"tkggo test 23\""}}
	fwRuelSet := firewallRules{[]firewallRule{fwRules, fwRules1}}

	for _, rr := range fwRuelSet.rules {
		log.Infoln(rr.chain, rr.rule)
		err = ipt.AppendUnique("filter", rr.chain, rr.rule...)
		if err != nil {
			log.Warnf("Append failed: %v", err)
		} else {
			log.Infof("LisAppend works: %v, %v", fwRules.chain, fwRules.rule)
		}
	}

	stats, err := ipt.Stats("filter", chain)
	if err != nil {
		fmt.Printf("stats failed: %v", err)
	}
	fmt.Println("================>>")
	fmt.Println(stats)

}
