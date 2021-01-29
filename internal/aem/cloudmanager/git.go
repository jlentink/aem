package cloudmanager

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/zalando/go-keyring"
	"os"
)

// AdobeCloudManagerRemote ...
const AdobeCloudManagerRemote = "AdobeCloudManagerRemote"

func setupGit() (*git.Repository, error){
	cwd, err := project.GetWorkDir()
	if err != nil {
		return nil, err
	}
	repo, err := git.PlainOpen(cwd)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func setupRemote(repo *git.Repository, cnf *objects.Config) error {
	remotes, err := repo.Remotes()
	if err != nil {
		return err
	}
	for _, remote := range remotes {
		if remote.Config().Name == AdobeCloudManagerRemote {
			return nil
		}
	}

	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: AdobeCloudManagerRemote,
		URLs: []string{cnf.CloudManagerGit},
	})

	return err
}

func getGitAuth(){
	//http.BasicAuth{Username: , Password: }
}

// GitPush ...
func GitPush(cnf *objects.Config) error {

	repo, err := setupGit()
	if err != nil {
		return err
	}

	err = setupRemote(repo, cnf)
	if err != nil {
		return err
	}

	err = repo.Push(&git.PushOptions{
		RemoteName: AdobeCloudManagerRemote,
		//Auth: ,
		Progress: os.Stdout,
	})

	return err
}

// GitSetAuthentication ...
func GitSetAuthentication(username, password string, cnf *objects.Config) error {
	return keyring.Set(fmt.Sprintf("%s-%s", cnf.ProjectName,AdobeCloudManagerRemote), username, password)
}