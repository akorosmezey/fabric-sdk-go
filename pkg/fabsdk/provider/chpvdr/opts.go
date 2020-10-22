package chpvdr

import (
	coptions "github.com/akorosmezey/fabric-sdk-go/pkg/common/options"
	"github.com/akorosmezey/fabric-sdk-go/pkg/common/providers/fab"
)


type providerParams struct {
  localProvider   fab.LocalDiscoveryProvider
}

func newDefaultProviderParams() *providerParams {
	return &providerParams{}
}

// WithErrorHandler sets the error handler
func WithLocalDiscovery(value fab.LocalDiscoveryProvider) coptions.Opt {
	return func(p coptions.Params) {
		logger.Debug("Checking LocalDiscoveryProvider")
		if setter, ok := p.(discoveryProviderSetter); ok {
      logger.Debug("Setting LocalDiscoveryProvider")
			setter.SetLocalDiscoveryProvider(value)
		}
	}
}

type discoveryProviderSetter interface {
	SetLocalDiscoveryProvider(value fab.LocalDiscoveryProvider)
}

func (o *providerParams) SetLocalDiscoveryProvider(value fab.LocalDiscoveryProvider) {
	logger.Debugf("LocalDiscoveryProvider: %T", value)
	o.localProvider = value
}
