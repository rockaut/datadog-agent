// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package net

import (
	"time"

	model "github.com/DataDog/agent-payload/v5/process"

	nppayload "github.com/DataDog/datadog-agent/pkg/networkpath/payload"
)

// SysProbeUtilGetter is a function that returns a SysProbeUtil for the given path
// The standard implementation is GetRemoteSysProbeUtil
type SysProbeUtilGetter func(string) (SysProbeUtil, error)

// SysProbeUtil fetches info from the SysProbe running remotely
type SysProbeUtil interface {
	GetConnections(clientID string) (*model.Connections, error)
	GetProcStats(pids []int32) (*model.ProcStatsWithPermByPID, error)
	Register(clientID string) error
	GetNetworkID() (string, error)
	GetPing(clientID string, host string, count int, interval time.Duration, timeout time.Duration) ([]byte, error)
	GetTraceroute(clientID string, host string, port uint16, protocol nppayload.Protocol, maxTTL uint8, timeout time.Duration) ([]byte, error)
}
