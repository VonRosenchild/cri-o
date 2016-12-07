// +build !apparmor

package apparmor

// IsEnabled returns false, when build without apparmor build tag.
func IsEnabled() bool {
	return false
}

// LoadDefaultAppArmorProfile dose nothing, when build without apparmor build tag.
func LoadDefaultAppArmorProfile() {
}

// GetProfileNameFromPodAnnotations dose nothing, when build without apparmor build tag.
func GetProfileNameFromPodAnnotations(annotations map[string]string, containerName string) string {
	return ""
}
