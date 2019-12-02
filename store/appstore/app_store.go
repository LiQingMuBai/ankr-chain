package appstore

import (
	"math/big"

	ankrcmm "github.com/Ankr-network/ankr-chain/common"
	ankrapscmm "github.com/Ankr-network/ankr-chain/store/appstore/common"
	"github.com/tendermint/iavl"
	"github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
)

type AccountStore interface {
	Nonce(address string, height int64, prove bool) (uint64, string, *iavl.RangeProof, []byte, error)
	SetNonce(address string, nonce uint64) error
	IncNonce(address string) (uint64, error)
	AddAccount(address string, accType ankrcmm.AccountType)
	SetBalance(address string, amount ankrcmm.Amount)
	Balance(address string, symbol string, height int64, prove bool) (*big.Int, string, *iavl.RangeProof, []byte, error)
	SetAllowance(addrSender string, addrSpender string, amount ankrcmm.Amount)
	Allowance(addrSender string, addrSpender string, symbol string) (*big.Int, error)
}

type BCStore interface {
	Height() int64
}

type TxStore interface {
	Commit() types.ResponseCommit
	SetCertKey(dcName string, pemBase64 string)
	CertKey(dcName string, height int64, prove bool) (string, string, *iavl.RangeProof, []byte)
	DeleteCertKey(dcName string)
	SetMetering(dcName string, nsName string, value string)
	Metering(dcName string, nsName string, height int64, prove bool) (string, string, *iavl.RangeProof, []byte)
	SetValidator(valInfo *ankrcmm.ValidatorInfo)
	Validator(valAddr string, height int64, prove bool) (*ankrcmm.ValidatorInfo, string, *iavl.RangeProof, []byte, error)
	RemoveValidator(valAddr string)
	TotalValidatorPowers() int64
	Get(key []byte) []byte
	Set(key []byte, val []byte)
	Delete(key []byte)
	Has(key []byte) bool
	TotalTx(height int64, prove bool) (int64, string, *iavl.RangeProof, []byte, error)
	SetTotalTx(totalTx int64)
	IncTotalTx() int64
}

type ContractStore interface {
	IsExist(cAddr string) bool
	CreateCurrency(symbol string, currency *ankrcmm.Currency) error
	CurrencyInfo(symbol string) (*ankrcmm.Currency, error)
	BuildCurrencyCAddrMap(symbol string, cAddr string) error
	ContractAddrBySymbol(symbol string) (string, error)
	SaveContract(cAddr string, cInfo *ankrcmm.ContractInfo) error
	LoadContract(cAddr string, height int64, prove bool) (*ankrcmm.ContractInfo, string, *iavl.RangeProof, []byte, error)
}

type QueryHandler interface {
	Query(reqQuery types.RequestQuery) (resQuery types.ResponseQuery)
}

type AppStore interface {
	AccountStore
	TxStore
	QueryHandler
	ContractStore
	BCStore
	SetChainID(chainID string)
	ChainID() string
	APPHash() []byte
	APPHashByHeight(height int64) []byte
	KVState() ankrapscmm.State
	ResetKVState()
    DB() dbm.DB
}