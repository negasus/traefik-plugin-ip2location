package traefik_plugin_ip2location

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
)

type Headers struct {
	CountryShort       string `json:"country_short,omitempty"`
	CountryLong        string `json:"country_long,omitempty"`
	Region             string `json:"region"`
	City               string `json:"city"`
	Isp                string `json:"isp"`
	Latitude           string `json:"latitude"`
	Longitude          string `json:"longitude"`
	Domain             string `json:"domain"`
	Zipcode            string `json:"zipcode"`
	Timezone           string `json:"timezone"`
	Netspeed           string `json:"netspeed"`
	Iddcode            string `json:"iddcode"`
	Areacode           string `json:"areacode"`
	Weatherstationcode string `json:"weatherstationcode"`
	Weatherstationname string `json:"weatherstationname"`
	Mcc                string `json:"mcc"`
	Mnc                string `json:"mnc"`
	Mobilebrand        string `json:"mobilebrand"`
	Elevation          string `json:"elevation"`
	Usagetype          string `json:"usagetype"`
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

// IP2Location plugin.
type IP2Location struct {
	next       http.Handler
	name       string
	fromHeader string
	db         *DB
	headers    Headers
}

// New created a new IP2Location plugin.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	db, err := OpenDB(config.Filename)
	if err != nil {
		return nil, fmt.Errorf("error open database file, %w", err)
	}

	return &IP2Location{
		next:       next,
		name:       name,
		fromHeader: config.FromHeader,
		db:         db,
		headers:    config.Headers,
	}, nil
}

func (a *IP2Location) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ip, err := a.getIP(req)
	if err != nil {
		req.Header.Add("X-IP2LOCATION-ERROR", err.Error())
		a.next.ServeHTTP(rw, req)
	}

	record, err := a.db.Get_all(ip.String())
	if err != nil {
		req.Header.Add("X-IP2LOCATION-ERROR", err.Error())
		a.next.ServeHTTP(rw, req)
	}

	a.addHeaders(req, &record)

	a.next.ServeHTTP(rw, req)
}

func (a *IP2Location) getIP(req *http.Request) (net.IP, error) {
	remoteAddr := req.RemoteAddr

	if a.fromHeader != "" {
		remoteAddr = req.Header.Get(a.fromHeader)
	}

	addr, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return nil, err
	}

	return net.ParseIP(addr), nil
}

func (a *IP2Location) addHeaders(req *http.Request, record *IP2Locationrecord) {
	if a.headers.CountryShort != "" {
		req.Header.Add(a.headers.CountryShort, record.Country_short)
	}
	if a.headers.CountryLong != "" {
		req.Header.Add(a.headers.CountryLong, record.Country_long)
	}
	if a.headers.Region != "" {
		req.Header.Add(a.headers.Region, record.Region)
	}
	if a.headers.City != "" {
		req.Header.Add(a.headers.City, record.City)
	}
	if a.headers.Isp != "" {
		req.Header.Add(a.headers.Isp, record.Isp)
	}
	if a.headers.Latitude != "" {
		req.Header.Add(a.headers.Latitude, strconv.FormatFloat(float64(record.Latitude), 'f', 0, 64))
	}
	if a.headers.Longitude != "" {
		req.Header.Add(a.headers.Longitude, strconv.FormatFloat(float64(record.Longitude), 'f', 0, 64))
	}
	if a.headers.Domain != "" {
		req.Header.Add(a.headers.Domain, record.Domain)
	}
	if a.headers.Zipcode != "" {
		req.Header.Add(a.headers.Zipcode, record.Zipcode)
	}
	if a.headers.Timezone != "" {
		req.Header.Add(a.headers.Timezone, record.Timezone)
	}
	if a.headers.Netspeed != "" {
		req.Header.Add(a.headers.Netspeed, record.Netspeed)
	}
	if a.headers.Iddcode != "" {
		req.Header.Add(a.headers.Iddcode, record.Iddcode)
	}
	if a.headers.Areacode != "" {
		req.Header.Add(a.headers.Areacode, record.Areacode)
	}
	if a.headers.Weatherstationcode != "" {
		req.Header.Add(a.headers.Weatherstationcode, record.Weatherstationcode)
	}
	if a.headers.Weatherstationname != "" {
		req.Header.Add(a.headers.Weatherstationname, record.Weatherstationname)
	}
	if a.headers.Mcc != "" {
		req.Header.Add(a.headers.Mcc, record.Mcc)
	}
	if a.headers.Mnc != "" {
		req.Header.Add(a.headers.Mnc, record.Mnc)
	}
	if a.headers.Mobilebrand != "" {
		req.Header.Add(a.headers.Mobilebrand, record.Mobilebrand)
	}
	if a.headers.Elevation != "" {
		req.Header.Add(a.headers.Elevation, strconv.FormatFloat(float64(record.Elevation), 'f', 0, 64))
	}
	if a.headers.Usagetype != "" {
		req.Header.Add(a.headers.Usagetype, record.Usagetype)
	}
}
