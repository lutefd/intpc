package app

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadInsomniaFile(filePath string) (*InsomniaExport, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var export InsomniaExport
	if err := yaml.Unmarshal(data, &export); err != nil {
		return nil, fmt.Errorf("failed to parse Insomnia YAML: %w", err)
	}

	return &export, nil
}

func WriteInsomniaFile(filePath string, export *InsomniaExport) error {
	data, err := yaml.Marshal(export)
	if err != nil {
		return fmt.Errorf("failed to marshal Insomnia YAML: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func ConvertPostmanToInsomnia(postman *PostmanCollection) (*InsomniaExport, error) {
	now := currentTimestamp()

	export := &InsomniaExport{
		Type: "collection.insomnia.rest/5.0",
		Name: postman.Info.Name,
		Meta: InsomniaMeta{
			ID:       "wrk_" + generateShortUUID(),
			Created:  now,
			Modified: now,
		},
		Collection: []InsomniaResource{},
		CookieJar: InsomniaCookieJar{
			Name: "Default Jar",
			Meta: InsomniaMeta{
				ID:       "jar_" + generateShortUUID(),
				Created:  now,
				Modified: now,
			},
		},
		Environments: InsomniaEnvironments{
			Name: "Base Environment",
			Meta: InsomniaMeta{
				ID:        "env_" + generateShortUUID(),
				Created:   now,
				Modified:  now,
				IsPrivate: false,
			},
		},
	}

	baseSort := now * -1
	processPostmanItems(postman.Items, &export.Collection, baseSort)

	return export, nil
}

func processPostmanItems(items []PostmanItem, parentResources *[]InsomniaResource, sortKeyBase int64) {
	for i, item := range items {
		now := currentTimestamp()
		sortKey := sortKeyBase - (int64(i) * 10000)
		if len(item.Items) > 0 {
			folder := InsomniaResource{
				Name: item.Name,
				Meta: InsomniaMeta{
					ID:       "fld_" + generateShortUUID(),
					Created:  now,
					Modified: now,
					SortKey:  sortKey,
				},
				Children: []InsomniaResource{},
			}

			processPostmanItems(item.Items, &folder.Children, sortKey-1000)

			*parentResources = append(*parentResources, folder)
		} else if item.Request != nil {
			request := InsomniaResource{
				Name:   item.Name,
				Method: item.Request.Method,
				URL:    item.Request.URL.Raw,
				Meta: InsomniaMeta{
					ID:        "req_" + generateShortUUID(),
					Created:   now,
					Modified:  now,
					IsPrivate: false,
					SortKey:   sortKey,
				},
				Settings: InsomniaSettings{
					RenderRequestBody: true,
					EncodeURL:         true,
					FollowRedirects:   "global",
					Cookies: InsomniaCookies{
						Send:  true,
						Store: true,
					},
					RebuildPath: true,
				},
			}

			if len(item.Request.Headers) > 0 {
				request.Headers = make([]InsomniaHeader, 0, len(item.Request.Headers))
				for _, header := range item.Request.Headers {
					if !header.Disabled {
						request.Headers = append(request.Headers, InsomniaHeader{
							Name:  header.Key,
							Value: header.Value,
						})
					}
				}
			}

			if len(item.Request.URL.Query) > 0 {
				request.Parameters = make([]InsomniaParameter, 0, len(item.Request.URL.Query))
				for _, query := range item.Request.URL.Query {
					if !query.Disabled {
						request.Parameters = append(request.Parameters, InsomniaParameter{
							Name:     query.Key,
							Value:    query.Value,
							Disabled: false,
						})
					}
				}
			}

			if item.Request.Body != nil && item.Request.Body.Raw != "" {
				mimeType := "text/plain"

				if item.Request.Body.Options != nil {
					if rawOpts, ok := item.Request.Body.Options["raw"].(map[string]interface{}); ok {
						if lang, ok := rawOpts["language"].(string); ok && lang == "json" {
							mimeType = "application/json"
						}
					}
				}

				request.Body = &InsomniaRequestBody{
					Text:     item.Request.Body.Raw,
					MimeType: mimeType,
				}
			}

			*parentResources = append(*parentResources, request)
		}
	}
}
