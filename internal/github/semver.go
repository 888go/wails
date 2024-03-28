package github

import (
	"fmt"

	"github.com/Masterminds/semver"
)

// SemanticVersion 是一个结构体，包含语义化版本信息
type SemanticVersion struct {
	Version *semver.Version
}

// NewSemanticVersion 根据给定的版本字符串创建一个新的 SemanticVersion 对象

// ff:
// version:
func NewSemanticVersion(version string) (*SemanticVersion, error) {
	semverVersion, err := semver.NewVersion(version)
	if err != nil {
		return nil, err
	}
	return &SemanticVersion{
		Version: semverVersion,
	}, nil
}

// IsRelease 返回 true，如果它是一个发布版本

// ff:
func (s *SemanticVersion) IsRelease() bool {
	// Limit to v2
	if s.Version.Major() != 2 {
		return false
	}
	return len(s.Version.Prerelease()) == 0 && len(s.Version.Metadata()) == 0
}

// IsPreRelease 判断是否为预发布版本，如果是则返回 true

// ff:
func (s *SemanticVersion) IsPreRelease() bool {
	// Limit to v1
	if s.Version.Major() != 2 {
		return false
	}
	return len(s.Version.Prerelease()) > 0
}


// ff:
func (s *SemanticVersion) String() string {
	return s.Version.String()
}

// IsGreaterThan 返回一个布尔值，表示如果当前版本大于给定版本，则返回true

// ff:
// version:
func (s *SemanticVersion) IsGreaterThan(version *SemanticVersion) (bool, error) {
	// Set up new constraint
	constraint, err := semver.NewConstraint("> " + version.Version.String())
	if err != nil {
		return false, err
	}

	// 检查期望值是否大于请求值
	success, msgs := constraint.Validate(s.Version)
	if !success {
		return false, msgs[0]
	}
	return true, nil
}

// IsGreaterThanOrEqual 返回一个布尔值，若当前版本大于或等于给定版本，则返回true

// ff:
// version:
func (s *SemanticVersion) IsGreaterThanOrEqual(version *SemanticVersion) (bool, error) {
	// Set up new constraint
	constraint, err := semver.NewConstraint(">= " + version.Version.String())
	if err != nil {
		return false, err
	}

	// 检查期望值是否大于请求值
	success, msgs := constraint.Validate(s.Version)
	if !success {
		return false, msgs[0]
	}
	return true, nil
}

// MainVersion 函数返回任何包含主版本+预发布版本+元数据的版本号中的主版本部分
// 例如：MainVersion("1.2.3-pre") => "1.2.3"

// ff:
func (s *SemanticVersion) MainVersion() *SemanticVersion {
	mainVersion := fmt.Sprintf("%d.%d.%d", s.Version.Major(), s.Version.Minor(), s.Version.Patch())
	result, _ := NewSemanticVersion(mainVersion)
	return result
}

// SemverCollection 是 SemanticVersion 对象的集合
type SemverCollection []*SemanticVersion

// Len 返回一个集合的长度。即切片中 Version 实例的数量。

// ff:
func (c SemverCollection) Len() int {
	return len(c)
}

// Less 是为了满足 sort 接口的需求，以便在切片上比较两个 Version 对象。它用于检查一个版本是否小于另一个版本。

// ff:
// j:
// i:
func (c SemverCollection) Less(i, j int) bool {
	return c[i].Version.LessThan(c[j].Version)
}

// Swap 是为了满足 sort 接口的需求，用于在切片中交换两个不同位置的 Version 对象。

// ff:
// j:
// i:
func (c SemverCollection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
