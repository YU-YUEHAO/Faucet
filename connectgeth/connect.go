package connectgeth

import (

	faucet "YU/Smart"
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"io/ioutil"
	"math/big"
)

var client *ethclient.Client

//连接geth
func init() {

	rpcDel, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}
client = ethclient.NewClient(rpcDel)
	//fmt.Println(client)
}

func Getaccout() (*ecdsa.PrivateKey, common.Address) {
	const file = "E:/go/go project/src/golastwork/privateChain/node1/nodedata/keystore/UTC--2021-09-12T17-06-06.881126000Z--00dc6e8b60fa02a5d83e525bbef3240e8ea54dc5"
	account, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	password := "1111"
	key, err := keystore.DecryptKey(account, password)
	if err != nil {
		panic(err)
	}
	//fmt.Println(key.PrivateKey, key.Address)
	return key.PrivateKey, key.Address

}
func Getsmartcontract() *faucet.Faucet{
	ins, err := faucet.NewFaucet(common.HexToAddress("0xCb90c08f7b3F7a21D8b2e6F3f3faDC5CD56c3EA5"), client)
	if err != nil {
		panic(err)
	}
	return ins
}


//进行转账交易
func Dotransfer(ins *faucet.Faucet,address common.Address,ops *bind.TransactOpts) (*types.Transaction,error){
   transfermessage,err :=ins.Dotransfer(ops,address)
	if err!=nil {
		panic(err)
	}
	return transfermessage,err
}
//获取gasprice
func GetgasPrice() (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return big.NewInt(0), err
	} else {
		return gasPrice, nil
	}

}
//获取nonce
func Getnonce(address common.Address) (uint64, error) {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, err
	} else {
		return nonce, nil
	}

}

//获取用户当天得到的货币数量
func Gettodaytoken(ins *faucet.Faucet,address common.Address,header *types.Header) *big.Int {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	todaytoken ,err:=  ins.Gettodaytoken(&opts,address)
	if err!=nil {
		panic(err)
	}
	return todaytoken
}

//获取用户账户得币总数量
func Gettransfercount(ins *faucet.Faucet,address common.Address,header *types.Header) *big.Int {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	   transfercount,err:= ins.Gettransfercount(&opts,address)
	if err != nil{
		panic(err)
	}
	return transfercount
}

//获取账户最新交易时间
func Gettransfertime(ins *faucet.Faucet,address common.Address,header *types.Header) *big.Int {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
    transfertime, err := ins.Gettransfertime(&opts,address)
	if err !=nil{
		panic(err)
	}
	return transfertime
}

//获取水龙头合约余额
func GetBalancecontract (ins *faucet.Faucet,address common.Address,header *types.Header) (*big.Int){
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	Balancecontract, err :=ins.GetBalancecontract(&opts)
	if err!=nil {
		 panic(err)
	}
	return Balancecontract
}



//获取本用户余额
func GetBalanceUser(ins *faucet.Faucet, address common.Address, header *types.Header) (*big.Int) {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	balance, err := ins.GetBalanceUser(&opts)
	if err != nil {
		panic(err)
	}
	return balance
	//fmt.Println(balance)
}
//获取区块数
func GetBlockNumber() (*types.Header,error) {
    header,err :=client.HeaderByNumber(context.Background(),nil)
	if err!=nil {
		panic(err)
	}
	//fmt.Println(header)
	return header, err
}
//设置TransactOpts
func setopts(privateKey *ecdsa.PrivateKey, address common.Address) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(10001))
	if err != nil {
		panic(err)
	}
	nonce, err := Getnonce(address)
	if err != nil {
		panic(err)
	}
	gasPrice, err := GetgasPrice()
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth
}
func GetTxopts() *bind.TransactOpts{
	privateKey, publicKey := Getaccout()
	opts := setopts(privateKey, publicKey)
	return opts
}
