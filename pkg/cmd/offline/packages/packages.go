package packages

import (
	"github.com/package-url/packageurl-go"
)

func IndexFromInstance(instance packageurl.PackageURL) string {
	switch instance.Namespace {
	case "alma":
		return "alma"
	case "alpine":
		return "alpine"
	default:
		return instance.Type
	}
}

func IsOS(instance packageurl.PackageURL) bool {
	switch instance.Namespace {
	case "wolfi":
		return true
	case "ubuntu":
		return true
	case "suse":
		return true
	case "rocky":
		return true
	case "redhat":
		return true
	case "oracle":
		return true
	case "fedora":
		return true
	case "debian":
		return true
	case "chainguard":
		return true
	case "centos":
		return true
	case "cbl-mariner":
		return true
	case "arch":
		return true
	case "alma":
		return true
	case "alpine":
		return true
	case "deb":
		return true
	case "rpm":
		return true
	default:
		return false
	}
}
