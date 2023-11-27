package midtrans

import (
	"github.com/arvinpaundra/ngekost-api/pkg/util/config"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type Midtrans struct {
	serverKey   string
	clientKey   string
	appEnv      string
	midtransEnv midtrans.EnvironmentType
}

func New() *Midtrans {
	m := Midtrans{
		serverKey: config.GetString("MIDTRANS_SERVER_KEY"),
		clientKey: config.GetString("MIDTRANS_CLIENT_KEY"),
		appEnv:    config.GetString("APP_ENV"),
	}

	midtrans.ServerKey = m.serverKey
	midtrans.Environment = midtrans.Sandbox
	m.midtransEnv = midtrans.Sandbox

	if m.appEnv == "production" {
		midtrans.Environment = midtrans.Production
		m.midtransEnv = midtrans.Production
	}

	return &m
}

func (m *Midtrans) Snap() *snap.Client {
	sc := snap.Client{
		ServerKey: m.serverKey,
		Env:       m.midtransEnv,
		Options:   &midtrans.ConfigOptions{},
	}

	return &sc
}
