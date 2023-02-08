package token

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type T struct{}

const (
	BSC_RPC_URL      = "https://endpoints.omniatech.io/v1/bsc/testnet/public"
	CONTRACT_ADDR    = "0xbF47aBE2d374313Eda18C77359E50dC3281A4D36"
	ACCOUNT_PUB_KEY  = "0x0eecD003Aa72554354527BeD08dC88f971d881DF"
	ACCOUNT_PRIV_KEY = "f4a7287be8fde0ea3a350956ab2c0e49c0c1e63be354d28c1e1cca8b3978af8e"
)

var tokenRpc TokenRPC
var err error

type TokenRPC struct {
	conn  *ethclient.Client
	token *Token
}

func init() {
	tokenRpc.conn, err = ethclient.Dial(BSC_RPC_URL)
	if err != nil {
		defer tokenRpc.conn.Close()
		log.Fatal("rpc error!: ", err)
	}

	tokenRpc.token, err = NewToken(common.HexToAddress(CONTRACT_ADDR), tokenRpc.conn)
	if err != nil {
		log.Fatal("NewToken error!: ", err)
	}

	// symbol, err := tokenRpc.token.Symbol(&bind.CallOpts{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("symbol:", symbol)
	// total_supply, err := tokenRpc.token.TotalSupply(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("total_supply:", total_supply)
}

func (t *T) Name() string {
	name, err := tokenRpc.token.TokenCaller.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Get Name Error: ", err)
	}
	return name
}

func (t *T) Symbol() string {
	symbol, err := tokenRpc.token.TokenCaller.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Get Symbol Error: ", err)
	}
	return symbol
}

func (t *T) Owner() string {
	ownerAddress, err := tokenRpc.token.TokenCaller.Owner(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Get Owner Error: ", err)
	}
	return ownerAddress.String()
}

func (t *T) TotalSupply() int64 {
	totalSupply, err := tokenRpc.token.TokenCaller.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Get TotalSupply Error: ", err)
	}
	return totalSupply.Int64()
}

func (t *T) GetOwner() string {
	ownerAddress, err := tokenRpc.token.TokenCaller.GetOwner(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Get GetOwner Error: ", err)
	}
	return ownerAddress.String()
}

func (t *T) Decimals() uint8 {
	decimals, err := tokenRpc.token.TokenCaller.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Get Decimals Error: ", err)
	}
	return decimals
}

func (t *T) BalanceOf(address string) string {
	balanceOf, err := tokenRpc.token.TokenCaller.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		log.Fatal("Get BalanceOf Error: ", err)
	}
	return balanceOf.String()
}

func (t *T) Allowance(addr0, addr1 string) string {
	allowance, err := tokenRpc.token.TokenCaller.Allowance(&bind.CallOpts{}, common.HexToAddress(addr0), common.HexToAddress(addr1))
	if err != nil {
		log.Fatal("Get BalanceOf Error: ", err)
	}
	return allowance.String()
}

func (t *T) AlTrlowance(addr0, addr1 string) string {
	privateKey, err := crypto.HexToECDSA(ACCOUNT_PRIV_KEY)
	if err != nil {
		log.Fatal("Err: ", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := tokenRpc.conn.PendingNonceAt(context.Background(), fromAddress)

	allowance, err := tokenRpc.token.TokenTransactor.Transfer(&bind.TransactOpts{
		From:  fromAddress,
		Nonce: big.NewInt(int64(nonce)),
	}, common.HexToAddress(addr0), big.NewInt(10))
	if err != nil {
		log.Fatal("Get BalanceOf Error: ", err)
	}
	return allowance.Value().String()
}

// 自己发送交易
func (t *T) Transfer(toAddress string, count int64) {
	privateKey, err := crypto.HexToECDSA(ACCOUNT_PRIV_KEY)
	if err != nil {
		log.Fatal("Err: ", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddr := common.HexToAddress(toAddress)
	nonce, err := tokenRpc.conn.PendingNonceAt(context.Background(), fromAddress)
	value := big.NewInt(int64(math.Pow(10, 18)) * count) // 1 wei
	gasLimit := uint64(21000)
	gasFeeCap, err := tokenRpc.conn.SuggestGasPrice(context.Background())
	gasTipCap, err := tokenRpc.conn.SuggestGasTipCap(context.Background())
	var data []byte
	chainID, err := tokenRpc.conn.NetworkID(context.Background())
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Gas:       gasLimit,
		To:        &toAddr,
		Value:     value,
		Data:      data,
	})
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	err = tokenRpc.conn.SendTransaction(context.Background(), signedTx)
}

func (t *T) TransferFrom(privKey string, toAddress string, count int64) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal("Err: ", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddr := common.HexToAddress(toAddress)
	nonce, err := tokenRpc.conn.PendingNonceAt(context.Background(), fromAddress)
	value := big.NewInt(int64(math.Pow(10, 18)) * count) // 1 wei
	gasLimit := uint64(21000)
	gasFeeCap, err := tokenRpc.conn.SuggestGasPrice(context.Background())
	gasTipCap, err := tokenRpc.conn.SuggestGasTipCap(context.Background())
	var data []byte
	chainID, err := tokenRpc.conn.NetworkID(context.Background())
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Gas:       gasLimit,
		To:        &toAddr,
		Value:     value,
		Data:      data,
	})
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	err = tokenRpc.conn.SendTransaction(context.Background(), signedTx)
}
