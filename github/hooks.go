// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"fmt"
	"time"
)

// Hook represents a github hook.
type Hook struct {
	URL         string     `json:"url,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	ID          int        `json:"id,omitempty"`
	HookOptions
}

// HookOptions are user configurable options for hooks.
type HookOptions struct {
	Name       string                 `json:"name,omitempty"`
	Events     []string               `json:"events,omitempty"`
	Active     bool                   `json:"active,omitempty"`
	HookConfig          `json:"config,omitempty"`
}

type HookConfig map[string]interface{}

// ListHooks retrieves a list of hooks for the given repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/hooks/#list
func (s *RepositoriesService) ListHooks(owner, repo string) ([]Hook, error) {
	u := fmt.Sprintf("repos/%v/%v/hooks", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	hooks := new([]Hook)
	_, err = s.client.Do(req, hooks)
	return *hooks, err
}

// GetHook retrieves a single hook for the given repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/hooks/#get-single-hook
func (s *RepositoriesService) GetHook(owner, repo string, id uint) (*Hook, error) {
	u := fmt.Sprintf("repos/%v/%v/hooks/%v", owner, repo, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	hook := new(Hook)
	_, err = s.client.Do(req, hook)
	return hook, err
}

// CreateHook creates a hook to the authenticated user's given repo.
//
// GitHub API docs: http://developer.github.com/v3/repos/hooks/#create-a-hook
func (s *RepositoriesService) CreateHook(owner, repo string, hook *Hook) (*Hook, error) {
	u := fmt.Sprintf("repos/%v/%v/hooks", owner, repo)
	req, err := s.client.NewRequest("POST", u, hook.HookOptions)
	if err != nil {
		return nil, err
	}
	response := new(Hook)
	_, err = s.client.Do(req, response)
	return response, err
}

// EditHook edits a specified hook for the given repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/hooks/#edit-a-hook
func (s *RepositoriesService) EditHook(owner, repo string, hook *Hook) (*Hook, error) {
	u := fmt.Sprintf("repos/%v/%v/hooks/%v", owner, repo, hook.ID)
	req, err := s.client.NewRequest("PATCH", u, hook.HookOptions)
	if err != nil {
		return nil, err
	}
	response := new(Hook)
	_, err = s.client.Do(req, response)
	return response, err
}

// TestPushHook triggers a hook with the lastest push to the given repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/hooks/#test-a-push-hook
func (s *RepositoriesService) TestPushHook(owner, repo string, id int) error {
	u := fmt.Sprintf("repos/%v/%v/hooks/%v/tests", owner, repo, id)
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	return err
}

// DeleteHook deletes the specified hook from a given repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/hooks/#delete-a-hook
func (s *RepositoriesService) DeleteHook(owner, repo string, id int) error {
	u := fmt.Sprintf("repos/%v/%v/hooks/%v", owner, repo, id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req, nil)
	return err
}
