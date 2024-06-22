package main

import "os"

func main() {
	log.Info("Starting newsletter service")
	log.Debugf("Config values: %v", config)

	gl, err := NewGitLab()
	if err != nil {
		log.Errorf("Error creating GitLab instance: %s", err.Error())
		os.Exit(1)
	}

	gitlabReport, err := gl.GenerateReport()
	if err != nil {
		log.Errorf("Error generating GitLab report: %s", err.Error())
		os.Exit(1)
	}

	if gitlabReport == "" {
		log.Info("No new commits found")
	} else {
		gitlabReport += "GitLab report\n"
		_, err = bot.SendMessage(gitlabReport)
		if err != nil {
			log.Errorf("Error sending GitLab report: %s", err.Error())
			os.Exit(1)
		}
	}
}
