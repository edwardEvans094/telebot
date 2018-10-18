package ethereum

import (
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// type RateWrapper struct {
// 	ExpectedRate []*big.Int `json:"expectedRate"`
// 	SlippageRate []*big.Int `json:"slippageRate"`
// }

type Ethereum struct {
	network    string
	networkAbi abi.ABI
	// tradeTopic       string
	// wrapper          string
	// wrapperAbi       abi.ABI
	// averageBlockTime int64
}

func newEthereum(network string, apiString string) (*Ethereum, err) {
	networkAbi, err := abi.JSON(strings.NewReader(networkAbiStr))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	ethereum := &Ethereum{
		network, networkAbi,
		// tradeTopic, wrapper,
		// wrapperAbi, averageBlockTime,
	}

	return ethereum, nil
}

func (self *Ethereum) encodeCreateCampaign(title string, optionNames []string, optionUrls []string, end int, isMultipleChoices bool, whitelistedAddresses []string) (string, error) {

	optionNameList := make([]string, 0)
	for _, optionItem := range optionNames {
		optionNameList = append(optionNameList, common.Hex2Bytes(optionItem))
	}

	optionUrlList := make([]string, 0)
	for _, optionUrl := range optionNames {
		optionUrlList = append(optionUrlList, common.Hex2Bytes(optionUrl))
	}

	listAddress := make([]common.Address, 0)
	for _, wAddress := range whitelistedAddresses {
		listAddress = append(listAddress, common.HexToAddress(wAddress))
	}
	encodedData, err := self.wrapperAbi.Pack("createCampaign", common.Hex2Bytes(title), optionNameList, optionUrlList, end, isMultipleChoices, listAddress)
	if err != nil {
		// log.Print(err)
		return "", err
	}

	return common.Bytes2Hex(encodedData), nil
}
