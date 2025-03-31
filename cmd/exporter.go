package cmd

type CMDExporter struct{}

func (c *CMDExporter) Command() string {
	return "exporter"
}

func (c *CMDExporter) Run() error {
	return nil
}
