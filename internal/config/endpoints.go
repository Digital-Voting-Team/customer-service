package config

import (
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type RawEndpointsConfig struct {
	Docker    bool                `fig:"docker"`
	Endpoints []rawEndpointConfig `fig:"services"`
}

type rawEndpointConfig struct {
	Endpoint    string   `fig:"service"`
	EntryPoints []string `fig:"entry_points"`
}

type EndpointsConfig struct {
	Endpoints map[string]string
}

type EndpointsConfiger interface {
	EndpointsConfig() *EndpointsConfig
}

func NewEndpointConfiger(getter kv.Getter) EndpointsConfiger {
	return &endpointsConfig{
		getter: getter,
	}
}

type endpointsConfig struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *endpointsConfig) EndpointsConfig() *EndpointsConfig {
	return c.once.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "Endpoints")
		config := RawEndpointsConfig{}
		err := figure.Out(&config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}
		resConf := parseRawEndpointsConfig(config)
		return &resConf
	}).(*EndpointsConfig)
}

func parseRawEndpointsConfig(raw RawEndpointsConfig) EndpointsConfig {
	resConf := EndpointsConfig{}
	resConf.Endpoints = make(map[string]string)
	for _, endpoint := range raw.Endpoints {
		if raw.Docker {
			resConf.Endpoints[endpoint.Endpoint] = "http://" + endpoint.EntryPoints[1]
		} else {
			resConf.Endpoints[endpoint.Endpoint] = "http://" + endpoint.EntryPoints[0]
		}
	}

	return resConf
}
