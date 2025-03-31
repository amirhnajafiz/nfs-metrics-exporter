package cmd

type CMDRequirements struct{}

func (c *CMDRequirements) Command() string {
	return "requirements"
}

func (c *CMDRequirements) Run() error {
	return nil
}
