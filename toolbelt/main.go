package main

import (
	"fmt"
	"github.com/apis-list/apis-list/toolbelt/builder"
	"github.com/apis-list/apis-list/toolbelt/list"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "build",
				Action: func(c *cli.Context) error {
					apis, err := list.ReadList(c.Args().First())
					if err != nil {
						return err
					}

					dir, err := os.Getwd()
					if err != nil {
						return err
					}

					return builder.Render(apis, dir)
				},
			},
			{
				Name: "check-links",
				Action: func(c *cli.Context) error {
					apis, err := list.ReadList(c.Args().First())
					if err != nil {
						return err
					}

					for _, v := range apis {
						if v.IsActive == false {
							continue
						}

						for _, vv := range v.Links {
							resp, err := http.Get(vv.Url)
							if err != nil || resp.StatusCode != http.StatusOK {
								if err == nil {
									err = fmt.Errorf("unexpected status code = %d", resp.StatusCode)
								}
								fmt.Printf("Wrong link %s for %s (%s) - %s\r\n ", vv.Name, v.Name, vv.Url, err)
							}
						}

					}

					return nil
				},
			},
			{
				Name: "check-libraries-links",
				Action: func(c *cli.Context) error {
					apis, err := list.ReadList(c.Args().First())
					if err != nil {
						return err
					}

					for _, v := range apis {
						if v.IsActive == false {
							continue
						}

						for _, vv := range v.Libraries {
							resp, err := http.Get(vv.DocumentationURL)
							if err != nil || resp.StatusCode != http.StatusOK {
								if err == nil {
									err = fmt.Errorf("unexpected status code = %d", resp.StatusCode)
								}
								fmt.Printf("Wrong docs link for %s - %s (%s): %s\r\n ", v.Name, vv.Name, vv.DocumentationURL, err)
							}

							resp, err = http.Get(vv.SourceCodeURL)
							if err != nil || resp.StatusCode != http.StatusOK {
								if err == nil {
									err = fmt.Errorf("unexpected status code = %d", resp.StatusCode)
								}
								fmt.Printf("Wrong source code link for %s - %s (%s): %s\r\n ", v.Name, vv.Name, vv.SourceCodeURL, err)
							}
						}
					}

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
