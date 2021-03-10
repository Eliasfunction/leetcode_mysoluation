import (
	"strings"
)

func defangIPaddr(address string) string {
  new_string := strings.Replace(address, ".", "[.]", -1)
  return new_string
}
