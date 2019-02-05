package neoutils

import (
	"fmt"
	"log"

	"github.com/o3labs/ont-mobile/ontmobile"
	"github.com/o3labs/ont-mobile/ontmobile/ontrpc"
)

func OntologyTransfer(endpoint string, gasPrice int, gasLimit int, wif string, asset string, to string, amount float64) (string, error) {
	raw, err := ontmobile.Transfer(uint(gasPrice), uint(gasLimit), wif, asset, to, amount)
	if err != nil {
		return "", err
	}
	log.Printf("raw = %x", raw.Data)
	txid, err := ontmobile.SendRawTransaction(endpoint, fmt.Sprintf("%x", raw.Data))
	if err != nil {
		return "", err
	}

	return txid, nil
}

func ClaimONG(endpoint string, gasPrice int, gasLimit int, wif string) (string, error) {
	raw, err := ontmobile.WithdrawONG(uint(gasPrice), uint(gasLimit), endpoint, wif)
	if err != nil {
		return "", err
	}

	txid, err := ontmobile.SendRawTransaction(endpoint, fmt.Sprintf("%x", raw.Data))
	if err != nil {
		return "", err
	}

	return txid, nil
}

func BuildOntologyInvocationTransaction(contract string, method string, args string, gasPrice int, gasLimit int, wif string, payer string) (string, error) {
	raw, err := ontmobile.BuildInvocationTransaction(contract, method, args, uint(gasPrice), uint(gasLimit), wif, payer)
	if err != nil {
		return "", err
	}

	return raw, nil
}

// OntologyInvoke : Invoke a neovm contract in Ontology
func OntologyInvoke(endpoint string, contract string, method string, args string, gasPrice int, gasLimit int, wif string, payer string) (string, error) {
	raw, err := ontmobile.BuildInvocationTransaction(contract, method, args, uint(gasPrice), uint(gasLimit), wif, payer)
	if err != nil {
		return "", err
	}

	txid, err := ontmobile.SendRawTransaction(endpoint, raw)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func OntologyGetBlockCount(endpoint string) (ontrpc.GetBlockCountResponse, error) {
	client := ontrpc.NewRPCClient(endpoint)
	response, err := client.GetBlockCount()
	if err != nil {
		return response, err
	}
	return response, nil
}

func OntologyGetBalance(endpoint string, address string) (ontrpc.GetBalanceResponse, error) {
	client := ontrpc.NewRPCClient(endpoint)
	response, err := client.GetBalance(address)
	if err != nil {
		return response, err
	}
	return response, nil
}

func OntologyGetSmartCodeEvent(endpoint string, txHash string) (ontrpc.GetSmartCodeEventResponse, error) {
	client := ontrpc.NewRPCClient(endpoint)
	response, err := client.GetSmartCodeEvent(txHash)
	if err != nil {
		return response, err
	}
	return response, nil
}

func OntologySendRawTransaction(endpoint string, raw string) (string, error) {
	txid, err := ontmobile.SendRawTransaction(endpoint, raw)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func OntologyGetStorage(endpoint string, scriptHash string, key string) (ontrpc.GetStorageResponse, error) {
	client := ontrpc.NewRPCClient(endpoint)
	response, err := client.GetStorage(scriptHash, key)
	if err != nil {
		return response, err
	}
	return response, nil
}

func OntologyGetRawTransaction(endpoint string, txID string) (ontrpc.GetRawTransactionResponse, error) {
	client := ontrpc.NewRPCClient(endpoint)
	response, err := client.GetRawTransaction(txID)
	if err != nil {
		return response, err
	}
	return response, nil
}

func OntologyGetBlockWithHash(endpoint string, blockHash string) (ontrpc.GetBlockResponse, error) {
	client := ontrpc.NewRPCClient(endpoint)
	response, err := client.GetBlockWithHash(blockHash)
	if err != nil {
		return response, err
	}
	return response, nil
}

func OntologyGetBlockWithHeight(endpoint string, blockHeight int) (ontrpc.GetBlockResponse, error) {
	client := ontrpc.NewRPCClient(endpoint)
	response, err := client.GetBlockWithHeight(blockHeight)
	if err != nil {
		return response, err
	}
	return response, nil
}
