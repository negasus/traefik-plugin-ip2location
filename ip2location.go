package traefik_plugin_ip2location

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

type Headers struct {
	CountryShort string `json:"country_short,omitempty"`
}

// Config the plugin configuration.
type Config struct {
	Filename   string  `json:"filename,omitempty"`
	FromHeader string  `json:"from_header,omitempty"`
	Headers    Headers `json:"headers,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// GeoIP2 a GeoIP2 plugin.
type GeoIP2 struct {
	next       http.Handler
	name       string
	fromHeader string
	db         *DB
	headers    Headers
}

// New created a new ip2location plugin.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {

	db, err := OpenDB(config.Filename)
	if err != nil {
		return nil, fmt.Errorf("error open database file, %w", err)
	}

	return &GeoIP2{
		next:       next,
		name:       name,
		fromHeader: config.FromHeader,
		db:         db,
		headers:    config.Headers,
	}, nil
}

func (a *GeoIP2) getIP(req *http.Request) (net.IP, error) {
	if a.fromHeader != "" {
		ip := net.ParseIP(req.Header.Get(a.fromHeader))
		return ip, nil
	}

	addr, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, err
	}

	ip := net.ParseIP(addr)
	return ip, nil
}

func (a *GeoIP2) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	//ip, err := a.getIP(req)
	//if err != nil {
	//	req.Header.Add("X-IP2LOCATION-ERROR", err.Error())
	//	a.next.ServeHTTP(rw, req)
	//}

	record, err := a.db.Get_all("4.0.0.0")
	//record, err := a.db.Get_all(ip.String())
	//if err != nil {
	//	req.Header.Add("X-IP2LOCATION-ERROR", err.Error())
	//	a.next.ServeHTTP(rw, req)
	//}
	//
	////log.Printf("[[]] " + a.headers.CountryShort)
	//
	//if a.headers.CountryShort != "" {
	//	req.Header.Add(a.headers.CountryShort, record.Country_short)
	//}
	//req.Header.Add("XXX", record.Country_short)
	_, _ = record, err
	//req.Header.Add("XXX", ip.String())

	a.next.ServeHTTP(rw, req)
}
