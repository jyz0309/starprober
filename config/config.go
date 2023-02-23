package config

import "time"

type Config struct {
	GlobalConfig GlobalConfig
	HttpConfig   []*HttpConfig
	TcpConfig    []*TCPConfig
	UDPConfig    []*UDPConfig
	ICMPConfig   []*ICMPConfig
}

type GlobalConfig struct {
	ProbeInterval time.Time
}

type HttpConfig struct {
}

type TCPConfig struct {
}

type ICMPConfig struct {
}

type UDPConfig struct {
}
