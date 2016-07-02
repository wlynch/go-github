// Copyright 2016 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"fmt"
)

// AdminStatsService handles communication with the Admin Stats related methods
// of the GitHub Enterprise API.
//
// Note: These endpoints are only available to authenticated site administrators.
// Normal users will receive a 404 response if they try to access it.
//
// GitHub API docs:
// https://developer.github.com/enterprise/2.6/v3/enterprise/admin_stats/
type AdminStatsService service

// AdminStats represents a list of metrics regarding a GitHub Enterprise
// instance.
type AdminStats struct {
	Comments   *CommentStats   `json:"comments,omitempty"`
	Gists      *GistStats      `json:"gists,omitempty"`
	Hooks      *HookStats      `json:"hooks,omitempty"`
	Issues     *IssueStats     `json:"issues,omitempty"`
	Orgs       *OrgStats       `json:"orgs,omitempty"`
	Pages      *PageStats      `json:"pages,omitempty"`
	Pulls      *PullStats      `json:"pulls,omitempty"`
	Repos      *RepoStats      `json:"repos,omitempty"`
	Users      *UserStats      `json:"users,omitempty"`
	Milestones *MilestoneStats `json:"milestones,omitempty"`
}

func (a AdminStats) String() string {
	return Stringify(a)
}

// CommentStats represents the number of comments on issues, pull requests, commits,
// and gists.
type CommentStats struct {
	TotalCommitComments      int `json:"total_commit_comments"`
	TotalGistComments        int `json:"total_gist_comments"`
	TotalIssueComments       int `json:"total_issue_comments"`
	TotalPullRequestComments int `json:"total_pull_request_comments"`
}

// GistStats represents the number of private and public gists.
type GistStats struct {
	PrivateGists int `json:"private_gists"`
	PublicGists  int `json:"public_gists"`
	TotalGists   int `json:"total_gists"`
}

// HookStats represents the number of active and inactive hooks.
type HookStats struct {
	ActiveHooks   int `json:"active_hooks"`
	InactiveHooks int `json:"inactive_hooks"`
	TotalHooks    int `json:"total_hooks"`
}

// IssueStats represents the number of open and closed issues.
type IssueStats struct {
	ClosedIssues int `json:"closed_issues"`
	OpenIssues   int `json:"open_issues"`
	TotalIssues  int `json:"total_issues"`
}

// MilestoneStats represents the number of open and closed milestones.
type MilestoneStats struct {
	ClosedMilestones int `json:"closed_milestones"`
	OpenMilestones   int `json:"open_milestones"`
	TotalMilestones  int `json:"total_milestones"`
}

// OrgStats represents the number of organizations, teams, team members, and
// disabled organizations.
type OrgStats struct {
	DisabledOrgs     int `json:"disabled_orgs"`
	TotalOrgs        int `json:"total_orgs"`
	TotalTeamMembers int `json:"total_team_members"`
	TotalTeams       int `json:"total_teams"`
}

// PageStats represents the number of GitHub Pages sites.
type PageStats struct {
	TotalPages int `json:"total_pages"`
}

// PullStats represents the number of merged, mergeable, and unmergeable pull
// requests.
type PullStats struct {
	MergeablePulls   int `json:"mergeable_pulls"`
	MergedPulls      int `json:"merged_pulls"`
	TotalPulls       int `json:"total_pulls"`
	UnmergeablePulls int `json:"unmergeable_pulls"`
}

// RepoStats represents the number of organization-owned repositories, root
// repositories, forks, pushed commits, and wikis.
type RepoStats struct {
	ForkRepos   int `json:"fork_repos"`
	OrgRepos    int `json:"org_repos"`
	RootRepos   int `json:"root_repos"`
	TotalPushes int `json:"total_pushes"`
	TotalRepos  int `json:"total_repos"`
	TotalWikis  int `json:"total_wikis"`
}

// UserStats represents the number of suspended and admin users.
type UserStats struct {
	AdminUsers     int `json:"admin_users"`
	SuspendedUsers int `json:"suspended_users"`
	TotalUsers     int `json:"total_users"`
}

// Get fetches the statistics for a GitHub Enterprise installation.
func (s *AdminStatsService) Get(statType string) (*AdminStats, *Response, error) {
	u := fmt.Sprintf("enterprise/stats/%s", statType)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	stats := new(AdminStats)
	resp, err := s.client.Do(req, stats)
	if err != nil {
		return nil, resp, err
	}
	return stats, resp, err
}
