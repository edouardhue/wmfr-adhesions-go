package iraiser

type IRaiser struct {
	config *Config
}

func NewIRaiser(config *Config) *IRaiser {
	return &IRaiser{
		config: config,
	}
}