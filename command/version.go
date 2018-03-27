package command

import "fmt"

const (
	versionNumber = "0.1"
	githubUrl     = "https://github.com/i183/gossh"
)

type Version struct {
}

func (v *Version) Init(args []string) bool {
	return true
}

func (v *Version) Execute() {
	fmt.Println("Version:", versionNumber)
	fmt.Println("Github:", githubUrl)
}

func (v *Version) Help() {
}
