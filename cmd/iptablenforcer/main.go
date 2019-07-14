package main

import (
	"fmt"
	//	"reflect"
	"os"
	//"encoding/json"
	//"strings"
	"github.com/bigg01/go-iptables/pkg/iptablenforcer"
	log "github.com/sirupsen/logrus"
)


func main() {

iptablenforcer.ApplRules()
}