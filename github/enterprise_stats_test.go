// Copyright 2016 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAdminStatsUnmarshal(t *testing.T) {
	payload := `{
  "repos": {
    "total_repos": 212,
    "root_repos": 194,
    "fork_repos": 18,
    "org_repos": 51,
    "total_pushes": 3082,
    "total_wikis": 15
  },
  "hooks": {
    "total_hooks": 27,
    "active_hooks": 23,
    "inactive_hooks": 4
  },
  "pages": {
    "total_pages": 36
  },
  "orgs": {
    "total_orgs": 33,
    "disabled_orgs": 0,
    "total_teams": 60,
    "total_team_members": 314
  },
  "users": {
    "total_users": 254,
    "admin_users": 45,
    "suspended_users": 21
  },
  "pulls": {
    "total_pulls": 86,
    "merged_pulls": 60,
    "mergeable_pulls": 21,
    "unmergeable_pulls": 3
  },
  "issues": {
    "total_issues": 179,
    "open_issues": 83,
    "closed_issues": 96
  },
  "milestones": {
    "total_milestones": 7,
    "open_milestones": 6,
    "closed_milestones": 1
  },
  "gists": {
    "total_gists": 178,
    "private_gists": 151,
    "public_gists": 25
  },
  "comments": {
    "total_commit_comments": 6,
    "total_gist_comments": 28,
    "total_issue_comments": 366,
    "total_pull_request_comments": 30
  }
}`
	got := new(AdminStats)
	json.Unmarshal([]byte(payload), got)

	want := &AdminStats{
		Repos: &RepoStats{
			TotalRepos:  212,
			RootRepos:   194,
			ForkRepos:   18,
			OrgRepos:    51,
			TotalPushes: 3082,
			TotalWikis:  15,
		},
		Hooks: &HookStats{
			TotalHooks:    27,
			ActiveHooks:   23,
			InactiveHooks: 4,
		},
		Pages: &PageStats{
			TotalPages: 36,
		},
		Orgs: &OrgStats{
			TotalOrgs:        33,
			DisabledOrgs:     0,
			TotalTeams:       60,
			TotalTeamMembers: 314,
		},
		Users: &UserStats{
			TotalUsers:     254,
			AdminUsers:     45,
			SuspendedUsers: 21,
		},
		Pulls: &PullStats{
			TotalPulls:       86,
			MergedPulls:      60,
			MergeablePulls:   21,
			UnmergeablePulls: 3,
		},
		Issues: &IssueStats{
			TotalIssues:  179,
			OpenIssues:   83,
			ClosedIssues: 96,
		},
		Milestones: &MilestoneStats{
			TotalMilestones:  7,
			OpenMilestones:   6,
			ClosedMilestones: 1,
		},
		Gists: &GistStats{
			TotalGists:   178,
			PrivateGists: 151,
			PublicGists:  25,
		},
		Comments: &CommentStats{
			TotalCommitComments:      6,
			TotalGistComments:        28,
			TotalIssueComments:       366,
			TotalPullRequestComments: 30,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("AdminStats unmarshal: got %+v, want %+v", got, want)
	}
}
