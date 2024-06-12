package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/docker/docker/client"
	"github.com/strangelove-ventures/interchaintest/v6/ibc"
	"github.com/strangelove-ventures/interchaintest/v6/relayer"
	"go.uber.org/zap"
)

const (
	RlyDefaultUidGid = "100:1000"
)

const (
	DefaultContainerImage   = "ghcr.io/cosmos/relayer"
	DefaultContainerVersion = "v2.5.2"
)

func NewIcqRelayer(log *zap.Logger, testName string, cli *client.Client, networkID string, options ...relayer.RelayerOption) *IcqRelayer {
	c := Commander{Log: log}
	for _, opt := range options {
		switch o := opt.(type) {
		case relayer.RelayerOptionExtraStartFlags:
			c.ExtraStartFlags = o.Flags
		}
	}
	dr, err := relayer.NewDockerRelayer(context.TODO(), log, testName, cli, networkID, c, options...)
	if err != nil {
		panic(err) // TODO: return
	}

	r := &IcqRelayer{
		DockerRelayer: dr,
	}

	return r
}

func (Commander) Name() string {
	return "rly"
}

func (Commander) DockerUser() string {
	return RlyDefaultUidGid // docker run -it --rm --entrypoint echo ghcr.io/cosmos/relayer "$(id -u):$(id -g)"
}

func (Commander) AddChainConfiguration(containerFilePath, homeDir string) []string {
	return []string{
		"rly", "chains", "add", "-f", containerFilePath,
		"--home", homeDir,
	}
}

func (Commander) AddKey(chainID, keyName, coinType, homeDir string) []string {
	return []string{
		"rly", "keys", "add", chainID, keyName,
		"--coin-type", fmt.Sprint(coinType), "--home", homeDir,
	}
}

func (Commander) CreateChannel(pathName string, opts ibc.CreateChannelOptions, homeDir string) []string {
	return []string{
		"rly", "tx", "channel", pathName,
		"--src-port", opts.SourcePortName,
		"--dst-port", opts.DestPortName,
		"--order", opts.Order.String(),
		"--version", opts.Version,

		"--home", homeDir,
	}
}

func (Commander) CreateClients(pathName string, opts ibc.CreateClientOptions, homeDir string) []string {
	return []string{
		"rly", "tx", "clients", pathName, "--client-tp", opts.TrustingPeriod,
		"--home", homeDir,
	}
}

// passing a value of 0 for customeClientTrustingPeriod will use default
func (Commander) CreateClient(pathName, homeDir, customeClientTrustingPeriod string) []string {
	return []string{
		"rly", "tx", "client", pathName, "--client-tp", customeClientTrustingPeriod,
		"--home", homeDir,
	}
}

func (Commander) CreateConnections(pathName string, homeDir string) []string {
	return []string{
		"rly", "tx", "connection", pathName,
		"--home", homeDir,
	}
}

func (Commander) Flush(pathName, channelID, homeDir string) []string {
	cmd := []string{"rly", "tx", "flush"}
	if pathName != "" {
		cmd = append(cmd, pathName)
		if channelID != "" {
			cmd = append(cmd, channelID)
		}
	}
	cmd = append(cmd, "--home", homeDir)
	return cmd
}

func (Commander) GeneratePath(srcChainID, dstChainID, pathName, homeDir string) []string {
	return []string{
		"rly", "paths", "new", srcChainID, dstChainID, pathName,
		"--home", homeDir,
	}
}

func (Commander) UpdatePath(pathName, homeDir string, filter ibc.ChannelFilter) []string {
	return []string{
		"rly", "paths", "update", pathName,
		"--home", homeDir,
		"--filter-rule", filter.Rule,
		"--filter-channels", strings.Join(filter.ChannelList, ","),
	}
}

func (Commander) GetChannels(chainID, homeDir string) []string {
	return []string{
		"rly", "q", "channels", chainID,
		"--home", homeDir,
	}
}

