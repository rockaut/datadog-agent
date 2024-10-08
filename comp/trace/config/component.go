// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Package config implements a component to handle trace-agent configuration.  This
// component temporarily wraps pkg/trace/config.
//
// This component initializes pkg/config based on the bundle params, and
// will return the same results as that package.  This is to support migration
// to a component architecture.  When no code still uses pkg/config, that
// package will be removed.
//
// The mock component does nothing at startup, beginning with an empty config.
// It also overwrites the pkg/config.SystemProbe for the duration of the test.
package config

import (
	"net/http"

	"go.uber.org/fx"

	"github.com/DataDog/datadog-agent/pkg/config/model"
	traceconfig "github.com/DataDog/datadog-agent/pkg/trace/config"
	"github.com/DataDog/datadog-agent/pkg/util/fxutil"
)

// team: agent-apm

// Component is the component type.
type Component interface {
	// Warnings returns config warnings collected during setup.
	Warnings() *model.Warnings

	// SetHandler returns a handler for runtime configuration changes.
	SetHandler() http.Handler

	// GetConfigHandler returns a handler to fetch the runtime configuration.
	GetConfigHandler() http.Handler

	// SetMaxMemCPU
	SetMaxMemCPU(isContainerized bool)

	// Object returns wrapped config
	Object() *traceconfig.AgentConfig

	// OnUpdateAPIKey registers a callback for API Key changes
	OnUpdateAPIKey(func(oldKey, newKey string))
}

// Module defines the fx options for this component.
func Module() fxutil.Module {
	return fxutil.Component(
		fx.Provide(NewConfig),
		fx.Supply(Params{
			FailIfAPIKeyMissing: true,
		}))
}
