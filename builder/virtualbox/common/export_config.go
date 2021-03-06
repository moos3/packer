package common

import (
	"errors"
	"fmt"
	"github.com/mitchellh/packer/packer"
)

type ExportConfig struct {
	Format string `mapstruture:"format"`
	Product string `mapstruture:"product"`
	ProductUrl string `mapstructure:"product-url"`
	Vendor string `mapstruture:"vendor"`
	VendorUrl string `mapstruture:"vendor-url"`
	Version string `mapstruture:"version"`
	EULA string `mapstruture:"eula"`
	EULAFile string `mapstruture:"eula-file-path"`
}

func (c *ExportConfig) Prepare(t *packer.ConfigTemplate) []error {
	if c.Format == "" {
		c.Format = "ovf"
	}

	templates := map[string]*string{
		"format": &c.Format,
	}

	errs := make([]error, 0)
	for n, ptr := range templates {
		var err error
		*ptr, err = t.Process(*ptr, nil)
		if err != nil {
			errs = append(errs, fmt.Errorf("Error processing %s: %s", n, err))
		}
	}

	if c.Format != "ovf" && c.Format != "ova" {
		errs = append(errs,
			errors.New("invalid format, only 'ovf' or 'ova' are allowed"))
	}

	return errs
}