func (Commander) GetConnections(chainID, homeDir string) []string {
	return []string{
		"rly", "q", "connections", chainID,
		"--home", homeDir,
	}
}

func (Commander) GetClients(chainID, homeDir string) []string {
	return []string{
		"rly", "q", "clients", chainID,
		"--home", homeDir,
	}
}

func (Commander) LinkPath(pathName, homeDir string, channelOpts ibc.CreateChannelOptions, clientOpt ibc.CreateClientOptions) []string {
	return []string{
		"rly", "tx", "link", pathName,
		"--src-port", channelOpts.SourcePortName,
		"--dst-port", channelOpts.DestPortName,
		"--order", channelOpts.Order.String(),
		"--version", channelOpts.Version,
		"--client-tp", clientOpt.TrustingPeriod,
		"--debug",

		"--home", homeDir,
	}
}

func (Commander) RestoreKey(chainID, keyName, coinType, mnemonic, homeDir string) []string {
	return []string{
		"rly", "keys", "restore", chainID, keyName, mnemonic,
		"--coin-type", fmt.Sprint(coinType), "--home", homeDir,
	}
}

func (c Commander) StartRelayer(homeDir string, pathNames ...string) []string {
	cmd := []string{
		"rly", "start", "--debug",
		"--home", homeDir,
	}
	cmd = append(cmd, c.ExtraStartFlags...)
	cmd = append(cmd, pathNames...)
	return cmd
}

func (Commander) UpdateClients(pathName, homeDir string) []string {
	return []string{
		"rly", "tx", "update-clients", pathName,
		"--home", homeDir,
	}
}

type CosmosRelayerChainConfigValue struct {
	AccountPrefix  string  `json:"account-prefix"`
	ChainID        string  `json:"chain-id"`
	Debug          bool    `json:"debug"`
	GRPCAddr       string  `json:"grpc-addr"`
	GasAdjustment  float64 `json:"gas-adjustment"`
	GasPrices      string  `json:"gas-prices"`
	Key            string  `json:"key"`
	KeyringBackend string  `json:"keyring-backend"`
	OutputFormat   string  `json:"output-format"`
	RPCAddr        string  `json:"rpc-addr"`
	SignMode       string  `json:"sign-mode"`
	Timeout        string  `json:"timeout"`
}

type CosmosRelayerChainConfig struct {
	Type  string                        `json:"type"`
	Value CosmosRelayerChainConfigValue `json:"value"`
}

func ChainConfigToCosmosRelayerChainConfig(chainConfig ibc.ChainConfig, keyName, rpcAddr, gprcAddr string) CosmosRelayerChainConfig {
	chainType := chainConfig.Type
	if chainType == "polkadot" || chainType == "parachain" || chainType == "relaychain" {
		chainType = "substrate"
	}
	return CosmosRelayerChainConfig{
		Type: chainType,
		Value: CosmosRelayerChainConfigValue{
			Key:            keyName,
			ChainID:        chainConfig.ChainID,
			RPCAddr:        rpcAddr,
			GRPCAddr:       gprcAddr,
			AccountPrefix:  chainConfig.Bech32Prefix,
			KeyringBackend: keyring.BackendTest,
			GasAdjustment:  chainConfig.GasAdjustment,
			GasPrices:      chainConfig.GasPrices,
			Debug:          true,
			Timeout:        "10s",
			OutputFormat:   "json",
			SignMode:       "direct",
		},
	}
}

