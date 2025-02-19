package packages

import (
	"github.com/package-url/packageurl-go"
)

var OS = []string{"alma", "alpine", "amazon", "arch", "cbl-mariner", "centos", "chainguard", "debian", "fedora", "oracle", "redhat", "rocky", "suse", "ubuntu", "wolfi"}

func IndexFromInstance(instance packageurl.PackageURL) string {
	for _, osName := range OS {
		if instance.Namespace == osName {
			return osName
		}
	}
	return instance.Type
}

func IsOS(instance packageurl.PackageURL) bool {
	for _, osName := range OS {
		if instance.Namespace == osName {
			return true
		}
	}
	return false
}
