package web

import (
	"database/sql"
	"fmt"
	"github.com/mraron/njudge/internal/njudge/email"
	"log"
	"log/slog"
	"net/url"
	"strconv"
	"time"
)

type DatabaseConfig struct {
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
	Host     string `mapstructure:"host" yaml:"host"`
	Name     string `mapstructure:"name" yaml:"name"`
	Port     int    `mapstructure:"port" yaml:"port"`
	SSLMode  bool   `mapstructure:"ssl_mode" yaml:"ssl_mode"`
}

func (db DatabaseConfig) Connect() (*sql.DB, error) {
	SSLMode := "require"
	if !db.SSLMode {
		SSLMode = "disable"
	}

	if db.Port == 0 {
		db.Port = 5432
	}

	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=%s", db.User, db.Password, db.Host, db.Name, db.Port, SSLMode)
	return sql.Open("postgres", connStr)
}

func (db DatabaseConfig) ConnectAndPing(log *slog.Logger) (*sql.DB, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	for {
		log.Info("Trying to ping database...")
		if err := conn.Ping(); err == nil {
			log.Info("OK, connected to database")
			break
		} else {
			log.Error("Failed to connect to database", "error", err)
		}
		time.Sleep(5 * time.Second)
	}
	return conn, nil
}

type Mode string

func (m Mode) Valid() bool {
	return m == ModeDebug || m == ModeDevelopment || m == ModeDemo || m == ModeProduction
}

func (m Mode) UsesDB() bool {
	return m != ModeDemo
}

const (
	ModeDebug       Mode = "debug"
	ModeDevelopment Mode = "development"
	ModeDemo        Mode = "demo"
	ModeProduction  Mode = "production"
)

type GoogleAuthConfig struct {
	Enabled   bool   `mapstructure:"enabled" yaml:"enabled"`
	ClientKey string `mapstructure:"client_key" yaml:"client_key"`
	Secret    string `mapstructure:"secret" yaml:"secret"`
}

type SendgridConfig struct {
	Enabled       bool   `yaml:"enabled" mapstructure:"enabled"`
	ApiKey        string `yaml:"api_key" mapstructure:"api_key"`
	SenderName    string `yaml:"sender_name" mapstructure:"sender_name"`
	SenderAddress string `yaml:"sender_address" mapstructure:"sender_address"`
}

type SMTPConfig struct {
	Enabled             bool   `yaml:"enabled" mapstructure:"enabled"`
	MailAccount         string `yaml:"mail_account" mapstructure:"mail_account"`
	MailServerHost      string `yaml:"mail_server" mapstructure:"mail_server"`
	MailServerPort      int    `yaml:"mail_port" mapstructure:"mail_port"`
	MailAccountPassword string `yaml:"mail_password" mapstructure:"mail_password"`
}

type Config struct {
	Mode     Mode   `mapstructure:"mode" yaml:"mode"`
	Url      string `mapstructure:"url" yaml:"url"`
	Port     string `mapstructure:"port" yaml:"port"`
	TimeZone string `mapstructure:"time_zone" yaml:"time_zone"`

	CookieSecret string `mapstructure:"cookie_secret" yaml:"cookie_secret"`

	GoogleAuth GoogleAuthConfig `mapstructure:"google_auth" yaml:"google_auth"`

	Sendgrid SendgridConfig `yaml:"sendgrid" mapstructure:"sendgrid"`

	SMTP SMTPConfig `yaml:"smtp" mapstructure:"smtp"`

	DatabaseConfig `yaml:"db" mapstructure:"db"`
}

func (s Config) EmailService() email.Service {
	if s.SMTP.Enabled {
		return email.SMTPService{
			From:     s.SMTP.MailAccount,
			Host:     s.SMTP.MailServerHost,
			Port:     s.SMTP.MailServerPort,
			User:     s.SMTP.MailAccount,
			Password: s.SMTP.MailAccountPassword,
		}
	} else if s.Sendgrid.Enabled {
		return email.SendgridService{
			SenderName:    s.Sendgrid.SenderName,
			SenderAddress: s.Sendgrid.SenderAddress,
			APIKey:        s.Sendgrid.ApiKey,
		}
	} else {
		if s.Mode == ModeDevelopment || s.Mode == ModeDebug || s.Mode == ModeDemo {
			return email.LogService{Logger: log.Default()}
		} else {
			return email.ErrorService{}
		}
	}
}

func (s Config) Valid() error {
	if !s.Mode.Valid() {
		return fmt.Errorf("invalid mode: %q", s.Mode)
	}
	u, err := url.Parse(s.Url)
	if err != nil {
		return fmt.Errorf("error parsing url: %w", err)
	} else if u.Scheme == "" {
		return fmt.Errorf("url scheme is required")
	} else if u.Host == "" {
		return fmt.Errorf("url host is required")
	} else if u.Path != "" {
		return fmt.Errorf("no url path is needed")
	}
	p, err := strconv.Atoi(s.Port)
	if err != nil || p <= 0 || p > 65535 {
		return fmt.Errorf("port must be an integer between 1 and 65535")
	}
	if s.CookieSecret == "" {
		return fmt.Errorf("cookiesecret must not be empty")
	}
	if s.GoogleAuth.Enabled {
		if s.GoogleAuth.ClientKey == "" || s.GoogleAuth.Secret == "" {
			return fmt.Errorf("google auth secret and client key is required")
		}
	}
	if s.Sendgrid.Enabled {
		if s.Sendgrid.ApiKey == "" {
			return fmt.Errorf("sendgrid api key is required")
		}
	}
	return nil
}
