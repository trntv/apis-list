package builder

import (
	"fmt"
	"github.com/gosimple/slug"
	"os"
	"path"
	"strings"
	"text/template"
)

const readmeTmplName = "README.md.gotmpl"
const apiTmplName = "api.md.gotmpl"

func Render(l APIs, dir string) error {
	funcs := template.FuncMap{
		"slug": func(s string) string {
			return slug.MakeLang(s, "en")
		},
		"short": func(s string) string {
			return FirstSentence(s)
		},
	}
	templates, err := template.New("readme").Funcs(funcs).ParseGlob(path.Join(dir, "templates", "*.gotmpl"))
	if err != nil {
		return err
	}

	readmeFile, err := os.OpenFile(path.Join(dir, "README.md"), os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer readmeFile.Close()

	err = templates.Lookup(readmeTmplName).Execute(readmeFile, map[string]interface{}{
		"apis": l.ByCategory(),
		"graveyard": l.Graveyard(),
	})
	if err != nil {
		return err
	}

	for _, a := range l {
		trgt := path.Join(dir, "apis", a.Slug, fmt.Sprintf("%s.md", a.Slug))
		err := os.MkdirAll(path.Dir(trgt), os.ModePerm)
		if err != nil {
			return err
		}

		apiFile, err := os.OpenFile(trgt, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
		defer apiFile.Close()

		err = templates.Lookup(apiTmplName).ExecuteTemplate(apiFile, apiTmplName, map[string]interface{}{
			"api":       a,
			"libraries": a.Libraries.ByPlatform(),
		})
		if err != nil {
			return err
		}

	}

	return nil
}

func FirstSentence(s string) string {
	var sep = []string{".", "!", "?"}
	for _, v := range sep {
		if strings.Contains(s, v) {
			return strings.Split(s, v)[0] + v
		}
	}

	return s
}
