// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

// HookConfig represents configuration settings for an arbitrary hook type.
type HookConfig interface{}

// WebHookConfig represents the configuration settings for web hooks.
// This type of hook trigger PostReceiveHooks to be sent.
type WebHookConfig struct {
	URL         string `json:"url,omitempty"`
	Secret      string `json:"secret,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	SSLVersion  string `json:"ssl_version,omitempty"`
	InsecureSSL bool   `json:"insecure_ssl,omitempty"`
}

// NewWebHook creates a new hook with the given user configurations.
func NewWebHook(events []string, active bool, config WebHookConfig) *Hook {
	return &Hook{
		"", nil, nil, 0, HookOptions{
			"Web", events, active, config,
		},
	}
}
