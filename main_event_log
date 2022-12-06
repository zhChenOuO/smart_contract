package main

import (
	"context"
	"fmt"
	"log"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// const OPABI = `[{"inputs":[{"internalType":"address","name":"_oneInchRouter","type":"address"},{"internalType":"address","name":"_imToken","type":"address"},{"internalType":"address","name":"_fundsProvider","type":"address"},{"internalType":"address payable","name":"_swapFeeTo","type":"address"},{"internalType":"address payable","name":"_gasFeeTo","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bool","name":"_prev","type":"bool"},{"indexed":false,"internalType":"bool","name":"_curr","type":"bool"}],"name":"FlipRunning","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"_prev","type":"address"},{"indexed":false,"internalType":"address","name":"_curr","type":"address"}],"name":"FundsProvider","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"_prev","type":"address"},{"indexed":false,"internalType":"address","name":"_curr","type":"address"}],"name":"GasFeeTo","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"token","type":"address"},{"indexed":false,"internalType":"uint256","name":"amt","type":"uint256"},{"indexed":false,"internalType":"address","name":"to","type":"address"}],"name":"Pull","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"token","type":"address"},{"indexed":false,"internalType":"uint256","name":"amt","type":"uint256"}],"name":"Push","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"_addr","type":"address"},{"indexed":false,"internalType":"bool","name":"_isWhitelist","type":"bool"}],"name":"SetWhitelist","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes","name":"id","type":"bytes"},{"indexed":false,"internalType":"bytes","name":"uniqueId","type":"bytes"},{"indexed":false,"internalType":"enum Operator.ACTION","name":"action","type":"uint8"},{"indexed":false,"internalType":"address","name":"srcToken","type":"address"},{"indexed":false,"internalType":"address","name":"dstToken","type":"address"},{"indexed":false,"internalType":"address","name":"tokenFrom","type":"address"},{"indexed":false,"internalType":"address","name":"tokenTo","type":"address"},{"indexed":false,"internalType":"uint256","name":"retAmt","type":"uint256"}],"name":"Swap","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"_prev","type":"address"},{"indexed":false,"internalType":"address","name":"_curr","type":"address"}],"name":"SwapFeeTo","type":"event"},{"inputs":[{"internalType":"bytes","name":"_id","type":"bytes"},{"internalType":"bytes","name":"_uniqueId","type":"bytes"},{"internalType":"uint256","name":"_gasFeeAmt","type":"uint256"},{"internalType":"bytes","name":"_data","type":"bytes"}],"name":"crossSwap","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes","name":"_id","type":"bytes"},{"internalType":"bytes","name":"_uniqueId","type":"bytes"},{"internalType":"uint256","name":"_swapFeeAmt","type":"uint256"},{"internalType":"bytes","name":"_data","type":"bytes"}],"name":"doSwap","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"flipRunning","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes","name":"_id","type":"bytes"},{"internalType":"bytes","name":"_uniqueId","type":"bytes"},{"internalType":"uint256","name":"_amt","type":"uint256"},{"internalType":"uint256","name":"_swapFeeAmt","type":"uint256"}],"name":"fromUCross","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"gasFeeTo","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address[]","name":"_tokens","type":"address[]"}],"name":"getBalance","outputs":[{"internalType":"uint256[]","name":"balances","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getFundsProvider","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"imToken","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"isRunning","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"oneInchRouter","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_token","type":"address"},{"internalType":"uint256","name":"_amt","type":"uint256"},{"internalType":"address","name":"_to","type":"address"}],"name":"pull","outputs":[{"internalType":"uint256","name":"amt","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_token","type":"address"},{"internalType":"uint256","name":"_amt","type":"uint256"}],"name":"push","outputs":[{"internalType":"uint256","name":"amt","type":"uint256"}],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_newFundsProvider","type":"address"}],"name":"setFundsProvider","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_newGasFeeTo","type":"address"}],"name":"setGasFeeTo","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_newSwapFeeTo","type":"address"}],"name":"setSwapFeeTo","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address[]","name":"_addrArr","type":"address[]"},{"internalType":"bool[]","name":"_flags","type":"bool[]"}],"name":"setWhitelist","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"swapFeeTo","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes","name":"_id","type":"bytes"},{"internalType":"bytes","name":"_uniqueId","type":"bytes"},{"internalType":"uint256","name":"_amt","type":"uint256"},{"internalType":"uint256","name":"_gasFeeAmt","type":"uint256"},{"internalType":"address","name":"_to","type":"address"}],"name":"toUCross","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"useless","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"whitelist","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"stateMutability":"payable","type":"receive"}]`

// 0x7c02520000000000000000000000000005ad60d9a2f1aa30ba0cdbaf1e0a0a145fbea16f00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000180000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee000000000000000000000000e9e7cea3dedca5984780bafc599bd69add087d5600000000000000000000000005ad60d9a2f1aa30ba0cdbaf1e0a0a145fbea16f000000000000000000000000b983a7fd629d943582eff49bb3883cd52a8c519b000000000000000000000000000000000000000000000000000009184e72a000000000000000000000000000000000000000000000000000000aa1dea5e101da000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000038000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c000000000000000000000000000000000000000000000000000009184e72a00000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000004d0e30db0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000064d1660f99000000000000000000000000bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c00000000000000000000000051e6d27fa57373d8d4c256231241053a70cb1d93000000000000000000000000000000000000000000000000000009184e72a0000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a4b757fed600000000000000000000000051e6d27fa57373d8d4c256231241053a70cb1d93000000000000000000000000bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c000000000000000000000000e9e7cea3dedca5984780bafc599bd69add087d560000000000000000001e84801111111254fb6c44bac0bed2854e76f90643097d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e7af3465
func main() {
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

	txHash := common.HexToHash("0x19c8f8bff8f2b62b6f452ba1e86983aca6a147ebb9d3970ab5381a58ead292c9")

	// contractAbi, err := OpMetaData.GetAbi()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	op, err := NewOpFilterer(common.HexToAddress("0x3AC2b846711897f1c287A6489011dC2c5Ef5C33c"), client)
	if err != nil {
		log.Fatal(err)
	}
	// oneinch, err := NewOneinchFilterer(common.HexToAddress("0x3AC2b846711897f1c287A6489011dC2c5Ef5C33c"), client)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	// eventSignature := []byte("Swap(string, string, address, uint8, address, address, address, address, uint256, uint256, uint256, uint256)")
	// hash := crypto.Keccak256Hash(eventSignature)

	for _, txLog := range receipt.Logs {
		// if val, err := oneinch.ParseSwapped(*txLog); err == nil {
		// 	fmt.Printf("ParseSwapped %+v\n", val)
		// 	continue
		// }

		val, err := op.ParseSwap(*txLog)
		if err != nil {
			fmt.Println(txLog.Topics[0])
			continue
		}
		fmt.Println("val", val.FeeAmt)
		fmt.Println("swap", txLog.Topics[0])

		// v := map[string]interface{}{}

		// fmt.Println("?", common.Bytes2Hex(txLog.Data))
		// if err := contractAbi.UnpackIntoMap(v, "Swap", txLog.Data); err != nil {
		// 	continue
		// }
		// fmt.Printf("ParseSwap %+v\n", v)
		// vv := v["id"].([]byte)

		// b, err := hex.DecodeString(string(vv))
		// if err != nil {
		// 	fmt.Errorf("???", err.Error())
		// }
		// fmt.Println("???", string(b))

		// reqID := hex.EncodeToString([]byte("1"))

		// bb, berr := hex.DecodeString(reqID)
		// if berr != nil {
		// 	fmt.Errorf("???", berr.Error())
		// }
		// fmt.Println("???", string(bb))

		// var topics [4]string
		// for i := range txLog.Topics {
		// 	topics[i] = txLog.Topics[i].Hex()
		// }

		// fmt.Println(topics)

		// fmt.Println("?", contractAbi.Events[topics[1]])
		// val := map[string]interface{}{}
		// if err := contractAbi.UnpackIntoMap(val, "doSwap", txLog.Data); err != nil {
		// 	continue
		// }

	}

	// blockHash := common.HexToHash("0x7f43e1f005d0afee68ae47d07007374c57433746fb035d6dd3dee817eb4bf001")
	// block := big.NewInt(22106437)

	// logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
	// 	BlockHash: &blockHash,
	// 	FromBlock: block,
	// 	ToBlock:   block,
	// 	Addresses: []common.Address{
	// 		common.HexToAddress("0x3AC2b846711897f1c287A6489011dC2c5Ef5C33c"),
	// 	},
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	// }

	// for i := range logs {
	// 	val, err := contractAbi.Unpack("ItemSet", vLog.Data)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
}
