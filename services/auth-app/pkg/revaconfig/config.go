package revaconfig

import (
	"path/filepath"

	"github.com/opencloud-eu/opencloud/pkg/config/defaults"
	"github.com/opencloud-eu/opencloud/services/auth-app/pkg/config"
)

// AuthAppConfigFromStruct will adapt an OpenCloud config struct into a reva mapstructure to start a reva service.
func AuthAppConfigFromStruct(cfg *config.Config) map[string]interface{} {
	appAuthJSON := filepath.Join(defaults.BaseDataPath(), "appauth.json")

	rcfg := map[string]interface{}{
		"shared": map[string]interface{}{
			"jwt_secret":                cfg.TokenManager.JWTSecret,
			"gatewaysvc":                cfg.Reva.Address,
			"skip_user_groups_in_token": cfg.SkipUserGroupsInToken,
			"grpc_client_options":       cfg.Reva.GetGRPCClientConfig(),
		},
		"grpc": map[string]interface{}{
			"network": cfg.GRPC.Protocol,
			"address": cfg.GRPC.Addr,
			"tls_settings": map[string]interface{}{
				"enabled":     cfg.GRPC.TLS.Enabled,
				"certificate": cfg.GRPC.TLS.Cert,
				"key":         cfg.GRPC.TLS.Key,
			},
			"services": map[string]interface{}{
				"authprovider": map[string]interface{}{
					"auth_manager": "appauth",
					"auth_managers": map[string]interface{}{
						"appauth": map[string]interface{}{
							"gateway_addr": cfg.Reva.Address,
						},
					},
				},
				"applicationauth": map[string]interface{}{
					"driver": "jsoncs3",
					"drivers": map[string]interface{}{
						"json": map[string]interface{}{
							"file": appAuthJSON,
						},
						"jsoncs3": map[string]interface{}{
							"provider_addr":       cfg.StorageDrivers.JSONCS3.ProviderAddr,
							"service_user_id":     cfg.StorageDrivers.JSONCS3.SystemUserID,
							"service_user_idp":    cfg.StorageDrivers.JSONCS3.SystemUserIDP,
							"machine_auth_apikey": cfg.StorageDrivers.JSONCS3.SystemUserAPIKey,
						},
					},
				},
			},
			"interceptors": map[string]interface{}{
				"prometheus": map[string]interface{}{
					"namespace": "opencloud",
					"subsystem": "auth_app",
				},
			},
		},
	}
	return rcfg
}