func (Commander) ConfigContent(ctx context.Context, cfg ibc.ChainConfig, keyName, rpcAddr, grpcAddr string) ([]byte, error) {
	cosmosRelayerChainConfig := ChainConfigToCosmosRelayerChainConfig(cfg, keyName, rpcAddr, grpcAddr)
	jsonBytes, err := json.Marshal(cosmosRelayerChainConfig)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func (Commander) DefaultContainerImage() string {
	return DefaultContainerImage
}

func (Commander) DefaultContainerVersion() string {
	return DefaultContainerVersion
}

func (Commander) ParseAddKeyOutput(stdout, stderr string) (ibc.Wallet, error) {
	var wallet WalletModel
	err := json.Unmarshal([]byte(stdout), &wallet)
	rlyWallet := NewWallet("", wallet.Address, wallet.Mnemonic)
	return rlyWallet, err
}

func (Commander) ParseRestoreKeyOutput(stdout, stderr string) string {
	return strings.Replace(stdout, "\n", "", 1)
}

func (c Commander) ParseGetChannelsOutput(stdout, stderr string) ([]ibc.ChannelOutput, error) {
	var channels []ibc.ChannelOutput
	channelSplit := strings.Split(stdout, "\n")
	for _, channel := range channelSplit {
		if strings.TrimSpace(channel) == "" {
			continue
		}
		var channelOutput ibc.ChannelOutput
		err := json.Unmarshal([]byte(channel), &channelOutput)
		if err != nil {
			c.Log.Error("Failed to parse channels json", zap.Error(err))
			continue
		}
		channels = append(channels, channelOutput)
	}

	return channels, nil
}

func (c Commander) ParseGetConnectionsOutput(stdout, stderr string) (ibc.ConnectionOutputs, error) {
	var connections ibc.ConnectionOutputs
	for _, connection := range strings.Split(stdout, "\n") {
		if strings.TrimSpace(connection) == "" {
			continue
		}

		var connectionOutput ibc.ConnectionOutput
		if err := json.Unmarshal([]byte(connection), &connectionOutput); err != nil {
			c.Log.Error(
				"Error parsing connection json",
				zap.Error(err),
			)

			continue
		}
		connections = append(connections, &connectionOutput)
	}

	return connections, nil
}

func (c Commander) ParseGetClientsOutput(stdout, stderr string) (ibc.ClientOutputs, error) {
	var clients ibc.ClientOutputs
	for _, client := range strings.Split(stdout, "\n") {
		if strings.TrimSpace(client) == "" {
			continue
		}

		var clientOutput ibc.ClientOutput
		if err := json.Unmarshal([]byte(client), &clientOutput); err != nil {
			c.Log.Error(
				"Error parsing client json",
				zap.Error(err),
			)

			continue
		}
		clients = append(clients, &clientOutput)
	}

	return clients, nil
}

func (Commander) Init(homeDir string) []string {
	return []string{
		"rly", "config", "init",
		"--home", homeDir,
	}
}

func (c Commander) CreateWallet(keyName, address, mnemonic string) ibc.Wallet {
	return NewWallet(keyName, address, mnemonic)
}

// Commander satisfies relayer.RelayerCommander.
type Commander struct {
	Log             *zap.Logger
	ExtraStartFlags []string
}

// IcqRelayer is the ibc.Relayer implementation for github.com/quicksilver/icq-relayer.
type IcqRelayer struct {
	// Embedded DockerRelayer so commands just work.
	*relayer.DockerRelayer
}

var _ ibc.Wallet = &RlyWallet{}

type WalletModel struct {
	Mnemonic string `json:"mnemonic"`
	Address  string `json:"address"`
}

type RlyWallet struct {
	mnemonic string
	address  string
	keyName  string
}

func NewWallet(keyname string, address string, mnemonic string) *RlyWallet {
	return &RlyWallet{
		mnemonic: mnemonic,
		address:  address,
		keyName:  keyname,
	}
}

func (w *RlyWallet) KeyName() string {
	return w.keyName
}

func (w *RlyWallet) FormattedAddress() string {
	return w.address
}

// Get mnemonic, only used for relayer wallets
func (w *RlyWallet) Mnemonic() string {
	return w.mnemonic
}

// Get Address
func (w *RlyWallet) Address() []byte {
	return []byte(w.address)
}
