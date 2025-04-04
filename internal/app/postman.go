package app

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadPostmanFile(filePath string) (*PostmanCollection, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var collection PostmanCollection
	if err := json.Unmarshal(data, &collection); err != nil {
		return nil, fmt.Errorf("failed to parse Postman JSON: %w", err)
	}

	return &collection, nil
}

func WritePostmanFile(filePath string, collection *PostmanCollection) error {
	data, err := json.MarshalIndent(collection, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal Postman JSON: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func ConvertInsomniaToPostman(insomnia *InsomniaExport) (*PostmanCollection, error) {
	collection := &PostmanCollection{
		Info: PostmanInfo{
			Name:        insomnia.Name,
			Description: "",
			Schema:      "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
			PostmanID:   generateUUID(),
		},
		Items: []PostmanItem{},
	}

	for _, resource := range insomnia.Collection {
		item := convertInsomniaResourceToPostmanItem(resource)
		collection.Items = append(collection.Items, item)
	}

	return collection, nil
}

func convertInsomniaResourceToPostmanItem(resource InsomniaResource) PostmanItem {
	item := PostmanItem{
		Name:        resource.Name,
		Description: "",
	}

	if len(resource.Children) > 0 {
		item.Items = make([]PostmanItem, 0, len(resource.Children))
		for _, child := range resource.Children {
			childItem := convertInsomniaResourceToPostmanItem(child)
			item.Items = append(item.Items, childItem)
		}
	} else if resource.URL != "" {
		request := &PostmanRequest{
			Method:      resource.Method,
			Description: "",
		}

		parsedURL := parseURL(resource.URL)
		request.URL = parsedURL

		if len(resource.Headers) > 0 {
			request.Headers = make([]PostmanHeader, 0, len(resource.Headers))
			for _, header := range resource.Headers {
				request.Headers = append(request.Headers, PostmanHeader{
					Key:      header.Name,
					Value:    header.Value,
					Disabled: header.Disabled,
				})
			}
		}

		if len(resource.Parameters) > 0 {
			if request.URL.Query == nil {
				request.URL.Query = make([]PostmanQuery, 0, len(resource.Parameters))
			}

			for _, param := range resource.Parameters {
				request.URL.Query = append(request.URL.Query, PostmanQuery{
					Key:      param.Name,
					Value:    param.Value,
					Disabled: param.Disabled,
				})
			}
		}

		if resource.Body != nil && resource.Body.Text != "" {
			request.Body = &PostmanBody{
				Mode: getBodyMode(resource.Body.MimeType),
				Raw:  resource.Body.Text,
			}

			if resource.Body.MimeType == "application/json" {
				request.Body.Options = map[string]interface{}{
					"raw": map[string]interface{}{
						"language": "json",
					},
				}
			}
		}

		item.Request = request
	}

	return item
}
