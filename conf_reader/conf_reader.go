package conf_reader

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func ReadFlag() *Config {
	flagConfig := NewFlagConfig()
	flag.Parse()
	flagConfig.Validate(dburl, jaeger_url, sentry_url)
	return flagConfig

}

func NewFlagConfig() *Config {
	return &Config{
		Port:            &port,
		JaegerUrl:       &url.URL{},
		SentryUrl:       &url.URL{},
		KafkaBrokerPort: &kafka,
		AppID:           &appID,
		AppKey:          &appKey,
	}
}

func init() {
	flag.IntVar(&port, "port", port, "Set port, 8080 if doen't set")
	flag.StringVar(&dburl, "dburl", dburl, "Set if you want to chenge db url")
	flag.StringVar(&jaeger_url, "jaeger", jaeger_url, "set if you whant to change jaeger url")
	flag.StringVar(&sentry_url, "sentry", sentry_url, "set if you whant to change sentry url")
	flag.IntVar(&kafka, "kafka", kafka, "Set to change kafka port if it's different from 9092")
	flag.StringVar(&appID, "appid", appID, "change id, default - testid")
	flag.StringVar(&appKey, "appkey", appKey, "appkey, default testkey")
}

var (
	port       = 8080
	dburl      = "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"
	jaeger_url = "http://jaeger:16686"
	sentry_url = "http://sentry:9000"
	kafka      = 9090
	appKey     = "testKey"
	appID      = "testID"
)

type Config struct {
	Port            *int
	Dburl           *url.URL
	JaegerUrl       *url.URL
	SentryUrl       *url.URL
	KafkaBrokerPort *int
	AppID           *string
	AppKey          *string
}

func validate(rawUrl string) (*url.URL, error) {
	tUrl, err := url.Parse(rawUrl)
	if err != nil || tUrl.Scheme == "" || tUrl.Host == "" {
		err = fmt.Errorf("%s is not valid url", rawUrl)
		return nil, err
	}
	return tUrl, nil
}

func NewConfig() *Config {
	port := 8080
	return &Config{
		Port:            &port,
		Dburl:           &url.URL{},
		JaegerUrl:       &url.URL{},
		SentryUrl:       &url.URL{},
		KafkaBrokerPort: new(int),
		AppID:           new(string),
		AppKey:          new(string),
	}
}

func useValidOrDefault(u, defaultValue string) *url.URL {
	tUrl, err := validate(u)
	if err != nil {
		tUrl, _ = url.Parse(defaultValue)
	}
	return tUrl
}

func (c *Config) Validate(db, jaeger, sentry string) {
	c.Dburl = useValidOrDefault(db, dburl)
	c.JaegerUrl = useValidOrDefault(jaeger, jaeger_url)
	c.SentryUrl = useValidOrDefault(sentry, sentry_url)
}

func ReadEnv() *Config {
	c := NewConfig()
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		c.Port = &port
	}
	c.Validate(os.Getenv("DBURL"), os.Getenv("JAEGERURL"), os.Getenv("SENTRYURL"))
	a := os.Getenv("APPID")
	if a == "" {
		c.AppID = &appID
	} else {
		c.AppID = &a
	}
	aK := os.Getenv("APPKEY")
	if aK == "" {
		c.AppKey = &appKey
	} else {
		c.AppKey = &aK
	}

	kafkaT, err := strconv.Atoi(os.Getenv("KAFKA"))
	if err != nil {
		c.KafkaBrokerPort = &kafka
	}
	c.KafkaBrokerPort = &kafkaT
	return c
}
