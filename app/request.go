package app

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

var errUnsupportedProxyScheme = errors.New("unsupported proxy scheme")

func makeRequest(
	ctx context.Context,
	method string,
	requestURL *url.URL,
	headers http.Header,
	body io.Reader,
	regOpts *registryOptions,
) (*http.Response, error) {
	if requestURL.Scheme != "http" && regOpts != nil && regOpts.Insecure {
		requestURL.Scheme = "http"
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		requestURL.String(),
		body,
	)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		req.Header = headers
	}

	if regOpts != nil {
		if regOpts.Token != "" {
			req.Header.Set("Authorization", "Bearer "+regOpts.Token)
		} else if regOpts.Username != "" && regOpts.Password != "" {
			req.SetBasicAuth(regOpts.Username, regOpts.Password)
		}
	}

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:129.0) Gecko/20100101 Firefox/129.0",
	)

	if s := req.Header.Get("Content-Length"); s != "" {
		contentLength, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}

		req.ContentLength = contentLength
	}

	client, err := newHTTPClient()
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

func newHTTPClient() (*http.Client, error) {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,

		// Keep the normal HTTP transport settings.
		ForceAttemptHTTP2: true,

		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,

		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// If a SOCKS5 proxy is explicitly configured, use it.
	//
	// SOCKS5_PROXY takes precedence over the normal HTTP proxy
	// environment variables.
	socksProxy := getEnv(
		"SOCKS5_PROXY",
		"socks5_proxy",
	)

	if socksProxy != "" {
		dialer, err := newSOCKS5Dialer(socksProxy)
		if err != nil {
			return nil, err
		}

		transport.Proxy = nil
		transport.DialContext = func(
			ctx context.Context,
			network string,
			address string,
		) (net.Conn, error) {

			return dialer.Dial(network, address)
		}
	}

	return &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second,
	}, nil
}

func newSOCKS5Dialer(proxyURL string) (proxy.Dialer, error) {
	u, err := url.Parse(proxyURL)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "socks5" && u.Scheme != "socks5h" {
		return nil, &url.Error{
			Op:  "parse proxy",
			URL: proxyURL,
			Err: errUnsupportedProxyScheme,
		}
	}

	var auth *proxy.Auth

	if u.User != nil {
		auth = &proxy.Auth{
			User: u.User.Username(),
		}

		if password, ok := u.User.Password(); ok {
			auth.Password = password
		}
	}

	return proxy.SOCKS5(
		"tcp",
		u.Host,
		auth,
		proxy.Direct,
	)
}

func getEnv(names ...string) string {
	for _, name := range names {
		if value := strings.TrimSpace(os.Getenv(name)); value != "" {
			return value
		}
	}

	return ""
}
