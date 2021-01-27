package builder

import (
	"fmt"
	"github.com/apis-list/apis-list/toolbelt/list"
	"gopkg.in/yaml.v3"
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
		"sort_categories": func(s []string) []string {
			sort.Sort(byCategoryName(s))
			return s
		},
		"category_icon": func(s string) string {
			icon := list.Categories[s]
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

	categoriesNames := make([]string, 0)
	for k := range l.ByCategory() {
		categoriesNames = append(categoriesNames, k)
	}
	sort.Sort(byCategoryName(categoriesNames))

	err = templates.Lookup(readmeTmplName).Execute(readmeFile, map[string]interface{}{
		"CategoriesNames": categoriesNames,
		"APIs":            l.ByCategory(),
		"Graveyard":       l.Graveyard(),
	})
	if err != nil {
		return err
	}

	var links = make([]string, len(l))
	for k, a := range l {
		trgt := path.Join(dir, "apis", a.Slug, fmt.Sprintf("%s.md", a.Slug))
		err := os.MkdirAll(path.Dir(trgt), os.ModePerm)
		if err != nil {
			return err
		}

		apiFile, err := os.OpenFile(trgt, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}
		err = templates.Lookup(apiTmplName).ExecuteTemplate(apiFile, apiTmplName, map[string]interface{}{
			"API":        a,
			"Libraries":  a.Libraries.ByPlatform(),
			"EditorLink": fmt.Sprintf("https://github.com/apis-list/apis-list/edit/main/apis/%s/%s.yaml", a.Slug, a.Slug),
		})

		apiFile.Close()

		if err != nil {
			return err
		}

		links[k] = fmt.Sprintf("https://raw.githubusercontent.com/apis-list/apis-list/main/apis/%s/%s.yaml", a.Slug, a.Slug)
	}

	index, err := yaml.Marshal(links)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path.Join(dir, "index.yaml"), os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(index)
	if err != nil {
		return err
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

type byCategoryName []string

func (a byCategoryName) Len() int      { return len(a) }
func (a byCategoryName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byCategoryName) Less(i, j int) bool {
	if a[i] == "Other" {
		return false
	}

	if a[j] == "Other" {
		return true
	}

	return a[i] < a[j]
}
