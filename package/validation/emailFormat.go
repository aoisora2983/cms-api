package validation

import (
	"errors"
	"regexp"
)

var emailPattern = `^(?i:[^ @"<>]+|".*")@(?i:[a-z1-9.])+.(?i:[a-z])+$`
var emailReg = regexp.MustCompile(emailPattern)

func EmailFormat(value interface{}) error {
	email, _ := value.(string)

	if len(emailReg.FindAllString(email, -1)) == 0 {
		return errors.New("メールアドレスの形式で入力してください。")
	}
	return nil
}
