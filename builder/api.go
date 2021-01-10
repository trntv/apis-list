package builder

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type APIs []*API
type Libraries []*APILibrary

type API struct {
	Name          string   `yaml:"name"`
	Slug          string   `yaml:"slug"`
	Categories    []string `yaml:"categories"`
	Description   string   `yaml:"description"`
	URI           string   `yaml:"uri"`
	IsPaid        bool     `yaml:"is_paid"`
	Logo          string   `yaml:"logo,omitempty"`
	DiscussionURI string   `yaml:"discussion_uri,omitempty"`
	Type          string   `yaml:"type"`
	Contact       string   `yaml:"contact,omitempty"`
	IsActive        bool     `yaml:"is_active"`

	Libraries Libraries `yaml:"libraries,omitempty"`
}

type APILibrary struct {
	Name          string `yaml:"name"`
	Description   string `yaml:"-"`
	HomepageURI   string `yaml:"homepage_uri" yaml:"homepage_uri,omitempty"`
	SourceCodeURI string `yaml:"source_code_uri" yaml:"source_code_uri"`
	Version       string `yaml:"version,omitempty"`
	Platform      string `yaml:"platform"`
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

func ReadList(path string) ([]*API, error) {
	apisData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var list []*API
	err = yaml.Unmarshal(apisData, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
