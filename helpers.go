package mikrotik

import (
	"crypto/tls"
	"time"

	routeros "gopkg.in/routeros.v2"
)

//Dial to mikrotik router
func Dial(addr, user, pass string) (*Mikrotik, error) {
	c, err := routeros.Dial(addr, user, pass)
	if err != nil {
		return nil, err
	}

	mik := &Mikrotik{Conn: c}
	mik.setMikrotikCommands()

	return mik, nil
}

//DialTimeout dial to mikrotik router with timeout
func DialTimeout(addr, user, pass string, timeout time.Duration) (*Mikrotik, error) {
	c, err := routeros.DialTimeout(addr, user, pass, timeout)
	if err != nil {
		return nil, err
	}

	mik := &Mikrotik{Conn: c}
	mik.setMikrotikCommands()

	return mik, nil
}

//DialTLSTimeout dial to mikrotik router with timeout and TLS
func DialTLSTimeout(addr, user, pass string, tlsConfig *tls.Config, timeout time.Duration) (*Mikrotik, error) {
	c, err := routeros.DialTLSTimeout(addr, user, pass, tlsConfig, timeout)
	if err != nil {
		return nil, err
	}

	mik := &Mikrotik{Conn: c}
	mik.setMikrotikCommands()

	return mik, nil
}

//DialTLS to mikrotik router with TLS
func DialTLS(addr, user, pass string, tlsConfig *tls.Config) (*Mikrotik, error) {
	c, err := routeros.DialTLS(addr, user, pass, tlsConfig)
	if err != nil {
		return nil, err
	}

	mik := &Mikrotik{Conn: c}
	mik.setMikrotikCommands()

	return mik, nil
}
