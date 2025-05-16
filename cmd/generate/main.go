package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	directoryUrl = "https://api.hubspot.com/public/api/spec/v1/specs"
	replacer     = strings.NewReplacer(" ", "_", "-", "_")
)

// schemaOpenAPI contains the OpenAPI schema definition.
type schemaOpenAPI struct {
	Info struct {
		Title   string `json:"title"`
		Version string `json:"version"`
	} `json:"info"`
}

// directory is the schema for the https://api.hubspot.com/public/api/spec/v1/specs page which lists all
// the publicly available APIs.
type directory struct {
	Results []struct {
		Name     string `json:"name"`
		Group    string `json:"group"`
		Versions []struct {
			Version   int    `json:"version"`
			Stage     string `json:"stage"`
			OpenApi   string `json:"openApi"`
		} `json:"versions"`
	} `json:"results"`
}

// retrieveSchema downloads a schema from the given url.
// It returns the schema text as a string.
func retrieveSchema(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	schema, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(schema), nil
}

// saveSchema saves the schema to a file.
// It returns a path to the saved file.
func saveSchema(group string, schema string) (string, error) {
	// Create schema directory if it doesn't exist
	err := os.MkdirAll("./schema", 0755)
	if err != nil {
		return "", err
	}

	filename := "./schema/" + group + ".json"
	err = ioutil.WriteFile(filename, []byte(schema), 0644)

	if err != nil {
		return "", err
	}

	return filename, nil
}

// versionFromSchema retrieves the OpenAPI schema version from the input JSON string.
// It returns the API version.
func versionFromSchema(input string) (string, error) {
	var schema schemaOpenAPI
	err := json.Unmarshal([]byte(input), &schema)

	if err != nil {
		return "", err
	}

	return schema.Info.Version, nil
}

func main() {
	// Add user agent to prevent 403 errors
	client := &http.Client{}
	req, err := http.NewRequest("GET", directoryUrl, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var r directory
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		panic(err)
	}
	err = os.Setenv("GIT_USER_ID", "clarkmcc")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("GIT_REPO_ID", "go-hubspot")
	if err != nil {
		panic(err)
	}
	generator, err := exec.LookPath("openapi-generator")
	if err != nil {
		panic(err)
	}

	// Process all APIs
	var successCount, skipCount, errorCount int


	for _, api := range r.Results {
		// Find the latest stable version for each API
		var latestVersion *struct {
			Version int    `json:"version"`
			Stage   string `json:"stage"`
			OpenApi string `json:"openApi"`
		}

		fmt.Printf("Processing API: %s (Group: %s)\n", api.Name, api.Group)

		for i, version := range api.Versions {
			if latestVersion == nil || version.Stage == "LATEST" {
				latestVersion = &api.Versions[i]
			}
		}

		if latestVersion == nil {
			fmt.Printf("  Skipping %s/%s (no available version)\n", api.Group, api.Name)
			skipCount++
			continue
		}

		fmt.Printf("  Using version: v%d (Stage: %s)\n", latestVersion.Version, latestVersion.Stage)

		// Fetch the OpenAPI spec from the OpenApi URL
		schema, err := retrieveSchema(latestVersion.OpenApi)

		if err != nil {
			fmt.Printf("  Error retrieving schema for %s/%s: %v\n", api.Group, api.Name, err)
			errorCount++
			continue
		}

		// Save the schema to a file
		name := api.Name
		filename, err := saveSchema(name, schema)

		if err != nil {
			fmt.Printf("  Error saving schema for %s/%s: %v\n", api.Group, api.Name, err)
			errorCount++
			continue
		}

		version, err := versionFromSchema(schema)

		if err != nil {
			fmt.Printf("  Error extracting version for %s/%s: %v\n", api.Group, api.Name, err)
			errorCount++
			continue
		}

		// Convert API name to safe package name
		name = replacer.Replace(strings.ToLower(name))
		fmt.Printf("  Generating API client for %s/%s (version %s)\n", api.Group, name, version)

		output, err := exec.Command(generator, "generate",
			"-i", filename,
			"-g", "go",
			"--package-name", name,
			"--additional-properties=isGoSubmodule=false",
			"--skip-validate-spec",
			"-o", "./generated/"+version+"/"+name,
		).CombinedOutput()
		if err != nil {
			fmt.Printf("  Error generating client for %s/%s: %v\n", api.Group, api.Name, err)
			errorCount++
			continue
		}

		fmt.Printf("  Output: %s\n", string(output))

		fmt.Printf("  Successfully generated API client for %s/%s (version %s)\n", api.Group, name, version)
		successCount++
	}

	fmt.Printf("\nGeneration Summary:\n")
	fmt.Printf("  Successfully generated: %d APIs\n", successCount)
	fmt.Printf("  Skipped: %d APIs\n", skipCount)
	fmt.Printf("  Failed: %d APIs\n", errorCount)

	fmt.Println("updating generated files")
	err = filepath.Walk("generated", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		// API files check for the presence of an API key in the request context using type
		// assertion. Because we generate each module separately, each module has its own
		// type for storing the API key in the context (contacts.APIKey vs deals.APIKey). In
		// order to allow us to share the same authorizer across packages, we'll do a find
		// and replace on all the package-specific API key structs and replace them with our
		// own global structs.
		if strings.HasSuffix(info.Name(), ".go") {
			// Open the file, read it, then close it
			var b []byte
			err = func() error {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()
				b, err = ioutil.ReadAll(file)
				if err != nil {
					return err
				}
				return nil
			}()
			if err != nil {
				return err
			}

			// Replace any imports of GIT_USER_ID/GIT_REPO_ID with our repo
			if bytes.Contains(b, []byte("github.com/GIT_USER_ID/GIT_REPO_ID")) {
				b = bytes.ReplaceAll(b, []byte("github.com/GIT_USER_ID/GIT_REPO_ID"), []byte("github.com/clarkmcc/go-hubspot"))
				fmt.Printf("  Fixed imports in %s\n", path)
			}

			for _, method := range findMethods {
				if !bytes.Contains(b, method) {
					continue
				}

				// Replace generated API key auth with custom API key auth
				b = bytes.ReplaceAll(b, method, replaceWith)

				// Add imports for authorization package
				idx := bytes.Index(b, []byte("\"net/url\""))
				if idx >= 0 {
					b = bytes.Join([][]byte{b[:idx], []byte("\t\"github.com/clarkmcc/go-hubspot\""), b[idx:]}, []byte("\n"))
				}
				break
			}

			// Write back the modified file
			err = os.WriteFile(path, b, 0666)
			if err != nil {
				return err
			}
		}
		
		// The client configuration gets generated with a go module for each generated
		// client. We'll go through and delete each submodule so that we can take advantage
		// of the root parent module.
		if info.Name() != "go.mod" && info.Name() != "go.sum" && info.Name() != "git_push.sh" {
			return nil
		}
		return os.Remove(path)
	})
	if err != nil {
		panic(err)
	}
}

