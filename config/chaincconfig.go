package config

type ChainConfig map[string]interface{}

var currentConfig *ChainConfig

var (
	ChainIDUnknown = "n/a"
	ChainIDBTS     = "4018d7844c78f6a6c41c6a552b898022310fc5dec06da467ee7905a8dad512c8"
	ChainIDMuse    = "45ad2d3f9ef92a49b55c2227eb06123f613bb35dd08bd876f2aea21925a67a67"
	ChainIDTest    = "39f5e2ede1f8bc1a3a54a7914414e3779e33193f1f5693510e73cb7a87617447"
	ChainIDObelisk = "1cfde7c388b9e8ac06462d68aadbd966b58f88797637d9af805b4560b0e9661e"
	ChainIDGPH     = "b8d1603965b3eb1acba27e62ff59f74efa3154d43a4188d381088ac7cdf35539"
	ChainIDSPM     = ""
	ChainIDSPMTest = "3b3b6ec2827fe6b27805203264e035e4a94f6ba95da7a14f7746b6f75612e676"
	ChainIDAIS     = "32a81eefffe46f2b486d9c7c6b3990f8a1c1d97713eed46bfbb3925d2a4e986e"
)

var (
	knownNetworks = []ChainConfig{
		ChainConfig{
			"name":           "Unknown",
			"core_asset":     "n/a",
			"address_prefix": "n/a",
			"chain_id":       ChainIDUnknown,
		},
		ChainConfig{
			"name":           "Graphene",
			"core_asset":     "CORE",
			"address_prefix": "GPH",
			"chain_id":       ChainIDGPH,
		},
		ChainConfig{
			"name":           "BitShares",
			"core_asset":     "BTS",
			"address_prefix": "BTS",
			"chain_id":       ChainIDBTS,
		},
		ChainConfig{
			"name":           "Muse",
			"core_asset":     "MUSE",
			"address_prefix": "MUSE",
			"chain_id":       ChainIDMuse,
		},
		ChainConfig{
			"name":           "Test",
			"core_asset":     "TEST",
			"address_prefix": "TEST",
			"chain_id":       ChainIDTest,
		},
		ChainConfig{
			"name":           "Obelisk",
			"core_asset":     "GOV",
			"address_prefix": "FEW",
			"chain_id":       ChainIDObelisk,
		},
		ChainConfig{
			"name":           "SPM",
			"core_asset":     "SPM",
			"address_prefix": "SPM",
			"chain_id":       ChainIDSPM,
		},
		ChainConfig{
			"name":           "SPM",
			"core_asset":     "SPM",
			"address_prefix": "SPM",
			"chain_id":       ChainIDSPMTest,
		}, ChainConfig{
			"name":           "AIS",
			"core_asset":     "AIS",
			"address_prefix": "AIS",
			"chain_id":       ChainIDAIS,
		},
	}
)

func (p ChainConfig) ID() string {
	if id, ok := p["chain_id"]; ok {
		return id.(string)
	}

	return "n/a"
}

func (p ChainConfig) Prefix() string {
	if id, ok := p["address_prefix"]; ok {
		return id.(string)
	}

	return "n/a"
}

func CurrentConfig() *ChainConfig {
	return currentConfig
}

var DefaultChain = ChainConfig{
	"name":           "BAR",
	"core_asset":     "BAR",
	"address_prefix": "BAR",
}

func SetCurrentConfig(chainID string) error {
	for _, cnf := range knownNetworks {
		if cnf["chain_id"] == chainID {
			currentConfig = &cnf
			return nil
		}
	}
	
	DefaultChain["chain_id"] = chainID
	
	currentConfig = &DefaultChain
	return nil

	//return errors.Errorf("ChainConfig for chainID %q not found", chainID)
}

func CustomCurrentConfig(cfg *ChainConfig) error {
	currentConfig = cfg
	return nil
}
