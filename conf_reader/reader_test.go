package conf_reader

import (
	"net/url"
	"os"
	"strconv"
	"testing"
)

// "PORT"
// "DBURL"
// "JAEGERURL"
// "SENTRYURL"
// "APPID"
// "APPKEY"
// "KAFKA"

var testCases = map[string]map[string]string{
	"valid": {"PORT": "8000",
		"DBURL":     "http://dbtesturl.com",
		"JAEGERURL": "http://jaegertesturl.com",
		"SENTRYURL": "http://sentryrtesturl.com",
		"APPKEY":    "testkey",
		"APPID":     "testid",
		"KAFKA":     "9000"},

	"notValidUrl": {"PORT": "8000",
		"DBURL":     "notValidUrl",
		"JAEGERURL": "notValidUrl",
		"SENTRYURL": "notValidUrl",
		"APPKEY":    "testkey",
		"APPID":     "testid",
		"KAFKA":     "9000"},
}

var defaults = map[string]string{
	"PORT":      "8080",
	"DBURL":     "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable",
	"JAEGERURL": "http://jaeger:16686",
	"SENTRYURL": "http://sentry:9000",
	"KAFKA":     "9090",
	"APPKEY":    "testKey",
	"APPID":     "testID",
}

func prepare(key string) {
	for k, v := range testCases[key] {
		os.Setenv(k, v)
	}
}

func clear() {
	os.Clearenv()
}

func TestValid(t *testing.T) {
	prepare("valid")
	c := ReadEnv()
	if p, _ := strconv.Atoi(testCases["valid"]["PORT"]); *c.Port != p {
		t.Fatalf("Port = %d, but should be %s", *c.Port, testCases["valid"]["PORT"])
	}
	if *c.AppID != testCases["valid"]["APPID"] {
		t.Fatalf("AppID = %s, but should be %s", *c.AppID, defaults["APPID"])
	}
	if *c.AppKey != testCases["valid"]["APPKEY"] {
		t.Fatalf("AppKey = %s, but should be %s", *c.AppKey, testCases["valid"]["APPKEY"])
	}
	if urlT, _ := url.Parse(testCases["valid"]["DBURL"]); *c.Dburl != *urlT {
		t.Fatalf("DB_url = %#v, but should be %#v", *c.Dburl, urlT)
	}
	if urlT, _ := url.Parse(testCases["valid"]["JAEGERURL"]); *c.JaegerUrl != *urlT {
		t.Fatalf("Jaeger_url = %#v, but should be %#v", *c.JaegerUrl, urlT)
	}
	if urlT, _ := url.Parse(testCases["valid"]["SENTRYURL"]); *c.SentryUrl != *urlT {
		t.Fatalf("sENTRY_url = %#v, but should be %#v", *c.SentryUrl, urlT)
	}
	clear()
}

func TestNotValidUrl(t *testing.T) {
	prepare("notValidUrl")
	c := ReadEnv()
	if urlT, _ := url.Parse(defaults["JAEGERURL"]); *c.JaegerUrl != *urlT {
		t.Fatalf("Jaeger_url = %#v, but should be %#v", *c.JaegerUrl, urlT)
	}
	if urlT, _ := url.Parse(defaults["DBURL"]); (*c.Dburl).Scheme != urlT.Scheme || (*c.Dburl).Host != urlT.Host {
		t.Fatalf("DB_url = %#v, but should be %#v", (*c.Dburl).Host, urlT.Host) // тут странно парсит, визуально они одинаковые
	}
	if urlT, _ := url.Parse(defaults["SENTRYURL"]); *c.SentryUrl != *urlT {
		t.Fatalf("sENTRY_url = %#v, but should be %#v", *c.SentryUrl, urlT)
	}
	clear()
}

func TestDefaults(t *testing.T) {
	c := ReadEnv()
	if p, _ := strconv.Atoi(defaults["PORT"]); *c.Port != p {
		t.Fatalf("Port = %d, but should be %s", *c.Port, defaults["PORT"])
	}
	if *c.AppID != defaults["APPID"] {
		t.Fatalf("AppID = %s, but should be %s", *c.AppID, defaults["APPID"])
	}
	if *c.AppKey != defaults["APPKEY"] {
		t.Fatalf("AppKey = %s, but should be %s", *c.AppKey, defaults["APPKEY"])
	}
	if urlT, _ := url.Parse(defaults["DBURL"]); (*c.Dburl).Host != urlT.Host || (*c.Dburl).Scheme != urlT.Scheme {
		t.Fatalf("DB_url = %#v, but should be %#v", *c.Dburl, urlT)
	}
	if urlT, _ := url.Parse(defaults["JAEGERURL"]); *c.JaegerUrl != *urlT {
		t.Fatalf("Jaeger_url = %#v, but should be %#v", *c.JaegerUrl, urlT)
	}
	if urlT, _ := url.Parse(defaults["SENTRYURL"]); *c.SentryUrl != *urlT {
		t.Fatalf("sENTRY_url = %#v, but should be %#v", *c.SentryUrl, urlT)
	}
}
