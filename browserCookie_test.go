package browsercookie

import (
	"testing"
	"fmt"
	"time"
	"os/user"
)

func TestReadChromeCookies(t *testing.T) {
	cookieJar, _ := LoadCookieJarFromChrome("")
	fmt.Println("CookieJar=", cookieJar)

	usr, _ := user.Current()
	cookiesFile := fmt.Sprintf("%s/Library/Application Support/Google/Chrome/Default/Cookies", usr.HomeDir)
	cookies, _ := ReadChromeCookies(cookiesFile, "", "", time.Time{})
	fmt.Println("Cookies=", cookies)
}
