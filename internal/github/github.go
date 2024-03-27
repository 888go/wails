package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"sort"
	"strings"

	"github.com/charmbracelet/glamour"
)


// ff:
// noColour:
// tagVersion:
func GetReleaseNotes(tagVersion string, noColour bool) string {
	resp, err := http.Get("https://api.github.com/repos/wailsapp/wails/releases/tags/" + url.PathEscape(tagVersion))
	if err != nil {
		return "Unable to retrieve release notes. Please check your network connection"
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Unable to retrieve release notes. Please check your network connection"
	}

	data := map[string]interface{}{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "Unable to retrieve release notes. Please check your network connection"
	}

	if data["body"] == nil {
		return "No release notes found"
	}

	result := "# Release Notes for " + tagVersion + "\n" + data["body"].(string)
	var renderer *glamour.TermRenderer

	var termRendererOpts []glamour.TermRendererOption

	if runtime.GOOS == "windows" || noColour {
		termRendererOpts = append(termRendererOpts, glamour.WithStyles(glamour.NoTTYStyleConfig))
	} else {
		termRendererOpts = append(termRendererOpts, glamour.WithAutoStyle())
	}

	renderer, err = glamour.NewTermRenderer(termRendererOpts...)
	if err != nil {
		return result
	}
	result, err = renderer.Render(result)
	if err != nil {
		return err.Error()
	}
	return result
}

// GetVersionTags 获取Wails仓库上的标签列表
// 返回一个按降序排序的标签列表

// ff:
func GetVersionTags() ([]*SemanticVersion, error) {
	result := []*SemanticVersion{}
	var err error

	resp, err := http.Get("https://api.github.com/repos/wailsapp/wails/tags")
	if err != nil {
		return result, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	data := []map[string]interface{}{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return result, err
	}

	// 将标签数据转换为Version结构体
	for _, tag := range data {
		version := tag["name"].(string)
		if !strings.HasPrefix(version, "v2") {
			continue
		}
		semver, err := NewSemanticVersion(version)
		if err != nil {
			return result, err
		}
		result = append(result, semver)
	}

	// Reverse Sort
	sort.Sort(sort.Reverse(SemverCollection(result)))

	return result, err
}

// GetLatestStableRelease 从 GitHub 获取最新的稳定版本

// ff:
// err:
// result:
func GetLatestStableRelease() (result *SemanticVersion, err error) {
	tags, err := GetVersionTags()
	if err != nil {
		return nil, err
	}

	for _, tag := range tags {
		if tag.IsRelease() {
			return tag, nil
		}
	}

	return nil, fmt.Errorf("no release tag found")
}

// GetLatestPreRelease 从GitHub获取最新的预发布版本

// ff:
// err:
// result:
func GetLatestPreRelease() (result *SemanticVersion, err error) {
	tags, err := GetVersionTags()
	if err != nil {
		return nil, err
	}

	for _, tag := range tags {
		if tag.IsPreRelease() {
			return tag, nil
		}
	}

	return nil, fmt.Errorf("no prerelease tag found")
}

// IsValidTag 判断给定的字符串是否为有效的标签并返回布尔值

// ff:
// tagVersion:
func IsValidTag(tagVersion string) (bool, error) {
	if tagVersion[0] == 'v' {
		tagVersion = tagVersion[1:]
	}
	tags, err := GetVersionTags()
	if err != nil {
		return false, err
	}

	for _, tag := range tags {
		if tag.String() == tagVersion {
			return true, nil
		}
	}
	return false, nil
}
