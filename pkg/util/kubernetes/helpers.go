// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package kubernetes

import (
	"strconv"
	"strings"
)

// KubeAllowedEncodeStringAlphaNums holds the charactes allowed in replicaset names from as parent deployment
// Taken from https://github.com/kow3ns/kubernetes/blob/96067e6d7b24a05a6a68a0d94db622957448b5ab/staging/src/k8s.io/apimachinery/pkg/util/rand/rand.go#L76
const KubeAllowedEncodeStringAlphaNums = "bcdfghjklmnpqrstvwxz2456789"

// Digits holds the digits used for naming replicasets in kubenetes < 1.8
const Digits = "1234567890"

// ParseDeploymentForReplicaSet gets the deployment name from a replicaset,
// or returns an empty string if no parent deployment is found.
func ParseDeploymentForReplicaSet(name string) string {
	return removeKubernetesNameSuffix(name)
}

// ParseDeploymentForPodName gets the deployment name from a pod name,
// or returns an empty string if no parent deployment is found.
func ParseDeploymentForPodName(name string) string {
	replicaSet := removeKubernetesNameSuffix(name)
	if replicaSet == "" {
		return ""
	}
	return ParseDeploymentForReplicaSet(replicaSet)
}

// ParseReplicaSetForPodName gets the replica set name from a pod name,
// or returns an empty string if no parent replica set is found.
func ParseReplicaSetForPodName(name string) string {
	return removeKubernetesNameSuffix(name)
}

// ParseCronJobForJob gets the cronjob name from a job,
// or returns an empty string if no parent cronjob is found.
// https://github.com/kubernetes/kubernetes/blob/b4e3bd381bd4d7c0db1959341b39558b45187345/pkg/controller/cronjob/utils.go#L156
func ParseCronJobForJob(name string) (string, int) {
	lastDash := strings.LastIndexByte(name, '-')
	if lastDash == -1 {
		// No dash
		return "", 0
	}
	suffix := name[lastDash+1:]
	if len(suffix) < 3 {
		// Suffix is variable length but we cutoff at 3+ characters
		return "", 0
	}

	if !stringInRuneset(suffix, Digits) {
		// Invalid suffix
		return "", 0
	}

	id, err := strconv.Atoi(suffix)
	if err != nil {
		// Cannot happen because of the test just above
		return "", 0
	}

	return name[:lastDash], id
}

// stringInRuneset tests whether all runes of a string are in a given subset
// returns false if any rune in the string is not found in the subset
func stringInRuneset(name, subset string) bool {
	for _, r := range name {
		if !strings.ContainsRune(subset, r) {
			// Found an unexpected rune in suffix
			return false
		}
	}
	return true
}

// removeKubernetesNameSuffix removes the suffix from a kubernetes name
// or returns an empty string if either the suffix or name are invalid.
func removeKubernetesNameSuffix(name string) string {
	lastDash := strings.LastIndexByte(name, '-')
	if lastDash == -1 {
		// No dash
		return ""
	}
	suffix := name[lastDash+1:]
	if len(suffix) < 3 {
		// Suffix is variable length but we cutoff at 3+ characters
		return ""
	}

	if !stringInRuneset(suffix, Digits) && !stringInRuneset(suffix, KubeAllowedEncodeStringAlphaNums) {
		// Invalid suffix
		return ""
	}

	return name[:lastDash]
}
