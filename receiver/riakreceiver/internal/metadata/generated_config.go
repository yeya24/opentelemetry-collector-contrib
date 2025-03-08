// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/filter"
)

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for riak metrics.
type MetricsConfig struct {
	RiakMemoryLimit              MetricConfig `mapstructure:"riak.memory.limit"`
	RiakNodeOperationCount       MetricConfig `mapstructure:"riak.node.operation.count"`
	RiakNodeOperationTimeMean    MetricConfig `mapstructure:"riak.node.operation.time.mean"`
	RiakNodeReadRepairCount      MetricConfig `mapstructure:"riak.node.read_repair.count"`
	RiakVnodeIndexOperationCount MetricConfig `mapstructure:"riak.vnode.index.operation.count"`
	RiakVnodeOperationCount      MetricConfig `mapstructure:"riak.vnode.operation.count"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		RiakMemoryLimit: MetricConfig{
			Enabled: true,
		},
		RiakNodeOperationCount: MetricConfig{
			Enabled: true,
		},
		RiakNodeOperationTimeMean: MetricConfig{
			Enabled: true,
		},
		RiakNodeReadRepairCount: MetricConfig{
			Enabled: true,
		},
		RiakVnodeIndexOperationCount: MetricConfig{
			Enabled: true,
		},
		RiakVnodeOperationCount: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
	// Experimental: MetricsInclude defines a list of filters for attribute values.
	// If the list is not empty, only metrics with matching resource attribute values will be emitted.
	MetricsInclude []filter.Config `mapstructure:"metrics_include"`
	// Experimental: MetricsExclude defines a list of filters for attribute values.
	// If the list is not empty, metrics with matching resource attribute values will not be emitted.
	// MetricsInclude has higher priority than MetricsExclude.
	MetricsExclude []filter.Config `mapstructure:"metrics_exclude"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for riak resource attributes.
type ResourceAttributesConfig struct {
	RiakNodeName ResourceAttributeConfig `mapstructure:"riak.node.name"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		RiakNodeName: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for riak metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
