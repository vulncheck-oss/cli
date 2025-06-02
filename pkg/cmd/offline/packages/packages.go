package packages

import (
	"github.com/package-url/packageurl-go"
)

var OS = []string{"alma", "alpine", "amazon", "arch", "cbl-mariner", "centos", "chainguard", "debian", "fedora", "oracle", "redhat", "rocky", "suse", "ubuntu", "wolfi"}
var OSSupported = []string{"alpine", "rocky", "debian"}

func IndexFromName(name string) string {
	for _, osName := range OSSupported {
		if name == osName {
			return osName + "-purls"
		}
	}
	return name
}

func IndexFromInstance(instance packageurl.PackageURL) string {
	for _, osName := range OSSupported {
		if instance.Namespace == osName {
			return osName + "-purls"
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

func ISOSSupported(instance packageurl.PackageURL) bool {
	for _, osName := range OSSupported {
		if instance.Namespace == osName {
			return true
		}
	}
	return false
}
