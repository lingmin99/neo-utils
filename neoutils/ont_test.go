package neoutils_test

import (
	"log"
	"testing"

	"github.com/o3labs/neo-utils/neoutils"
)

func TestONTTransfer(t *testing.T) {
	endpoint := "http://dappnode2.ont.io:20336"
	wif := ""
	asset := "ont"
	to := "AcydXy1MvrzaT8qD3Qe4B8mqEoinTvRy8U"
	amount := float64(2)
	gasPrice := uint(500)
	gasLimit := uint(20000)
	txid, err := neoutils.OntologyTransfer(endpoint, gasPrice, gasLimit, wif, asset, to, amount)
	if err != nil {
		log.Printf("err %v", err)
		return
	}
	log.Printf("tx id =%v", txid)
}
