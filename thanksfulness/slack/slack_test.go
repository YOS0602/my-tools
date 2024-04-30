package slack_test

import (
	"encoding/json"
	"my-tools/thanksfulness/notion"
	target "my-tools/thanksfulness/slack"
	"testing"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

func TestMakeReqBody(t *testing.T) {
	// given
	input := &notion.ThanksList{
		EditedAt: "2020/01/01",
		Items:    []string{"item1", "item2"},
	}
	// when
	actual := target.MakeAttachment(input)
	// then
	assert.NotNil(t, actual)
	b, err := json.MarshalIndent(*actual, "", "  ")
	assert.Nil(t, err)
	j := string(b)
	assert.Contains(t, j, "2020/01/01")
	assert.Contains(t, j, "item1")
	assert.Contains(t, j, "item2")
}

func TestRichTextElement(t *testing.T) {
	// given
	input := &notion.ThanksList{
		EditedAt: "2020/01/01",
		Items:    []string{"item1", "item2", "item3", "item4", "item5"},
	}
	// when
	actual := target.MakeRichTextElements(input)
	// then
	// 関数の戻り値型である []slack.RichTextElement{} で初期化できないので []slack.RichTextSection を使用している
	expected := []slack.RichTextSection{
		{Type: slack.RTESection, Elements: []slack.RichTextSectionElement{
			slack.NewRichTextSectionTextElement("item1", nil),
		}},
		{Type: slack.RTESection, Elements: []slack.RichTextSectionElement{
			slack.NewRichTextSectionTextElement("item2", nil),
		}},
		{Type: slack.RTESection, Elements: []slack.RichTextSectionElement{
			slack.NewRichTextSectionTextElement("item3", nil),
		}},
		{Type: slack.RTESection, Elements: []slack.RichTextSectionElement{
			slack.NewRichTextSectionTextElement("item4", nil),
		}},
		{Type: slack.RTESection, Elements: []slack.RichTextSectionElement{
			slack.NewRichTextSectionTextElement("item5", nil),
		}},
	}
	assert.IsType(t, []slack.RichTextElement{}, actual)
	assert.Len(t, actual, 5)

	// expected []slack.RichTextSection
	// actual []slack.RichTextElement
	// 上記は暗黙変換不可で型違いエラーが出るので assert.Equal(t, expected, actual) はできない
	// そのためJSON変換した要素が一致するかを確認している
	expectedBytes, _ := json.Marshal(expected)
	actualBytes, _ := json.Marshal(actual)
	assert.JSONEq(t, string(expectedBytes), string(actualBytes))
}
