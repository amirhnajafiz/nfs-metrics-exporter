package cmd

type CMDHealth struct{}

func (c *CMDHealth) Command() string {
	return "health"
}

func (c *CMDHealth) Run() error {
	return nil
}
