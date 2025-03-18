package sqlboiler_sqlite3

import (
	"github.com/hdget/common/intf"
	"github.com/pkg/errors"
)

type sqliteProviderConfig struct {
	DbName string `mapstructure:"dbname"`
}

const (
	configSection = "sdk.sqlite3"
)

var (
	errInvalidConfig = errors.New("invalid config")
	errEmptyConfig   = errors.New("empty config")
)

func newConfig(configProvider intf.ConfigProvider) (*sqliteProviderConfig, error) {
	if configProvider == nil {
		return nil, errInvalidConfig
	}

	var c *sqliteProviderConfig
	err := configProvider.Unmarshal(&c, configSection)
	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, errEmptyConfig
	}

	err = c.validate()
	if err != nil {
		return nil, errors.Wrap(err, "validate sqlite3 config")
	}

	return c, nil
}

func (c *sqliteProviderConfig) validate() error {
	if c == nil || c.DbName == "" {
		return errInvalidConfig
	}
	return nil
}
