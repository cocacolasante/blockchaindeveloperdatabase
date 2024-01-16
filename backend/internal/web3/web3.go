package web3

import (
	"fmt"
	"log"
	"math/big"
	"os"

	credits "github.com/cocacolasante/blockchaindeveloperdatabase/smartcontractinterfaces"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Web3Connect struct {
	ChainId    int
	Client     *ethclient.Client
	Blockchain string
}

const CHAINID =80001

func GetClient() (*ethclient.Client, error) {
	var client *ethclient.Client

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
	instance, err := credits.NewCredits(contAdd, w.Client)
	if err != nil {
		return nil, err

	}

	remainingCredits, err := instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		return nil, err
	}

	return remainingCredits, nil

}

func (w *Web3Connect) RedeemCredits(requestersAddress string) error {

	privateKey := os.Getenv("DEPLOYER_PRIVATE_KEY")
	contractAddy := os.Getenv("CREDIT_CONTRACT")
	wallet, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Println(err)
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(wallet, big.NewInt(CHAINID)) // Use ChainID 1 for the mainnet
	if err != nil {
		log.Println(err)
		return err
	}

	auth.GasPrice = big.NewInt(20000000000) // Replace with your preferred gas price
	auth.GasLimit = uint64(300000)

	contractAddress := common.HexToAddress(contractAddy)
	contractInstance, err := credits.NewCredits(contractAddress, w.Client)
	if err != nil {
		log.Println(err)
		return err
	}

	targetAddress := common.HexToAddress(requestersAddress)
	tx, err := contractInstance.RedeemCredits(auth, targetAddress, big.NewInt(1))
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())

	return nil
}
