package config

import (
	"os"

	"github.com/shono09835/bosh-cli/v7/uaa"
	"gopkg.in/yaml.v2"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

/*
environments:
- url: https://192.168.50.4:25555
  ca_cert: |...
  username: admin
  password: admin
*/

type FSConfig struct {
	path string
	fs   boshsys.FileSystem

	schema fsConfigSchema
}

type fsConfigSchema struct {
	Environments []fsConfigSchema_Environment `yaml:"environments"`
}

type fsConfigSchema_Environment struct {
	URL    string `yaml:"url"`
	CACert string `yaml:"ca_cert,omitempty"`

	Alias string `yaml:"alias,omitempty"`

	// Auth
	Username        string `yaml:"username,omitempty"`
	Password        string `yaml:"password,omitempty"`
	AccessTokenType string `yaml:"access_token_type,omitempty"`
	AccessToken     string `yaml:"access_token,omitempty"`
	RefreshToken    string `yaml:"refresh_token,omitempty"`
}

func NewFSConfigFromPath(path string, fs boshsys.FileSystem) (FSConfig, error) {
	var schema fsConfigSchema

	absPath, err := fs.ExpandPath(path)
	if err != nil {
		return FSConfig{}, err
	}

	if fs.FileExists(absPath) {
		bytes, err := fs.ReadFile(absPath)
		if err != nil {
			return FSConfig{}, bosherr.WrapErrorf(err, "Reading config '%s'", absPath)
		}

		err = yaml.Unmarshal(bytes, &schema)
		if err != nil {
			return FSConfig{}, bosherr.WrapError(err, "Unmarshalling config")
		}
	}

	return FSConfig{path: absPath, fs: fs, schema: schema}, nil
}

func (c FSConfig) Environments() []Environment {
	environments := []Environment{}

	for _, tg := range c.schema.Environments {
		environments = append(environments, Environment{URL: tg.URL, Alias: tg.Alias})
	}

	return environments
}

func (c FSConfig) ResolveEnvironment(urlOrAlias string) string {
	_, tg := c.findOrCreateEnvironment(urlOrAlias)

	return tg.URL
}

func (c FSConfig) AliasEnvironment(url, alias, caCert string) (Config, error) {
	if len(url) == 0 {
		return nil, bosherr.Error("Expected non-empty environment URL")
	}

	if len(alias) == 0 {
		return nil, bosherr.Error("Expected non-empty environment alias")
	}

	config := c.deepCopy()

	i, tg := config.findOrCreateEnvironmentByUrlOrAlias(url, alias)
	tg.URL = url
	tg.Alias = alias
	tg.CACert = caCert
	config.schema.Environments[i] = tg

	return config, nil
}

func (c FSConfig) UnaliasEnvironment(alias string) (Config, error) {
	if len(alias) == 0 {
		return nil, bosherr.Error("expected non-empty environment alias")
	}

	idx := -1
	for i, tg := range c.schema.Environments {
		if alias == tg.Alias {
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil, bosherr.Errorf("alias %s not found", alias)
	}
	config := c.deepCopy()
	config.schema.Environments = append(c.schema.Environments[:idx], c.schema.Environments[idx+1:]...)
	return config, nil
}

func (c FSConfig) CACert(urlOrAlias string) string {
	_, tg := c.findOrCreateEnvironment(urlOrAlias)

	return tg.CACert
}

func (c FSConfig) Credentials(urlOrAlias string) Creds {
	_, tg := c.findOrCreateEnvironment(urlOrAlias)

	return Creds{
		Client:       tg.Username,
		ClientSecret: tg.Password,

		AccessTokenType: tg.AccessTokenType,
		AccessToken:     tg.AccessToken,
		RefreshToken:    tg.RefreshToken,
	}
}

func (c FSConfig) SetCredentials(urlOrAlias string, creds Creds) Config {
	config := c.deepCopy()

	i, tg := config.findOrCreateEnvironment(urlOrAlias)
	tg.Username = creds.Client
	tg.Password = creds.ClientSecret
	tg.AccessTokenType = creds.AccessTokenType
	tg.AccessToken = creds.AccessToken
	tg.RefreshToken = creds.RefreshToken
	config.schema.Environments[i] = tg

	return config
}

func (c FSConfig) UnsetCredentials(urlOrAlias string) Config {
	config := c.deepCopy()

	i, tg := config.findOrCreateEnvironment(urlOrAlias)
	tg.Username = ""
	tg.Password = ""
	tg.AccessTokenType = ""
	tg.AccessToken = ""
	tg.RefreshToken = ""
	config.schema.Environments[i] = tg

	return config
}

func (c FSConfig) Save() error {
	bytes, err := yaml.Marshal(c.schema)
	if err != nil {
		return bosherr.WrapError(err, "Marshalling config")
	}

	err = c.fs.WriteFile(c.path, bytes)
	if err != nil {
		return bosherr.WrapErrorf(err, "Writing config '%s'", c.path)
	}
	err = c.fs.Chmod(c.path, os.FileMode(0600))
	if err != nil {
		return bosherr.WrapErrorf(err, "Setting config '%s' permissions", c.path)
	}

	return nil
}

func (c FSConfig) UpdateConfigWithToken(environment string, t uaa.AccessToken) error {
	creds := Creds{
		AccessToken:     t.Value(),
		AccessTokenType: t.Type(),
	}

	if refreshToken, ok := t.(uaa.RefreshableAccessToken); ok {
		creds.RefreshToken = refreshToken.RefreshValue()
	}
	config := c.SetCredentials(environment, creds)
	return config.Save()
}

func (c *FSConfig) findOrCreateEnvironment(urlOrAlias string) (int, fsConfigSchema_Environment) {
	// Always consider empty URL/alias as a new item
	if urlOrAlias != "" {
		for i, tg := range c.schema.Environments {
			if urlOrAlias == tg.URL || urlOrAlias == tg.Alias {
				return i, tg
			}
		}
	}

	return c.appendNewEnvironmentWithURL(urlOrAlias)
}

func (c *FSConfig) findOrCreateEnvironmentByUrlOrAlias(url, alias string) (int, fsConfigSchema_Environment) {
	for i, tg := range c.schema.Environments {
		if url == tg.URL || alias == tg.Alias {
			return i, tg
		}
	}

	i, tg := c.appendNewEnvironmentWithURL(url)
	tg.Alias = alias
	return i, tg
}

func (c *FSConfig) appendNewEnvironmentWithURL(url string) (int, fsConfigSchema_Environment) {
	tg := fsConfigSchema_Environment{URL: url}
	c.schema.Environments = append(c.schema.Environments, tg)
	return len(c.schema.Environments) - 1, tg
}

func (c FSConfig) deepCopy() FSConfig {
	bytes, err := yaml.Marshal(c.schema)
	if err != nil {
		panic("serializing config schema")
	}

	var schema fsConfigSchema

	err = yaml.Unmarshal(bytes, &schema)
	if err != nil {
		panic("deserializing config schema")
	}

	return FSConfig{path: c.path, fs: c.fs, schema: schema}
}
