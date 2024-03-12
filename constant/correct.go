package constant

// correct levelt
const (
	CORRECT_OK      = 0
	CORRECT_NG      = 1
	CORRECT_WARNING = 2
)

// correct kind
const (
	NON_TEXT     = 1
	REPLACE_WORD = 2
)

// 1.1.1
const (
	NO_ALT_ATTRIBUTE = "画像に代替文字が設定されていません。\nアイコンの場合は空文字を設定してください。"
	NO_ALT_TEXT      = "画像に代替文字が入力されていません。\nアイコンの場合は空文字のままで問題ありません。"
)

// 文字列置換（not WCAG）
const (
	REPLACE_WORD_NG      = "アクセシビリティに問題のある文字列があります。"
	REPLACE_WORD_WARNING = "アクセシビリティに問題のある文字列があります。"
)
