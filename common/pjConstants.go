package common

const NormalError string = "想定外のエラーが発生しました。管理員へご連絡ください。"

const AttrNameEntity string = "arawaseta"

const EmptyString string = ""

const HankakuPercentMark string = "%"

const AttrNameException string = "exception"

const DefaultPageSize uint8 = 5

var StrangeArray = []string{
	"insert", "delete", "update", "create", "drop",
	"#", "$", "%", "&", "(", ")", "\"", "'", "@", ":", "select",
}
