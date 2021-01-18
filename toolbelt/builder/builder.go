package builder

import (
	"fmt"
	"github.com/apis-list/apis-list/toolbelt/list"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

var slugRegexp = regexp.MustCompile("[^a-zA-Z0-9 ]")

const readmeTmplName = "readme.gomd"
const apiTmplName = "api.gomd"

func Render(l list.APIs, dir string) error {
	funcs := template.FuncMap{
		"short": func(s string) string {
			s = firstSentence(s)
			s = strings.ReplaceAll(s, "\n", ". ")
			return strings.TrimSpace(s)
		},
		"slug": slug,
		"sort": func(s []string) []string {
			sort.Strings(s)
			return s
		},
		"category_icon": func(s string) string {
			icon := CategoryIcons[s]
			if icon == "" {
				icon = "📃"
			}

			return icon
		},
	}
	templates, err := template.New("readme").Funcs(funcs).ParseGlob(path.Join(dir, "templates", "*.gomd"))
	if err != nil {
		return err
	}

	readmeFile, err := os.OpenFile(path.Join(dir, "README.md"), os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer readmeFile.Close()

	categoriesIndex := make(map[string]string)
	for k := range l.ByCategory() {
		categoriesIndex[k] = slug(k)
	}
	err = templates.Lookup(readmeTmplName).Execute(readmeFile, map[string]interface{}{
		"CategoriesIndex": categoriesIndex,
		"APIs":            l.ByCategory(),
		"Graveyard":       l.Graveyard(),
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

		apiFile, err := os.OpenFile(trgt, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}
		defer apiFile.Close()

		err = templates.Lookup(apiTmplName).ExecuteTemplate(apiFile, apiTmplName, map[string]interface{}{
			"API":        a,
			"Libraries":  a.Libraries.ByPlatform(),
			"EditorLink": fmt.Sprintf("https://github.com/apis-list/apis-list/edit/main/apis.yaml#L%d", a.Line),
		})
		if err != nil {
			return err
		}

	}

	return nil
}

func firstSentence(s string) string {
	var sep = []string{"\n", ".", "!", "?"}
	for _, v := range sep {
		if strings.Contains(s, v) {
			return strings.Split(s, v)[0]
		}
	}

	return s
}

func slug(s string) string {
	s = slugRegexp.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, " ", "-")
	return strings.ToLower(s)
}
