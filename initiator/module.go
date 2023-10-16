package initiator

import (
	"visitor_management/internal/module"
	gMod "visitor_management/internal/module/generic_module"
	"visitor_management/internal/module/user"
	"visitor_management/platform/logger"
)

type Module struct {
	// TODO implement
	UserModule    module.UserModule
	EstateModule  module.EstateModule
	GenericModule module.GenericModule
}

func InitModule(persistence Persistence, privateKeyPath string, platformLayer PlatformLayer, log logger.Logger) Module {

	gmod := gMod.InitGenericModule(log, persistence.Generic)
	return Module{
		UserModule:    user.InitOAuth(log, persistence.User),
		GenericModule: gmod,
	}
}
