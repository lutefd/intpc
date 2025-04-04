package app

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

func (c *Converter) DetectFormat() (string, error) {
	ext := filepath.Ext(c.InputFile)

	if ext == ".yaml" || ext == ".yml" {
		data, err := os.ReadFile(c.InputFile)
		if err != nil {
			return "", fmt.Errorf("failed to read file: %w", err)
		}

		if strings.Contains(string(data), "type: collection.insomnia") {
			return "insomnia", nil
		}
	} else if ext == ".json" {
		data, err := os.ReadFile(c.InputFile)
		if err != nil {
			return "", fmt.Errorf("failed to read file: %w", err)
		}

		var obj map[string]interface{}
		if err := json.Unmarshal(data, &obj); err == nil {
			if info, ok := obj["info"].(map[string]interface{}); ok {
				if schema, ok := info["schema"].(string); ok && strings.Contains(schema, "postman") {
					return "postman", nil
				}
			}
		}
	}

	data, err := os.ReadFile(c.InputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	var yamlObj map[string]interface{}
	if err := yaml.Unmarshal(data, &yamlObj); err == nil {
		if _, ok := yamlObj["type"]; ok {
			if typ, ok := yamlObj["type"].(string); ok && strings.Contains(typ, "insomnia") {
				return "insomnia", nil
			}
		}
	}

	var jsonObj map[string]interface{}
	if err := json.Unmarshal(data, &jsonObj); err == nil {
		if info, ok := jsonObj["info"].(map[string]interface{}); ok {
			if _, ok := info["schema"]; ok {
				return "postman", nil
			}
		}
	}

	return "", fmt.Errorf("unable to determine format of input file")
}

func parseURL(urlStr string) PostmanURL {
	postmanURL := PostmanURL{
		Raw: urlStr,
	}

	if strings.HasPrefix(urlStr, "http") {
		parts := strings.Split(urlStr, "://")
		if len(parts) > 1 {
			postmanURL.Protocol = parts[0]

			hostAndPath := strings.SplitN(parts[1], "/", 2)
			if len(hostAndPath) > 0 {
				postmanURL.Host = strings.Split(hostAndPath[0], ".")

				if len(hostAndPath) > 1 {
					pathStr := "/" + hostAndPath[1]
					postmanURL.Path = strings.Split(strings.Trim(pathStr, "/"), "/")
				}
			}
		}
	}

	return postmanURL
}

func getBodyMode(mimeType string) string {
	switch mimeType {
	case "application/json":
		return "raw"
	case "application/x-www-form-urlencoded":
		return "urlencoded"
	case "multipart/form-data":
		return "formdata"
	default:
		return "raw"
	}
}

func currentTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func generateShortUUID() string {
	return strings.Replace(generateUUID(), "-", "", -1)[:24]
}
