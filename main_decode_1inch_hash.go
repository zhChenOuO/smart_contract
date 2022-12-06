package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

var calldata string = "0x7c02520000000000000000000000000018101ac1d35230f1a3c005e2abaaeb25cae79e7f00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000180000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee000000000000000000000000ba2ae424d960c26247dd6c32edc70b295c744c4300000000000000000000000018101ac1d35230f1a3c005e2abaaeb25cae79e7f000000000000000000000000b983a7fd629d943582eff49bb3883cd52a8c519b00000000000000000000000000000000000000000000000000230dae5a611a1a00000000000000000000000000000000000000000000000000000000ba94732700000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e70000000000000000000000000000000000000000000000000000a900001a4041bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095cd0e30db00c20bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095cac109c8025f272414fd9e2faa805a583708a017f6ae40711b8002625a0ac109c8025f272414fd9e2faa805a583708a017f1111111254fb6c44bac0bed2854e76f90643097d00000000000000000000000000000000000000000000000000000000b8bb88d3bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c00000000000000000000000000000000000000000000000000230dae5a611a1a00000000000000000000000000000000000000000000000000e7af3465"

type TmpStruct struct {
	AggregationRouterV4SwapDescription `json:"desc"`
}

func main() {
	// 0.000999179926735771
	// 0.999179926735770363
	//

	amount, _ := decimal.NewFromString("0.000999179926735770364")
	amountTT, _ := decimal.NewFromString("0.998180746809034593")
	fmt.Println(amount.Add(amountTT))
	client, err := ethclient.Dial("https://bsc-dataseed1.binance.org/")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	tt := time.Now()
	t, err := client.EstimateGas(context.Background(), ethereum.CallMsg{})
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Since(tt).Seconds())
	fmt.Println(t)

	decodedData, err := hex.DecodeString(calldata[10:])
	if err != nil {
		log.Fatal(err)
	}

	oneinchABI, _ := OneinchMetaData.GetAbi()

	out := map[string]interface{}{}
	if err := oneinchABI.Methods["swap"].Inputs.UnpackIntoMap(out, decodedData); err != nil {
		log.Fatalf("%v", err)
	}
	// fmt.Printf("%T", out["desc"])
	var tmmp TmpStruct

	b, _ := json.Marshal(out)
	fmt.Printf("%s\n", string(b))
	json.Unmarshal(b, &tmmp)
	fmt.Printf("%+v\n", tmmp)

	return

	// decode txInput method signature
	decodedSig, err := hex.DecodeString(calldata[2:10])
	if err != nil {
		log.Fatal(err)
	}
	// recover Method from signature and ABI
	method, err := oneinchABI.MethodById(decodedSig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(method)

	// if err := method.Inputs.UnpackIntoMap(out, decodedData); err != nil {
	// 	log.Fatalf("%v", err)
	// }

	// fmt.Printf("%+v", out["desc"])
}
