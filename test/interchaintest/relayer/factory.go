package relayer

import (
	"github.com/docker/docker/client"
	ictest "github.com/strangelove-ventures/interchaintest/v6"
	"github.com/strangelove-ventures/interchaintest/v6/ibc"
	"github.com/strangelove-ventures/interchaintest/v6/relayer"
	"github.com/strangelove-ventures/interchaintest/v6/relayer/rly"

	"go.uber.org/zap"
)

// builtinRelayerFactory is the built-in relayer factory that understands
// how to start the cosmos relayer in a docker container.
type icqFactory struct {
	impl    ibc.RelayerImplementation
	log     *zap.Logger
	options relayer.RelayerOptions
}

func NewBuiltinRelayerFactory(impl ibc.RelayerImplementation, logger *zap.Logger, options ...relayer.RelayerOption) ictest.RelayerFactory {
	return icqFactory{impl: impl, log: logger, options: options}
}

// Build returns a relayer chosen depending on f.impl.
func (f icqFactory) Build(
	t ictest.TestName,
	cli *client.Client,
	networkID string,
) ibc.Relayer {
	return NewIcqRelayer(f.log, t.Name(), cli, networkID, f.options...)
}

// TODO: This should be a more descriptive name.
func (f icqFactory) Name() string {
	return "icq@" + "123"

}

// Capabilities returns the set of capabilities for the
// relayer implementation backing this factory.
func (f icqFactory) Capabilities() map[relayer.Capability]bool {
	return rly.Capabilities()
}
