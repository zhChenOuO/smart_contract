package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ERC20ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// const MulticallABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

type AggregateResponse struct {
	BlockNumber *big.Int
	ReturnData  [][32]byte
}

type Ttttt string

const (
	AVSD Ttttt = "ABC"
)

func main() {

	fmt.Println(Ttttt("ABC") == AVSD)

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/21e4ec6a577c4062a7a097278a55d9eb")
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

	erc20Abi, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		panic(err)
	}

	b, _ := os.ReadFile("./mult_call.abi")
	multAbi, err := abi.JSON(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	balanceOf, _ := erc20Abi.Pack("balanceOf", common.HexToAddress("0x00192Fb10dF37c9FB26829eb2CC623cd1BF599E8"))
	b, err = multAbi.Pack("aggregate", []Struct0{
		{
			Target:   common.HexToAddress("0x86b78c93Eea132BFD942E1f316bAB02A127dfC84"),
			CallData: balanceOf,
		},
		{
			Target:   common.HexToAddress("0x1dd34B35d96f5159567EF4caD2c69C0f87c3195c"),
			CallData: balanceOf,
		},
	})
	if err != nil {
		panic(err)
	}

	multcallContract := common.HexToAddress("0xeefba1e63905ef1d7acba5a8513c70307c1ce441")
	b, err = client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &multcallContract,
		Data: b,
	}, nil)
	if err != nil {
		fmt.Printf("%+v %+v\n", b, err)
		return
	}

	var out AggregateResponse

	err = multAbi.UnpackIntoInterface(&out, "aggregate", b)
	if err != nil {
		fmt.Printf("%+v %+v\n", b, err)
	}

	for _, i := range out.ReturnData {
		out, _ := erc20Abi.Unpack("balanceOf", i[:])
		fmt.Printf("%+v\n", out)
	}
}
