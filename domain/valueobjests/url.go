package valueobjests

import (
	"errors"
	"net"
	"net/url"
	"strconv"
)

type Url struct {
	scheme   string
	hostname string
	ip       []net.IP
	port     int
	endpoint string
	query    map[string][]string
}

func NewUrl(value string) (*Url, error) {

	URL := &Url{}
	parsedUrl, err := url.Parse(value)
	if err != nil {
		return nil, errors.New("Invalid URl: " + value)
	}

	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return nil, errors.New("Invalid Schema: " + parsedUrl.Scheme + " .Supported Schema: http, https")
	}
	URL.scheme = parsedUrl.Scheme

	if parsedUrl.Port() == "" {
		switch parsedUrl.Scheme {
		case "http":
			URL.port = 80
		case "https":
			URL.port = 443
		}
	} else if parsedUrl.Port() > "65536" {
		return nil, errors.New("Invalid Port: " + parsedUrl.Port() + " .Supported Values: 1-65536")
	} else {
		URL.port, _ = strconv.Atoi(parsedUrl.Port())
	}

	host := parsedUrl.Hostname()
	if ip := net.ParseIP(host); ip != nil {
		URL.ip = []net.IP{ip}
	} else {
		URL.hostname = host
		ipList, err := net.LookupIP(host)
		if err != nil {
			return nil, errors.New("Unable to resolve DNS of host: " + host)
		} else {
			URL.ip = ipList
		}
	}

	URL.endpoint = parsedUrl.EscapedPath()
	URL.query = parsedUrl.Query()

	return URL, nil

}
