// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    githubPR, err := UnmarshalGithubPR(bytes)
//    bytes, err = githubPR.Marshal()

package pr

import "encoding/json"

func UnmarshalGithubPR(data []byte) (GithubPR, error) {
	var r GithubPR
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GithubPR) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GithubPR struct {
	Action      string      `json:"action"`
	Number      int64       `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repo        `json:"repository"`
	Sender      Sender      `json:"sender"`
}

type PullRequest struct {
	URL                 string        `json:"url"`
	ID                  int64         `json:"id"`
	NodeID              string        `json:"node_id"`
	HTMLURL             string        `json:"html_url"`
	DiffURL             string        `json:"diff_url"`
	PatchURL            string        `json:"patch_url"`
	IssueURL            string        `json:"issue_url"`
	Number              int64         `json:"number"`
	State               string        `json:"state"`
	Locked              bool          `json:"locked"`
	Title               string        `json:"title"`
	User                Sender        `json:"user"`
	Body                string        `json:"body"`
	CreatedAt           string        `json:"created_at"`
	UpdatedAt           string        `json:"updated_at"`
	ClosedAt            interface{}   `json:"closed_at"`
	MergedAt            interface{}   `json:"merged_at"`
	MergeCommitSHA      interface{}   `json:"merge_commit_sha"`
	Assignee            interface{}   `json:"assignee"`
	Assignees           []interface{} `json:"assignees"`
	RequestedReviewers  []interface{} `json:"requested_reviewers"`
	RequestedTeams      []interface{} `json:"requested_teams"`
	Labels              []interface{} `json:"labels"`
	Milestone           interface{}   `json:"milestone"`
	Draft               bool          `json:"draft"`
	CommitsURL          string        `json:"commits_url"`
	ReviewCommentsURL   string        `json:"review_comments_url"`
	ReviewCommentURL    string        `json:"review_comment_url"`
	CommentsURL         string        `json:"comments_url"`
	StatusesURL         string        `json:"statuses_url"`
	Head                Base          `json:"head"`
	Base                Base          `json:"base"`
	Links               Links         `json:"_links"`
	AuthorAssociation   string        `json:"author_association"`
	AutoMerge           interface{}   `json:"auto_merge"`
	ActiveLockReason    interface{}   `json:"active_lock_reason"`
	Merged              bool          `json:"merged"`
	Mergeable           interface{}   `json:"mergeable"`
	Rebaseable          interface{}   `json:"rebaseable"`
	MergeableState      string        `json:"mergeable_state"`
	MergedBy            interface{}   `json:"merged_by"`
	Comments            int64         `json:"comments"`
	ReviewComments      int64         `json:"review_comments"`
	MaintainerCanModify bool          `json:"maintainer_can_modify"`
	Commits             int64         `json:"commits"`
	Additions           int64         `json:"additions"`
	Deletions           int64         `json:"deletions"`
	ChangedFiles        int64         `json:"changed_files"`
}

type Base struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	SHA   string `json:"sha"`
	User  Sender `json:"user"`
	Repo  Repo   `json:"repo"`
}

type Repo struct {
	ID                        int64         `json:"id"`
	NodeID                    string        `json:"node_id"`
	Name                      string        `json:"name"`
	FullName                  string        `json:"full_name"`
	Private                   bool          `json:"private"`
	Owner                     Sender        `json:"owner"`
	HTMLURL                   string        `json:"html_url"`
	Description               string        `json:"description"`
	Fork                      bool          `json:"fork"`
	URL                       string        `json:"url"`
	ForksURL                  string        `json:"forks_url"`
	KeysURL                   string        `json:"keys_url"`
	CollaboratorsURL          string        `json:"collaborators_url"`
	TeamsURL                  string        `json:"teams_url"`
	HooksURL                  string        `json:"hooks_url"`
	IssueEventsURL            string        `json:"issue_events_url"`
	EventsURL                 string        `json:"events_url"`
	AssigneesURL              string        `json:"assignees_url"`
	BranchesURL               string        `json:"branches_url"`
	TagsURL                   string        `json:"tags_url"`
	BlobsURL                  string        `json:"blobs_url"`
	GitTagsURL                string        `json:"git_tags_url"`
	GitRefsURL                string        `json:"git_refs_url"`
	TreesURL                  string        `json:"trees_url"`
	StatusesURL               string        `json:"statuses_url"`
	LanguagesURL              string        `json:"languages_url"`
	StargazersURL             string        `json:"stargazers_url"`
	ContributorsURL           string        `json:"contributors_url"`
	SubscribersURL            string        `json:"subscribers_url"`
	SubscriptionURL           string        `json:"subscription_url"`
	CommitsURL                string        `json:"commits_url"`
	GitCommitsURL             string        `json:"git_commits_url"`
	CommentsURL               string        `json:"comments_url"`
	IssueCommentURL           string        `json:"issue_comment_url"`
	ContentsURL               string        `json:"contents_url"`
	CompareURL                string        `json:"compare_url"`
	MergesURL                 string        `json:"merges_url"`
	ArchiveURL                string        `json:"archive_url"`
	DownloadsURL              string        `json:"downloads_url"`
	IssuesURL                 string        `json:"issues_url"`
	PullsURL                  string        `json:"pulls_url"`
	MilestonesURL             string        `json:"milestones_url"`
	NotificationsURL          string        `json:"notifications_url"`
	LabelsURL                 string        `json:"labels_url"`
	ReleasesURL               string        `json:"releases_url"`
	DeploymentsURL            string        `json:"deployments_url"`
	CreatedAt                 string        `json:"created_at"`
	UpdatedAt                 string        `json:"updated_at"`
	PushedAt                  string        `json:"pushed_at"`
	GitURL                    string        `json:"git_url"`
	SSHURL                    string        `json:"ssh_url"`
	CloneURL                  string        `json:"clone_url"`
	SvnURL                    string        `json:"svn_url"`
	Homepage                  interface{}   `json:"homepage"`
	Size                      int64         `json:"size"`
	StargazersCount           int64         `json:"stargazers_count"`
	WatchersCount             int64         `json:"watchers_count"`
	Language                  string        `json:"language"`
	HasIssues                 bool          `json:"has_issues"`
	HasProjects               bool          `json:"has_projects"`
	HasDownloads              bool          `json:"has_downloads"`
	HasWiki                   bool          `json:"has_wiki"`
	HasPages                  bool          `json:"has_pages"`
	HasDiscussions            bool          `json:"has_discussions"`
	ForksCount                int64         `json:"forks_count"`
	MirrorURL                 interface{}   `json:"mirror_url"`
	Archived                  bool          `json:"archived"`
	Disabled                  bool          `json:"disabled"`
	OpenIssuesCount           int64         `json:"open_issues_count"`
	License                   interface{}   `json:"license"`
	AllowForking              bool          `json:"allow_forking"`
	IsTemplate                bool          `json:"is_template"`
	WebCommitSignoffRequired  bool          `json:"web_commit_signoff_required"`
	Topics                    []interface{} `json:"topics"`
	Visibility                string        `json:"visibility"`
	Forks                     int64         `json:"forks"`
	OpenIssues                int64         `json:"open_issues"`
	Watchers                  int64         `json:"watchers"`
	DefaultBranch             string        `json:"default_branch"`
	AllowSquashMerge          *bool         `json:"allow_squash_merge,omitempty"`
	AllowMergeCommit          *bool         `json:"allow_merge_commit,omitempty"`
	AllowRebaseMerge          *bool         `json:"allow_rebase_merge,omitempty"`
	AllowAutoMerge            *bool         `json:"allow_auto_merge,omitempty"`
	DeleteBranchOnMerge       *bool         `json:"delete_branch_on_merge,omitempty"`
	AllowUpdateBranch         *bool         `json:"allow_update_branch,omitempty"`
	UseSquashPRTitleAsDefault *bool         `json:"use_squash_pr_title_as_default,omitempty"`
	SquashMergeCommitMessage  *string       `json:"squash_merge_commit_message,omitempty"`
	SquashMergeCommitTitle    *string       `json:"squash_merge_commit_title,omitempty"`
	MergeCommitMessage        *string       `json:"merge_commit_message,omitempty"`
	MergeCommitTitle          *string       `json:"merge_commit_title,omitempty"`
}

type Sender struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Links struct {
	Self           Comments `json:"self"`
	HTML           Comments `json:"html"`
	Issue          Comments `json:"issue"`
	Comments       Comments `json:"comments"`
	ReviewComments Comments `json:"review_comments"`
	ReviewComment  Comments `json:"review_comment"`
	Commits        Comments `json:"commits"`
	Statuses       Comments `json:"statuses"`
}

type Comments struct {
	Href string `json:"href"`
}
