// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package tests

import (
	"testing"
)

func TestBlockchain(t *testing.T) {
	t.Parallel()

	bt := new(testMatcher)
	// General state tests are 'exported' as blockchain tests, but we can run them natively.
	bt.skipLoad(`^GeneralStateTests/`)
	// Skip random failures due to selfish mining test.
	bt.skipLoad(`^bcForgedTest/bcForkUncle\.json`)
	bt.skipLoad(`^bcMultiChainTest/(ChainAtoChainB_blockorder|CallContractFromNotBestBlock)`)
	bt.skipLoad(`^bcTotalDifficultyTest/(lotsOfLeafs|lotsOfBranches|sideChainWithMoreTransactions)`)
	// Constantinople is not implemented yet.
	bt.skipLoad(`(?i)(constantinople)`)
	// Expected failures:
	bt.fails("^TransitionTests/bcEIP158ToByzantium", "byzantium not supported")
	bt.fails(`^TransitionTests/bcHomesteadToDao/DaoTransactions(|_UncleExtradata|_EmptyTransactionAndForkBlocksAhead)\.json`, "issue in test")
	bt.fails(`^bc(Exploit|Fork|Gas|Multi|Total|State|Random|Uncle|Valid|Wallet).*_Byzantium$`, "byzantium not supported")
	bt.fails(`^bcBlockGasLimitTest/(BlockGasLimit2p63m1|TransactionGasHigherThanLimit2p63m1|SuicideTransaction|GasUsedHigherThanBlockGasLimitButNotWithRefundsSuicideFirst|TransactionGasHigherThanLimit2p63m1_2).*_Byzantium$`, "byzantium not supported")

	// Skip invalid receipt root hash / invalid nonce quorum failures
	bt.skipLoad(`(TransactionSendingToZero|wrongParentHash|wrongMixHash|wrongStateRoot|timestampTooHigh|timestampTooLow|nonceWrong|gasLimitTooLowExactBound|gasLimitTooLow|gasLimitTooHighExactBound|gasLimitTooHigh|diffTooLow2|diffTooLow|diffTooHigh|suicideCoinbase|InternlCallStoreClearsSucces|StoreClearsAndInternlCallStoreClearsOOG|failed_tx_xcf416c53|TransactionSendingToZero|SuicidesAndInternlCallSuicidesSuccess|SuicidesAndInternlCallSuicidesOOG|SuicidesAndInternlCallSuicidesBonusGasAtCallFailed|SuicidesAndInternlCallSuicidesBonusGasAtCall|StoreClearsAndInternlCallStoreClearsSuccess|CallContractToCreateContractOOG).json`)

	bt.walk(t, blockTestDir, func(t *testing.T, name string, test *BlockTest) {
		if err := bt.checkFailure(t, name, test.Run()); err != nil {
			t.Error(err)
		}
	})
}
