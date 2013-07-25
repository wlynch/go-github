// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

// NewWebHook creates a new hook with the given user configurations.
func NewWebHook(events []string, active bool, config HookConfig) *Hook {
	return &Hook{
		"", nil, nil, 0, HookOptions{
			"web", events, active, config,
		},
	}
}
