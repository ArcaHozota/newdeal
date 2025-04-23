package common

const NormalError string = "想定外のエラーが発生しました。管理員へご連絡ください。"

const AttrNameEntity string = "arawaseta"

const EmptyString string = ""

const HankakuPercentMark string = "%"

const AttrNameException string = "exception"

const DefaultPageSize uint8 = 5

const NeedLoginMsg string = "ログインしてください"

const LoginedMsg string = "ログイン成功!"

const StudentError string = "当ユーザが存在しません。"

const PasswordError string = "パスワードが不一致になります。"

const LoginFormError string = "正しい形式で入力してください"

var StrangeArray = []string{
	"insert", "delete", "update", "create", "drop",
	"#", "$", "%", "&", "(", ")", "\"", "'", "@", ":", "select",
}
