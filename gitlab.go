package main

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
	"time"
)

type GitLab struct {
	client *gitlab.Client
}

func createClient() (*gitlab.Client, error) {
	log.Debug("Creating gitlab client")

	client, err := gitlab.NewClient(config.GitLabToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return client, nil
}

func NewGitLab() (*GitLab, error) {
	log.Debug("Creating gitlab instance")

	client, err := createClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return &GitLab{
		client: client,
	}, nil
}

func getSubgroups(client *gitlab.Client, groupID int) ([]*gitlab.Group, error) {
	log.Debugf("Getting subgroups for group %d", groupID)

	groups, _, err := client.Groups.ListSubGroups(groupID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list groups: %w", err)
	}
	log.Debugf("Found %d subgroups", len(groups))

	allGroups := groups
	for _, group := range groups {
		subgroups, err := getSubgroups(client, group.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get subgroups: %w", err)
		}
		allGroups = append(allGroups, subgroups...)
	}

	return allGroups, nil
}

func (g *GitLab) GetSubgroups() ([]*gitlab.Group, error) {
	log.Info("Getting subgroups")

	groups, err := getSubgroups(g.client, config.GitLabGroupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get subgroups: %w", err)
	}

	return groups, nil
}

func (g *GitLab) GetProjects() ([]*gitlab.Project, error) {
	log.Info("Getting projects")

	var projects []*gitlab.Project

	rootGroupProjects, _, err := g.client.Groups.ListGroupProjects(config.GitLabGroupID, &gitlab.ListGroupProjectsOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list group projects: %w", err)
	}
	projects = append(projects, rootGroupProjects...)

	groups, err := g.GetSubgroups()
	if err != nil {
		return nil, fmt.Errorf("failed to get subgroups: %w", err)
	}

	for _, group := range groups {
		log.Debugf("Group: %s", group.Name)

		groupProjects, _, err := g.client.Groups.ListGroupProjects(group.ID, &gitlab.ListGroupProjectsOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to list group projects: %w", err)
		}

		projects = append(projects, groupProjects...)
	}

	return projects, nil
}

func (g *GitLab) GetCommits(projectID int) ([]*gitlab.Commit, error) {
	log.Infof("Getting commits for project %d", projectID)

	currentTime := time.Now()
	oneDayAgo := currentTime.AddDate(0, 0, -1)

	commits, _, err := g.client.Commits.ListCommits(projectID, &gitlab.ListCommitsOptions{
		Since: &oneDayAgo,
		Until: &currentTime,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list commits: %w", err)
	}

	return commits, nil
}

func (g *GitLab) RenderCommits(project *gitlab.Project, commits []*gitlab.Commit) string {
	log.Debugf("Rendering commits")

	message := fmt.Sprintf("<b>%s:</b>\n<blockquote expandable>", project.Name)
	for _, commit := range commits {
		message += fmt.Sprintf(
			"📝 <a href='%s'>%s</a> by %s\n",
			commit.WebURL,
			commit.Title,
			commit.AuthorName,
		)
	}
	message += "</blockquote>"

	return message
}

func (g *GitLab) GenerateReport() string {
	report := ""

	projects, err := g.GetProjects()
	if err != nil {
		log.Errorf("Error getting projects: %s", err.Error())
	}

	for _, project := range projects {
		log.Infof("Project: %s", project.Name)

		commits, err := g.GetCommits(project.ID)
		if err != nil {
			log.Errorf("Error getting commits: %s", err.Error())
			continue
		}
		if len(commits) == 0 {
			log.Infof("No commits found for project %s", project.Name)
			continue
		}

		log.Infof("Commits: %d", len(commits))

		report += g.RenderCommits(project, commits) + "\n"
	}

	if report == "" {
		return ""
	} else {
		report = "GitLab report\n" + report
		return report
	}
}
