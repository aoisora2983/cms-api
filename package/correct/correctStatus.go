package correct

import "cms/constant"

type Replace struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type CorrectStatus struct {
	Kind            int
	Level           int
	Message         string
	Sentence        string
	Replace         []Replace
	ReplaceSentence string
}

func NewCorrectStatus() *CorrectStatus {
	return &CorrectStatus{
		Kind:            constant.NON_TEXT,
		Level:           constant.CORRECT_NO_CHECK,
		Message:         "",
		Sentence:        "",
		Replace:         []Replace{},
		ReplaceSentence: "",
	}
}

func (c *CorrectStatus) Response() map[string]interface{} {
	return map[string]interface{}{
		"kind":             c.GetKind(),
		"level":            c.GetLevel(),
		"message":          c.GetMessage(),
		"sentence":         c.GetSentence(),
		"replace":          c.GetReplace(),
		"replace_sentence": c.GetReplaceSentence(),
	}
}

func (c *CorrectStatus) GetKind() int {
	return c.Kind
}

func (c *CorrectStatus) SetKind(kind int) {
	c.Kind = kind
}

func (c *CorrectStatus) GetLevel() int {
	return c.Level
}

func (c *CorrectStatus) SetLevel(level int) {
	c.Level = level
}

func (c *CorrectStatus) GetSentence() string {
	return c.Sentence
}

func (c *CorrectStatus) SetSentence(sentence string) {
	c.Sentence = sentence
}

func (c *CorrectStatus) GetMessage() string {
	return c.Message
}

func (c *CorrectStatus) SetMessage(message string) {
	c.Message = message
}

func (c *CorrectStatus) GetReplace() []Replace {
	return c.Replace
}

func (c *CorrectStatus) SetReplace(replace []Replace) {
	c.Replace = replace
}

func (c *CorrectStatus) GetReplaceSentence() string {
	return c.ReplaceSentence
}

func (c *CorrectStatus) SetReplaceSentence(replaceSentence string) {
	c.ReplaceSentence = replaceSentence
}