var findPrivateAppsAuth = []byte("if r.ctx != nil {\n\t\t// API Key Authentication\n\t\tif auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {\n\t\t\tif apiKey, ok := auth[\"private_apps\"]; ok {\n\t\t\t\tvar key string\n\t\t\t\tif apiKey.Prefix != \"\" {\n\t\t\t\t\tkey = apiKey.Prefix + \" \" + apiKey.Key\n\t\t\t\t} else {\n\t\t\t\t\tkey = apiKey.Key\n\t\t\t\t}\n\t\t\t\tlocalVarHeaderParams[\"private-app\"] = key\n\t\t\t}\n\t\t}\n\t}")
var findPrivateAppsLegacyAuth = []byte("if r.ctx != nil {\n\t\t// API Key Authentication\n\t\tif auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {\n\t\t\tif apiKey, ok := auth[\"private_apps_legacy\"]; ok {\n\t\t\t\tvar key string\n\t\t\t\tif apiKey.Prefix != \"\" {\n\t\t\t\t\tkey = apiKey.Prefix + \" \" + apiKey.Key\n\t\t\t\t} else {\n\t\t\t\t\tkey = apiKey.Key\n\t\t\t\t}\n\t\t\t\tlocalVarHeaderParams[\"private-app-legacy\"] = key\n\t\t\t}\n\t\t}\n\t}")
var findDeveloperHapiKeyAuth = []byte("if r.ctx != nil {\n\t\t// API Key Authentication\n\t\tif auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {\n\t\t\tif apiKey, ok := auth[\"developer_hapikey\"]; ok {\n\t\t\t\tvar key string\n\t\t\t\tif apiKey.Prefix != \"\" {\n\t\t\t\t\tkey = apiKey.Prefix + \" \" + apiKey.Key\n\t\t\t\t} else {\n\t\t\t\t\tkey = apiKey.Key\n\t\t\t\t}\n\t\t\t\tlocalVarQueryParams.Add(\"hapikey\", key)\n\t\t\t}\n\t\t}\n\t}")
var replaceWith = []byte("if r.ctx != nil {\n\t\t// API Key Authentication\n\t\tif auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {\n\t\t\tauth.Apply(hubspot.AuthorizationRequest{\n\t\t\t\tQueryParams: localVarQueryParams,\n\t\t\t\tFormParams:  localVarFormParams,\n\t\t\t\tHeaders:     localVarHeaderParams,\n\t\t\t})\n\t\t}\n\t}")

var findMethods = [][]byte{
	findPrivateAppsAuth,
	findPrivateAppsLegacyAuth,
	findDeveloperHapiKeyAuth,
}
