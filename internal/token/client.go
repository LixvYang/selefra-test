package token

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type T struct{}

const (
	JSON                = `{"version":3,"id":"69f11cc4-7866-42bf-bf04-81bb0b021f1b","address":"0eecd003aa72554354527bed08dc88f971d881df","crypto":{"ciphertext":"0862b95ac264cb16324f4b5c26aac9727c7a14ae2ce35c0514b7b4f18828ee9d","cipherparams":{"iv":"0484e8078c0bbbb7dc27d1944ae08a47"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"8b20b0226c6c6f435c374ebcaf963f9c6dc93fedf4fdcea79fb0a2faddec389c","n":262144,"r":8,"p":1},"mac":"0545955a32b69b226703c9e1f14537c1abb6a77a268694ee7a1f02156ef370bd"}}`
	BSC_RPC_URL         = "https://endpoints.omniatech.io/v1/bsc/testnet/public"
	CONTRACT_ADDR       = "0xEf42420F5d2815CbB2700d03D527F0F89bdA9503"
	MY_ACCOUNT_PUB_KEY  = "0x0eecD003Aa72554354527BeD08dC88f971d881DF"
	MY_ACCOUNT_PRIV_KEY = "f4a7287be8fde0ea3a350956ab2c0e49c0c1e63be354d28c1e1cca8b3978af8e"
)

var tokenRpc TokenRPC
var err error

type TokenRPC struct {
	conn  *ethclient.Client
	token *Token
}

func Init() {
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
	privateKey, err := crypto.HexToECDSA(MY_ACCOUNT_PRIV_KEY)
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
func (t *T) Transfer(toAddress string, count int64) error {
	privateKey, err := crypto.HexToECDSA(MY_ACCOUNT_PRIV_KEY)
	if err != nil {
		log.Fatal("Err: ", err)
	}
	fmt.Println(111)

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
	if err != nil {
		fmt.Println("chainID error")
		return err
	}
	fmt.Println(222)
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
	if err != nil {
		fmt.Println(444)
		return err
	}
	fmt.Println(333)

	err = tokenRpc.conn.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(555)
		return err
	}
	return nil
}

// func (t *T) T(toAddress string, count int64) error {
// 	tx, err := tokenRpc.token.Transfer(&bind.TransactOpts{
// 		From:   MY_ACCOUNT_PUB_KEY,
// 		Signer: auth.Signer,
// 		Value:  nil,
// 	}, common.HexToAddress(toAddress), big.NewInt(520))

// }

func (t *T) T(toAddress string, count int64) error {
	chainID, err := tokenRpc.conn.NetworkID(context.Background())
	if err != nil {
		fmt.Println("chainID error")
		return err
	}
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(JSON), "night-wear-prevent-area-rail-concert-slab-intact-mutual-romance-defense-dish", chainID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	tx, err := tokenRpc.token.Transfer(auth, common.HexToAddress(toAddress), big.NewInt(int64(math.Pow(10, 15)*float64(count))))
	if err != nil {
		fmt.Println("token.Transfer 错误")
		log.Fatalf("Transfer err: %s", err)
	}
	// 等待交易完成
	ctx := context.Background()
	_, err = bind.WaitMined(ctx, tokenRpc.conn, tx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Transfer Successful!")
	return nil
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

// 授权B可以调用其中100个代币 —— approve(B，100)；
func (t *T) Approve(_spender string, _value int64) {
	_, err := tokenRpc.token.TokenTransactor.Approve(&bind.TransactOpts{}, common.HexToAddress(_spender), big.NewInt(_value))
	if err != nil {
		return
	}
}
