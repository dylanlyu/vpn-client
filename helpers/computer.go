package helpers
import "os"

func GetLocalName() string {
	host, err := os.Hostname()
	if err != nil {
		return ""
	} else {
		return host
	}

}
