package types

import (
	"gopkg.in/yaml.v2"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// ParamKeyTable for airdrop module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new airdrop Params instance.
func NewParams() Params {
	return Params{}
}

// DefaultParams default ics params.
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs implements params.ParamSet.
func (*Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		// paramtypes.NewParamSetPair(Key, &p.Attribute, validateAttrib),
	}
}

// Validate validates params.
func (*Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
