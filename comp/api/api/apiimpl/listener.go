// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package apiimpl

import (
	"fmt"
	"net"
	"strconv"

	pkgconfigsetup "github.com/DataDog/datadog-agent/pkg/config/setup"
)

// getIPCAddressPort returns a listening connection
func getIPCAddressPort() (string, error) {
	address, err := pkgconfigsetup.GetIPCAddress(pkgconfigsetup.Datadog())
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v:%v", address, pkgconfigsetup.Datadog().GetInt("cmd_port")), nil
}

// getListener returns a listening connection
func getListener(address string) (net.Listener, error) {
	return net.Listen("tcp", address)
}

// returns whether the IPC server is enabled, and if so its host and host:port
func getIPCServerAddressPort() (string, string, bool) {
	ipcServerPort := pkgconfigsetup.Datadog().GetInt("agent_ipc.port")
	if ipcServerPort == 0 {
		return "", "", false
	}

	ipcServerHost := pkgconfigsetup.Datadog().GetString("agent_ipc.host")
	ipcServerHostPort := net.JoinHostPort(ipcServerHost, strconv.Itoa(ipcServerPort))

	return ipcServerHost, ipcServerHostPort, true
}
