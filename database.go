package app

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/stdlib"
	//_ "github.com/lib/pq"
)

func DB() *sql.DB {
	cfg := Config()

	db := stdlib.OpenDB(pgx.ConnConfig{
		Config: pgconn.Config{
			Host:     cfg.Database.Host,
			Port:     cfg.Database.Port,
			Database: cfg.Database.Name,
			User:     cfg.Database.User,
			Password: cfg.Database.Pass,
			TLSConfig: &tls.Config{
				Rand:                        nil,
				Time:                        nil,
				Certificates:                nil,
				NameToCertificate:           nil,
				GetCertificate:              nil,
				GetClientCertificate:        nil,
				GetConfigForClient:          nil,
				VerifyPeerCertificate:       nil,
				VerifyConnection:            nil,
				RootCAs:                     &x509.CertPool{},
				NextProtos:                  nil,
				ServerName:                  "",
				ClientAuth:                  0,
				ClientCAs:                   &x509.CertPool{},
				InsecureSkipVerify:          false,
				CipherSuites:                nil,
				PreferServerCipherSuites:    false,
				SessionTicketsDisabled:      false,
				SessionTicketKey:            [32]byte{},
				ClientSessionCache:          nil,
				MinVersion:                  0,
				MaxVersion:                  0,
				CurvePreferences:            nil,
				DynamicRecordSizingDisabled: false,
				Renegotiation:               0,
				KeyLogWriter:                nil,
			},
			ConnectTimeout:  0,
			DialFunc:        nil,
			LookupFunc:      nil,
			BuildFrontend:   nil,
			RuntimeParams:   nil,
			KerberosSrvName: "",
			KerberosSpn:     "",
			Fallbacks:       nil,
			ValidateConnect: nil,
			AfterConnect:    nil,
			OnNotice:        nil,
			OnNotification:  nil,
		},
		Tracer:                   nil,
		StatementCacheCapacity:   0,
		DescriptionCacheCapacity: 0,
		DefaultQueryExecMode:     0,
	})

	//dsn := fmt.Sprintf("postgres://%s:%d/%s?sslmode=%s&user=%s&password=%s",
	//	cfg.Database.Host,
	//	cfg.Database.Port,
	//	cfg.Database.Name,
	//	cfg.Database.SslMode,
	//	cfg.Database.User,
	//	cfg.Database.Pass,
	//)

	//db, err := sql.Open("postgres", dsn)
	//if err != nil {
	//	_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}

	_, err := db.Exec("SELECT true")
	if err != nil {
		log.Panic(err)
	}

	return db
}
