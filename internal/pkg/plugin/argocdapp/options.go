package argocdapp

import (
	"path/filepath"

	"github.com/mitchellh/mapstructure"

	"github.com/devstream-io/devstream/internal/pkg/configmanager"
	"github.com/devstream-io/devstream/pkg/util/downloader"
	"github.com/devstream-io/devstream/pkg/util/file"
	"github.com/devstream-io/devstream/pkg/util/log"
	"github.com/devstream-io/devstream/pkg/util/scm"
	"github.com/devstream-io/devstream/pkg/util/scm/git"
	"github.com/devstream-io/devstream/pkg/util/template"
)

// options is the struct for parameters used by the argocdapp package.
type options struct {
	App         *app         `mapstructure:"app" validate:"required"`
	Destination *destination `mapstructure:"destination"`
	Source      *source      `mapstructure:"source" validate:"required"`
	ImageRepo   imageRepo    `mapstructure:"imageRepo"`
}

type imageRepo struct {
	Address string `mapstructure:"address"`
	Tag     string `mapstructure:"tag"`
}

// app is the struct for an ArgoCD app.
type app struct {
	Name      string `mapstructure:"name" validate:"dns1123subdomain"`
	Namespace string `mapstructure:"namespace"`
}

// destination is the struct for the destination of an ArgoCD app.
type destination struct {
	Server    string `mapstructure:"server"`
	Namespace string `mapstructure:"namespace"`
}

// source is the struct for the source of an ArgoCD app.
type source struct {
	Valuefile string `mapstructure:"valuefile"`
	Path      string `mapstructure:"path" validate:"required"`
	RepoURL   string `mapstructure:"repoURL" validate:"required"`
}

// / newOptions create options by raw options
func newOptions(rawOptions configmanager.RawOptions) (options, error) {
	var opts options
	if err := mapstructure.Decode(rawOptions, &opts); err != nil {
		return opts, err
	}
	return opts, nil
}

// checkPathExist will check argocdapp config already exist
func (s *source) checkPathExist(scmClient scm.ClientOperation) (bool, error) {
	existFiles, err := scmClient.GetPathInfo(s.Path)
	if err != nil {
		return false, err
	}
	if len(existFiles) > 0 {
		log.Debugf("argocdapp check config path %s already exist, pass...", s.Path)
		return true, nil
	}
	return false, nil
}

func (o *options) getArgocdDefaultConfigFiles(configLocation downloader.ResourceLocation) (git.GitFileContentMap, error) {
	// 1. get configs from remote url
	configFiles, err := configLocation.Download()
	if err != nil {
		return nil, err
	}
	// 2. get file content
	fContentFunc := func(filePath string) ([]byte, error) {
		renderContent, err := template.New().FromLocalFile(filePath).SetDefaultRender("argocd config", o).Render()
		if err != nil {
			return nil, err
		}
		return []byte(renderContent), nil
	}
	fNameFunc := func(filePath, srcPath string) string {
		relativePath, _ := filepath.Rel(srcPath, filePath)
		return filepath.Join(o.Source.Path, relativePath)
	}
	return file.GetFileMap(configFiles, file.DirFileFilterDefaultFunc, fNameFunc, fContentFunc)
}
