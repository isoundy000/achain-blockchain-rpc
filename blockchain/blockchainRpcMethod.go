package blockchain

import (
	"encoding/json"
	"reflect"
	"2018go-workspace/achain-blockchain/utils"
)

func Info(result string) *BlockInfo {
	if !utils.IsBlank(reflect.ValueOf(result)) {
		info := &BlockInfo{}
		json.Unmarshal([]byte(result), &info)
		return info
	}
	return nil
}

type BlockInfo struct {
	Id     string     `json:"id"`
	Result InfoResult `json:"result"`
}

type InfoResult struct {
	BlockchainHeadBlockNum                 int64   `json:"blockchain_head_block_num"`
	BlockchainHeadBlockAge                 int64   `json:"blockchain_head_block_age"`
	BlockchainHeadBlockTimestamp           string  `json:"blockchain_head_block_timestamp"`
	BlockchainHeadBlockId                  string  `json:"blockchain_head_block_id"`
	BlockchainAverageDelegateParticipation string  `json:"blockchain_average_delegate_participation"`
	BlockchainConfirmationRequirement      int     `json:"blockchain_confirmation_requirement"`
	BlockchainShareSupply                  int64   `json:"blockchain_share_supply"`
	BlockchainBlocksLeftInRound            int     `json:"blockchain_blocks_left_in_round"`
	BlockchainNextRoundTime                int     `json:"blockchain_next_round_time"`
	BlockchainNextRoundTimestamp           string  `json:"blockchain_next_round_timestamp"`
	BlockchainRandomSeed                   string  `json:"blockchain_random_seed"`
	ClientDataDir                          string  `json:"client_data_dir"`
	ClientVersion                          string  `json:"client_version"`
	NetworkNumConnections                  int     `json:"network_num_connections"`
	NetworkNumConnectionsMax               int     `json:"network_num_connections_max"`
	NetworkChainDownloaderRunning          bool    `json:"network_chain_downloader_running"`
	NetworkChainDownloaderBlocksRemaining  string  `json:"network_chain_downloader_blocks_remaining"`
	TtpTime                                string  `json:"ntp_time"`
	TtpTimeError                           float64 `json:"ntp_time_error"`
	WalletOpen                             *bool   `json:"wallet_open"`
	WalletUnlocked                         *bool   `json:"wallet_unlocked"`
}
