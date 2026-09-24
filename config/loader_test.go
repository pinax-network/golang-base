package base_config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type testServiceClient struct {
	ClientID      string `mapstructure:"client_id" validate:"required"`
	PrincipalGUID string `mapstructure:"principal_guid" validate:"required"`
}

type testAuthConfig struct {
	Domain         string              `mapstructure:"auth0_domain" validate:"required"`
	Secret         string              `mapstructure:"secret" validate:"required"`
	ServiceClients []testServiceClient `mapstructure:"admin_service_clients" validate:"dive"`
}

type testConfig struct {
	Application ApplicationConfig `mapstructure:"application"`
	Auth0       testAuthConfig    `mapstructure:"auth0" validate:"required"`
}

const secretYAML = `
application:
  gin_mode: release
auth0:
  auth0_domain: login.example.com
  secret: do-not-print-this-secret
`

const publicYAML = `
auth0:
  admin_service_clients:
    - client_id: MixedCaseClientID
      principal_guid: pinax-admin-ui
`

func writeConfig(t *testing.T, name, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	require.NoError(t, os.WriteFile(path, []byte(content), 0o600))
	return path
}

func TestLoadSingleFile(t *testing.T) {
	var config testConfig
	require.NoError(t, Load(writeConfig(t, "config.yaml", secretYAML), &config))
	require.Equal(t, "login.example.com", config.Auth0.Domain)
	require.Equal(t, Release, config.Application.GinMode)

	err := Load(writeConfig(t, "invalid.yaml", "application:\n  gin_mode: nope\nauth0:\n  auth0_domain: x\n  secret: y\n"), &config)
	require.ErrorContains(t, err, "config validation failed")
}

func TestLoadMergesPublicAndSecretFiles(t *testing.T) {
	public := writeConfig(t, "public.yaml", publicYAML)
	secret := writeConfig(t, "secret.yaml", secretYAML)
	var config testConfig
	require.NoError(t, Load(public+","+secret, &config))
	require.Equal(t, "login.example.com", config.Auth0.Domain)
	require.Equal(t, "do-not-print-this-secret", config.Auth0.Secret)
	require.Equal(t, []testServiceClient{{ClientID: "MixedCaseClientID", PrincipalGUID: "pinax-admin-ui"}}, config.Auth0.ServiceClients)
	require.Equal(t, Release, config.Application.GinMode)
}

func TestLoadRejectsSettingsDefinedInTwoFiles(t *testing.T) {
	for name, override := range map[string]string{
		"same setting":       "auth0:\n  secret: override-attempt\n",
		"value over a map":   "auth0: override-attempt\n",
		"map over a value":   "application:\n  gin_mode:\n    nested: override-attempt\n",
		"list over settings": "auth0:\n  - override-attempt\n",
	} {
		t.Run(name, func(t *testing.T) {
			files := writeConfig(t, "public.yaml", override) + "," + writeConfig(t, "secret.yaml", secretYAML)
			var config testConfig
			err := Load(files, &config)
			require.ErrorContains(t, err, "define each setting in only one file")
			require.NotContains(t, err.Error(), "do-not-print-this-secret")
			require.NotContains(t, err.Error(), "override-attempt")
		})
	}
	t.Run("same list setting", func(t *testing.T) {
		files := writeConfig(t, "a.yaml", publicYAML) + "," + writeConfig(t, "b.yaml", publicYAML)
		var config testConfig
		err := Load(files, &config)
		require.ErrorContains(t, err, "define each setting in only one file")
		require.NotContains(t, err.Error(), "MixedCaseClientID")
	})
}

func TestLoadFileListSyntax(t *testing.T) {
	public := writeConfig(t, "public.yaml", publicYAML)
	secret := writeConfig(t, "secret.yaml", secretYAML)
	var config testConfig
	require.NoError(t, Load(" "+public+" , "+secret+",", &config))
	require.Len(t, config.Auth0.ServiceClients, 1)

	require.ErrorContains(t, Load(" , ", &config), "no config file given")

	missing := filepath.Join(t.TempDir(), "missing.yaml")
	err := Load(secret+","+missing, &config)
	require.ErrorContains(t, err, "failed to read config file "+missing)
}

func TestLoadValidatesTheMergedConfig(t *testing.T) {
	secret := writeConfig(t, "secret.yaml", strings.Replace(secretYAML, "  auth0_domain: login.example.com\n", "", 1))
	domain := writeConfig(t, "public.yaml", "auth0:\n  auth0_domain: login.example.com\n")
	var config testConfig
	require.ErrorContains(t, Load(secret, &config), "config validation failed")
	require.NoError(t, Load(domain+","+secret, &config))
	require.Equal(t, "login.example.com", config.Auth0.Domain)

	incomplete := writeConfig(t, "clients.yaml", "auth0:\n  admin_service_clients:\n    - client_id: MixedCaseClientID\n")
	require.ErrorContains(t, Load(domain+","+incomplete+","+secret, &config), "config validation failed")
}
