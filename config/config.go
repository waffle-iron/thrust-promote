package config

import (
    "log"
    "strconv"
    "net"
    "net/url"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type ConnectionSettings struct {
        Pool int `yaml:pool`
        Url string `yaml:url`
        Timeout int `yaml:timeout`
        Host string `yaml:host`
        Port int `yaml:port`
        Database string `yaml:database`
        User string `yaml:user`
        Password string `yaml:password`
}

func (connSettings *ConnectionSettings) ParseUrl() {
    info, err := url.Parse(connSettings.Url)
    if err != nil {
        log.Fatalf("Could not parse url: %v", err)
    }

    host, port, err := net.SplitHostPort(info.Host)
    if err != nil {
        log.Fatalf("Could not parse host: %v", err)
    }
    connSettings.Host = host
    if info.User != nil {
        connSettings.User = info.User.Username()
        password, _ := info.User.Password()
        if password != "" {
            connSettings.Password = password
        }
    }

    portInt, err := strconv.Atoi(port)
    if err != nil {
        log.Fatalf("Invalid Port: %v", err)
    }
    connSettings.Port = portInt

    if info.Path != "" {
        connSettings.Database = info.Path[1:len(info.Path)]
    }

}

type Config struct {
    Db struct {
        Development ConnectionSettings `yaml:development`
        Production ConnectionSettings `yaml:production`
    } `yaml:db`
    Redis struct {
        Development ConnectionSettings `yaml:development`
        Production ConnectionSettings `yaml:production`
    } `yaml:redis`
    Instagram struct {
        ClientId string `yaml:client_id`
        ClientSecret string `yaml:client_secret`
    } `yaml:instagram`
    Soundcloud struct {
        ClientId string `yaml:client_id`
        ClientSecret string `yaml:client_secret`
    } `yaml:soundcloud`
    Youtube struct {
        ClientId string `yaml:client_id`
        ClientSecret string `yaml:client_secret`
        AccessToken string `yaml:access_token`
        RefreshToken string `yaml:refresh_token`
    } `yaml:youtube`
    Dropbox struct {
        AccessToken string `yaml:access_token`
        ImageFolder string `yaml:image_folder`
    } `yaml:dropbox`
    Twitter struct {
        ConsumerKey string `yaml:consumer_key`
        ConsumerSecret string `yaml:consumer_secret`
    } `yaml:twitter`
}

func LoadConfig(configPath string) *Config {
    data, err := ioutil.ReadFile(configPath)
    if err != nil {
        log.Fatalf("Unable to load file: %v", err)
        log.Panic(err)
    }

    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        log.Fatalf("Unable to parse file: %v", err)
        log.Panic(err)
    }

    return &config
}