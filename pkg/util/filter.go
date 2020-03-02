
package util

import (
    "regexp"
)

func RegEmail(email string) bool {
    pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
    reg := regexp.MustCompile(pattern)
    return reg.MatchString(email)
}
