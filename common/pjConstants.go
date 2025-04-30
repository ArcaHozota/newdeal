package common

const NormalError string = "想定外のエラーが発生しました。管理員へご連絡ください。"

const AttrNameEntity string = "arawaseta"

const EmptyString string = ""

const HankakuPercentMark string = "%"

const AttrNameException string = "exception"

const AttrNameTorokuMsg string = "torokuMsg"

const AttrNamePageNo string = "pageNum"

const DefaultPageSize uint8 = 5

const JwtErrorMsg string = "トークンの有効期間過ぎでまたは無効です。"

const NeedLoginMsg string = "ログインしてください"

const LoginedMsg string = "ログイン成功!"

const LogoutMsg string = "ログアウトしました、必要な場合はもう一度ログインしてください。"

const InsertedMsg string = "追加済み"

const UpsertedMsg string = "更新または追加済み"

const UpdatedMsg string = "更新済み"

const NochangeMsg string = "変更なし"

const StudentError string = "当ユーザが存在しません。"

const PasswordError string = "パスワードが不一致になります。"

const LoginFormError string = "正しい形式で入力してください"

const DateLayout string = "2006-01-02"

const DateTimeLayout string = "2006-01-02 15:04:05.000000"

var StrangeArray = []string{
	"insert", "delete", "update", "create", "drop",
	"#", "$", "%", "&", "(", ")", "\"", "'", "@", ":", "select",
}
