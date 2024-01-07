package web3

import (
	"fmt"
	"math/big"
	"os"

	ierc20 "github.com/cocacolasante/blockchaindeveloperdatabase/smartcontractinterfaces"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Web3Connect struct {
	ChainId    int
	Client     *ethclient.Client
	Blockchain string
}

func GetClient() (*ethclient.Client, error) {
	var client *ethclient.Client

	fmt.Printf("Getting Blockchain Client \n")

	ethURL := os.Getenv("POLYGON_MAINNET_URL")
	conn, err := ethclient.Dial(ethURL)
	if err != nil {
		return nil, err

	}

	client = conn

	return client, nil

}

func (w *Web3Connect) GetRemainingCredits(address string) (*big.Int, error) {

	account := common.HexToAddress(address)
	contractstring := os.Getenv("CREDIT_CONTRACT")
	contAdd := common.HexToAddress(contractstring)
	instance, err := ierc20.NewIerc20(contAdd, w.Client)
	if err != nil {
		return nil, err

	}

	remainingCredits, err := instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		return nil, err
	}

	return remainingCredits, nil

}
