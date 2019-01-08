package main

import (
	"fmt"
	//	"reflect"

	"github.com/coreos/go-iptables/iptables"
)

func contains(list []string, value string) bool {
	for _, val := range list {
		if val == value {
			return true
		}
	}
	return false
}

func main() {
	ipt, err := iptables.New()
	chain := "OUTPUT"
	// Saving the list of chains before executing tests
	// chain now exists
	err = ipt.ClearChain("filter", chain)
	if err != nil {
		fmt.Printf("ClearChain (of empty) failed: %v", err)
	}
	listChain, err := ipt.ListChains("filter")
	if err != nil {
		fmt.Printf("ListChains of Initial failed: %v", err)
	}
	fmt.Println(listChain)
	// check that chain is fully gone and that state similar to initial one
	listChain, err = ipt.ListChains("filter")
	if err != nil {
		fmt.Printf("ListChains failed: %v", err)
	}
	fmt.Println(listChain)
	//if !reflect.DeepEqual(originaListChain, listChain) {
	//	fmt.Printf("ListChains mismatch: \ngot  %#v \nneed %#v\n", originaListChain, listChain)
	//}

	// put a simple rule in
	err = ipt.AppendUnique("filter", chain, "-s", "0/0", "-j", "ACCEPT")
	if err != nil {
		fmt.Printf("Append failed: %v", err)
	}

	listChain, err = ipt.ListWithCounters("filter", chain)
	if err != nil {
		fmt.Printf("ListChains failed: %v", err)
	}
	fmt.Println(listChain)

	stats, err := ipt.Stats("filter", chain)
	if err != nil {
		fmt.Printf("stats failed: %v", err)
	}
	fmt.Println(stats)

}
