package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"

	"github.com/888go/wails/internal/s"
)

const versionFile = "../../cmd/wails/internal/version.txt"

func checkError(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

// TODO: 这部分可以替换为 "https://github.com/coreos/go-semver/blob/main/semver/semver.go" 的内容
func updateVersion() string {
	currentVersionData, err := os.ReadFile(versionFile)
	checkError(err)
	currentVersion := string(currentVersionData)
	vsplit := strings.Split(currentVersion, ".")
	minorVersion, err := strconv.Atoi(vsplit[len(vsplit)-1])
	checkError(err)
	minorVersion++
	vsplit[len(vsplit)-1] = strconv.Itoa(minorVersion)
	newVersion := strings.Join(vsplit, ".")
	err = os.WriteFile(versionFile, []byte(newVersion), 0o755)
	checkError(err)
	return newVersion
}

func runCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	checkError(err)
}


// ff:
// newVersion:
// currentVersion:
func IsPointRelease(currentVersion string, newVersion string) bool {
	// 版本的前n-1部分应保持相同
	if currentVersion[:len(currentVersion)-2] != newVersion[:len(newVersion)-2] {
		return false
	}
	// 在字符串的最后一个点处分割
	currentVersionSplit := strings.Split(currentVersion, ".")
	newVersionSplit := strings.Split(newVersion, ".")
// 比较版本
// 如果版本的最后一部分相同，那么这是一个小版本（点发布）
	currentMinor := lo.Must(strconv.Atoi(currentVersionSplit[len(currentVersionSplit)-1]))
	newMinor := lo.Must(strconv.Atoi(newVersionSplit[len(newVersionSplit)-1]))
	return newMinor == currentMinor+1
}

func main() {
	var newVersion string
	var isPointRelease bool
	if len(os.Args) > 1 {
		newVersion = os.Args[1]
		currentVersion, err := os.ReadFile(versionFile)
		checkError(err)
		err = os.WriteFile(versionFile, []byte(newVersion), 0o755)
		checkError(err)
		isPointRelease = IsPointRelease(string(currentVersion), newVersion)
	} else {
		newVersion = updateVersion()
	}

	// Update ChangeLog
	s.CD("../../../website")

	// 读取 `src/pages/changelog.mdx` 文件
	changelogData, err := os.ReadFile("src/pages/changelog.mdx")
	checkError(err)
	changelog := string(changelogData)
	// 在包含`## [未发布]`的行上进行分割
	changelogSplit := strings.Split(changelog, "## [Unreleased]")
	// 获取今天的日期，格式为YYYY-MM-DD
	today := time.Now().Format("2006-01-02")
	// 将新版本添加到变更日志的顶部
	newChangelog := changelogSplit[0] + "## [Unreleased]\n\n## " + newVersion + " - " + today + changelogSplit[1]
	// 将变更日志写回
	err = os.WriteFile("src/pages/changelog.mdx", []byte(newChangelog), 0o755)
	checkError(err)

	if !isPointRelease {
		runCommand("npx", "-y", "pnpm", "install")

		s.ECHO("Generating new Docs for version: " + newVersion)

		runCommand("npx", "pnpm", "run", "docusaurus", "docs:version", newVersion)

		runCommand("npx", "pnpm", "run", "write-translations")

		// 加载版本列表
		versionsData, err := os.ReadFile("versions.json")
		checkError(err)
		var versions []string
		err = json.Unmarshal(versionsData, &versions)
		checkError(err)
		oldestVersion := versions[len(versions)-1]
		s.ECHO(oldestVersion)
		versions = versions[0 : len(versions)-1]
		newVersions, err := json.Marshal(&versions)
		checkError(err)
		err = os.WriteFile("versions.json", newVersions, 0o755)
		checkError(err)

		s.ECHO("Removing old version: " + oldestVersion)
		s.CD("versioned_docs")
		s.RMDIR("version-" + oldestVersion)
		s.CD("../versioned_sidebars")
		s.RM("version-" + oldestVersion + "-sidebars.json")
		s.CD("..")

		runCommand("npx", "pnpm", "run", "build")
	}
}
