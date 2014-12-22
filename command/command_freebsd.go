package command

import (
	"github.com/mackerelio/mackerel-agent/config"
	"github.com/mackerelio/mackerel-agent/metrics"
	metricsFreebsd "github.com/mackerelio/mackerel-agent/metrics/freebsd"
	"github.com/mackerelio/mackerel-agent/spec"
	specFreebsd "github.com/mackerelio/mackerel-agent/spec/freebsd"
)

func specGenerators() []spec.Generator {
	return []spec.Generator{
		&specFreebsd.KernelGenerator{},
		&specFreebsd.MemoryGenerator{},
		&specFreebsd.CPUGenerator{},
		&specFreebsd.FilesystemGenerator{},
	}
}

func interfaceGenerator() spec.Generator {
	return &specFreebsd.InterfaceGenerator{}
}

func metricsGenerators(conf *config.Config) []metrics.Generator {
	generators := []metrics.Generator{
		&metricsFreebsd.Loadavg5Generator{},
		&metricsFreebsd.CpuusageGenerator{},
		&metricsFreebsd.FilesystemGenerator{},
	}

	return generators
}

func pluginGenerators(conf *config.Config) []metrics.PluginGenerator {
	generators := []metrics.PluginGenerator{}

	for _, pluginConfig := range conf.Plugin["metrics"] {
		generators = append(generators, metrics.NewPluginGenerator(pluginConfig))
	}

	return generators
}
