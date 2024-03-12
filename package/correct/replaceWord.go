package correct

import (
	"cms/constant"
	"cms/db/models"
	"fmt"
	"slices"
	"strings"
)

// 文字列置換（not WCAG）
func ReplaceWord(sentence string, status *CorrectStatus, skipIds []int, accessibility models.Accessibility) {
	words, err := models.GetReplaceableCorrectWordListById(accessibility.Id)
	if err != nil {
		status.SetLevel(constant.CORRECT_NG)
		status.SetMessage("")
		return
	}

	ngReplace := status.GetReplace()
	ngReplaceSentence := sentence
	ngTooltipReplaceSentence := sentence
	warnReplace := status.GetReplace()
	warnTooltipReplaceSentence := sentence
	warnReplaceSentence := sentence

	for _, word := range words {
		from := word.WordFrom
		to := word.WordTo
		level := word.Level
		index := strings.Index(sentence, from)

		if index != -1 && level != constant.CORRECT_NO_CHECK {
			if level == constant.CORRECT_NG {
				status.SetLevel(constant.CORRECT_NG)
				status.SetMessage(accessibility.Message)
				ngReplace = append(ngReplace, Replace{
					From: from,
					To:   to,
				})
			} else if level == constant.CORRECT_WARNING {
				// 警告OK or NGが混ざったならスキップ
				if slices.Contains(skipIds, constant.REPLACE_WORD) || status.GetLevel() == constant.CORRECT_NG {
					continue
				}

				status.SetLevel(constant.CORRECT_WARNING)
				status.SetMessage(accessibility.Message)
				warnReplace = append(warnReplace, Replace{
					From: from,
					To:   to,
				})
			}

			replace := fmt.Sprintf(
				"<span class='correct-tooltip correct-tooltip-%d' data-tooltip='%s'>%s</span>",
				status.GetLevel(),
				status.GetMessage(),
				from,
			)

			if status.GetLevel() == constant.CORRECT_NG {
				status.SetReplace(ngReplace)
				ngTooltipReplaceSentence = strings.Replace(ngTooltipReplaceSentence, from, replace, -1)
				ngReplaceSentence = strings.Replace(ngReplaceSentence, from, to, -1)
			} else if status.GetLevel() == constant.CORRECT_WARNING {
				status.SetReplace(warnReplace)
				warnTooltipReplaceSentence = strings.Replace(warnTooltipReplaceSentence, from, replace, -1)
				warnReplaceSentence = strings.Replace(warnReplaceSentence, from, to, -1)
			}
		}
	}

	if status.GetLevel() == constant.CORRECT_NG {
		status.SetSentence(ngTooltipReplaceSentence)
		status.SetReplaceSentence(ngReplaceSentence)
	} else if status.GetLevel() == constant.CORRECT_WARNING {
		status.SetSentence(warnTooltipReplaceSentence)
		status.SetReplaceSentence(warnReplaceSentence)
	}
}
