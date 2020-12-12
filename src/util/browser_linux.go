package util

func GetBrowserCmd(url string) (name string, arg []string) {
	return "xdg-open", []string{url}
}
