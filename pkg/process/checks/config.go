// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package checks

import (
	"github.com/DataDog/datadog-agent/pkg/config/model"
	pkgconfigsetup "github.com/DataDog/datadog-agent/pkg/config/setup"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

// getMaxBatchSize returns the maximum number of items (processes, containers, process_discoveries) in a check payload
var getMaxBatchSize = func(config model.Reader) int {
	return ensureValidMaxBatchSize(config.GetInt("process_config.max_per_message"))
}

func ensureValidMaxBatchSize(batchSize int) int {
	if batchSize <= 0 || batchSize > pkgconfigsetup.ProcessMaxPerMessageLimit {
		log.Warnf("Invalid max item count per message (%d), using default value of %d", batchSize, pkgconfigsetup.DefaultProcessMaxPerMessage)
		return pkgconfigsetup.DefaultProcessMaxPerMessage
	}
	return batchSize
}

// getMaxBatchSize returns the maximum number of bytes in a check payload
var getMaxBatchBytes = func(config model.Reader) int {
	return ensureValidMaxBatchBytes(config.GetInt("process_config.max_message_bytes"))
}

func ensureValidMaxBatchBytes(batchBytes int) int {
	if batchBytes <= 0 || batchBytes > pkgconfigsetup.ProcessMaxMessageBytesLimit {
		log.Warnf("Invalid max byte size per message (%d), using default value of %d", batchBytes, pkgconfigsetup.DefaultProcessMaxMessageBytes)
		return pkgconfigsetup.DefaultProcessMaxMessageBytes
	}
	return batchBytes
}
