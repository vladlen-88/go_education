package conf_reader

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

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
		// log.Fatal(err)
		return nil, err
	}
	return tUrl, nil
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

func NewConfig() *Config {
	// port := 8080
	return &Config{
		Port:            new(int),
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

func ReadFlag() *Config {
	flagConfig := NewFlagConfig()
	flag.Parse()
	flagConfig.Validate(dburl, jaeger_url, sentry_url)
	return flagConfig

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
		// log.Println("Unable to evaluate kafka broker port, we are going to use default")
		c.KafkaBrokerPort = &kafka
	}
	c.KafkaBrokerPort = &kafkaT
	return c
}

func (c *Config) UnmarshalJSON(j []byte) error {
	var rawData map[string]interface{}
	err := json.Unmarshal(j, &rawData)
	if err != nil {
		return err
	}
	for k, v := range rawData {
		switch strings.ToLower(k) {
		case "port":
			r, ok := v.(int)
			if !ok {
				c.Port = &port
			} else {
				c.Port = &r
			}
		case "db_url":
			c.Dburl = useValidOrDefault(fmt.Sprintf("%v", v), dburl)
		case "jaeger_url":
			c.JaegerUrl = useValidOrDefault(fmt.Sprintf("%v", v), jaeger_url)
		case "sentry_url":
			c.SentryUrl = useValidOrDefault(fmt.Sprintf("%v", v), sentry_url)
		case "kafka_broker":
			r, ok := v.(int)
			if !ok {
				c.KafkaBrokerPort = &kafka
			} else {
				c.KafkaBrokerPort = &r
			}
		case "app_id":
			r := fmt.Sprintf("%v", v)
			if r == "" {
				c.AppID = &appID
			} else {
				c.AppID = &r
			}
		case "app_key":
			r := fmt.Sprintf("%v", v)
			if r == "" {
				c.AppKey = &appKey
			} else {
				c.AppKey = &r
			}
		}
	}
	return nil
}

func ReadFromFile() *Config {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(path)
	file, err := os.Open(path + "/conf_reader/config.json")
	if err != nil {
		log.Fatalf("can't open the file %v", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("can't close the file %v", err)
		}
	}()
	c := NewConfig()
	info, err := file.Stat()
	if err != nil {
		log.Fatalf("can't get info about file %v", err)
	}
	data := make([]byte, info.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatalf("can't read from file %v", err)
	}
	err = c.UnmarshalJSON(data)
	if err != nil {
		log.Printf("can't unmarhall %v", err)
	}
	return c

}
