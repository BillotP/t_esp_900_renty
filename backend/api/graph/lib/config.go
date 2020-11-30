package lib

import (
	"os"
	"time"
)

// Config represent the graphql_server configuration variables
type Config struct {
	Stage                  string `env:"STAGE"`
	Release                string `env:"RELEASE"`
	Port                   string `env:"PORT"`
	Host                   string `env:"HOST"`
	URL                    string `env:"URL"`
	JwtIssuer              string `env:"JWT_ISSUER"`
	JwtSigningKey          string `env:"SIGN_KEY"`
	TotpExpiration         string `env:"TOTP_EXP"`
	SessionExpiration      string `env:"SESSION_EXP"`
	MaxSizeByte            int64
	TotpDuration           time.Duration
	SessionDuration        time.Duration
	SentryDSN              string `env:"SENTRY_DSN"`
}

// ServerConf is the exported config object singleton
var ServerConf = Config{}

func init() {
	ServerConf.Load()
}

// Load config options from environnement variables
func (c *Config) Load() {
	var err error
	c.Stage = GetDefVal("STAGE", "dev")
	c.Release = GetDefVal("RELEASE", "v0.1.11")
	c.Port = GetDefVal("PORT", "4000")
	c.Host = GetDefVal("HOST", "localhost")
	c.URL = GetDefVal("URL", "http://localhost:4000")
	c.JwtIssuer = GetDefVal("JWT_ISSUER", "renty-dev")
	c.JwtSigningKey = GetDefVal("SIGN_KEY", "SUPERSECRETTOBECHANGED")
	c.TotpExpiration = GetDefVal("TOTP_EXP", "15m")
	c.SessionExpiration = GetDefVal("SESSION_EXP", "12h")
	if c.TotpDuration, err = time.ParseDuration(c.TotpExpiration); err != nil {
		LogError("lib/Config", err.Error())
		os.Exit(1)
	}
	if c.SessionDuration, err = time.ParseDuration(c.SessionExpiration); err != nil {
		LogError("lib/Config", err.Error())
		os.Exit(1)
	}
	c.SentryDSN = GetDefVal("SENTRY_DSN", "https://67f58ee98e4e4b76a50ffc280a8eaaf1@sentry.io/5176087")
}
