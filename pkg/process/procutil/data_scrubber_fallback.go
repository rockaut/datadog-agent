// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !windows

package procutil

import (
	"strings"
)

var (
	defaultSensitiveWords = []string{
		"*password*", "*passwd*", "*mysql_pwd*",
		"*access_token*", "*auth_token*",
		"*api_key*", "*apikey*",
		"*secret*", "*credentials*", "stripetoken",
	}

	forbiddenSymbolsRegex = "[^a-zA-Z0-9_*]"
)

// stripArguments removes all arguments given a command line.
func (ds *DataScrubber) stripArguments(cmdline []string) []string {
	// We will sometimes see the entire command line come in via the first element -- splitting guarantees removal
	// of arguments in these cases.
	if len(cmdline) > 0 {
		return []string{strings.SplitN(cmdline[0], " ", 2)[0]}
	}
	return cmdline
}
