package main

func main() {
	log.Info("Starting newsletter service")
	log.Debugf("Config values: %v", config)

	gl, err := NewGitLab()
	if err != nil {
		log.Errorf("Error creating GitLab instance: %s", err.Error())
		return
	}

	gitlabReport := gl.GenerateReport()
	if gitlabReport == "" {
		log.Info("No new commits found")
	} else {
		gitlabReport += "GitLab report\n"
		bot.SendMessage(gitlabReport)
	}
}
