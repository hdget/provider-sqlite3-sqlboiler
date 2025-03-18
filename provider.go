package sqlite3_sqlboiler

import (
	"github.com/hdget/common/intf"
	"github.com/hdget/common/types"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	_ "modernc.org/sqlite"
)

type sqlite3Provider struct {
	client intf.DbClient
}

func New(configProvider intf.ConfigProvider, logger intf.LoggerProvider) (intf.DbProvider, error) {
	config, err := newConfig(configProvider)
	if err != nil {
		return nil, err
	}

	client, err := newClient(config)
	if err != nil {
		return nil, errors.Wrap(err, "new sqlite3 client")
	}

	// 设置boil的缺省db
	boil.SetDB(client)

	logger.Debug("init sqlite3 provider", "db", config.DbName)
	return &sqlite3Provider{client: client}, nil
}

// NewClient 从指定的文件创建创建数据库连接
func NewClient(dbFile string) (intf.DbClient, error) {
	client, err := newClient(nil, dbFile)
	if err != nil {
		return nil, errors.Wrapf(err, "connect sqlite3: %s", dbFile)
	}

	// 设置boil的缺省db
	boil.SetDB(client)
	return client, nil
}

func (p *sqlite3Provider) GetCapability() types.Capability {
	return Capability
}

func (p *sqlite3Provider) My() intf.DbClient {
	return p.client
}

func (p *sqlite3Provider) Master() intf.DbClient {
	return p.client
}

func (p *sqlite3Provider) Slave(i int) intf.DbClient {
	return p.client
}

func (p *sqlite3Provider) By(name string) intf.DbClient {
	return p.client
}
