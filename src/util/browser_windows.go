package util

func GetBrowserCmd(url string) (name string, arg []string) {
	return "cmd", []string{"/c", "start", url}
}
