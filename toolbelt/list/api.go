package list

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path"
	"path/filepath"
)

type APIs []*API
type Libraries []*APILibrary

type API struct {
	Name          string   `yaml:"name"`
	Slug          string   `yaml:"slug"`
	Categories    []string `yaml:"categories"`
	Description   string   `yaml:"description"`
	IsFree        bool     `yaml:"is_free"`
	Logo          string   `yaml:"logo,omitempty"`
	DiscussionURL string   `yaml:"discussion_url,omitempty"`
	Type          string   `yaml:"type,omitempty"`
	Contact       string   `yaml:"contact,omitempty"`
	IsActive      bool     `yaml:"is_active"`

	Specification APISpecification `yaml:"specification,omitempty"`

	Libraries Libraries `yaml:"libraries,omitempty"`
	Links     []APILink `yaml:"links,omitempty"`
}

type APISpecification struct {
	Type    string `yaml:"type,omitempty"`
	Url     string `yaml:"url,omitempty"`
	Version string `yaml:"version,omitempty"`
}

type APILibrary struct {
	Name             string `yaml:"name"`
	Description      string `yaml:"-"`
	DocumentationURL string `yaml:"homepage_uri" yaml:"homepage_uri,omitempty"`
	SourceCodeURL    string `yaml:"source_code_uri" yaml:"source_code_uri"`
	Version          string `yaml:"version,omitempty"`
	Platform         string `yaml:"platform"`
}

type APILink struct {
	Name string
	Url  string
}

func (a APIs) ByCategory() map[string][]*API {
	cm := make(map[string][]*API)
	for _, v := range a {
		for _, vv := range v.Categories {
			if _, ok := cm[vv]; !ok {
				cm[vv] = make([]*API, 0)
			}
			cm[vv] = append(cm[vv], v)
		}
	}

	return cm
}

func (a APIs) Graveyard() []*API {
	cm := make([]*API, 0)
	for _, v := range a {
		if v.IsActive == false {
			cm = append(cm, v)
		}
	}

	return cm
}

func (a Libraries) ByPlatform() map[string][]*APILibrary {
	cm := make(map[string][]*APILibrary)
	for _, v := range a {
		if _, ok := cm[v.Platform]; !ok {
			cm[v.Platform] = make([]*APILibrary, 0)
		}
		cm[v.Platform] = append(cm[v.Platform], v)
	}

	return cm
}

func ReadList(repositoryPath string) ([]*API, error) {
	var list []*API

	matches, _ := filepath.Glob(path.Join(repositoryPath, "apis", "*/*.yaml"))
	for _, match := range matches {
		var api API

		data, err := ioutil.ReadFile(match)
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal(data, &api)
		if err != nil {
			return nil, err
		}

		api.Slug = path.Base(path.Dir(match))

		list = append(list, &api)
	}

	return list, nil
}
