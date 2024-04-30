package slack

import (
	"log"
	"my-tools/thanksfulness/notion"
	"os"

	"github.com/slack-go/slack"
)

// usecase ロジック
// Slack へ感謝リストを送信する
func PostThanksList(thanksList *notion.ThanksList) error {
	client := createClient(os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN"))

	// リクエストするデータを生成する
	attachment := MakeAttachment(thanksList)
	msgOptions := makeMsgOptions(attachment)

	err := Request(client, msgOptions)
	if err != nil {
		log.Println("Failed to request Slack API")
		return err
	}
	return nil
}

func createClient(token string) *slack.Client {
	return slack.New(token)
}

// slack の chat.postMessage webAPI にリクエストするデータを生成する
func MakeAttachment(thanksList *notion.ThanksList) *slack.Attachment {
	return &slack.Attachment{
		Fallback: "Thanksfulness List",
		Color:    "#ffb6c1",
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{
				slack.NewHeaderBlock(&slack.TextBlockObject{
					Type: slack.PlainTextType,
					Text: ":heart_hands:感謝！:heart_hands:",
				}),
				slack.NewContextBlock("5minStaring", &slack.TextBlockObject{
					Type: slack.MarkdownType,
					Text: "5分間じっくり噛み締めよう:eyes:",
				}),
				slack.NewDividerBlock(),
				slack.NewRichTextBlock(
					"ThanksList",
					slack.NewRichTextList(slack.RTEListBullet, 0, MakeRichTextElements(thanksList)...,
					)),
				slack.NewContextBlock("EditedAt", &slack.TextBlockObject{
					Type: slack.MarkdownType,
					Text: "Edited on " + thanksList.EditedAt,
				}),
			},
		},
	}
}

// thanksList.Items の数だけ rich_text_section を作成する
//
// interface slack.RichTextElement を返さないと呼び出し側でうまく使えないため
// 実際に返している型 []slack.RichTextSection とは異なる型定義を関数に付与している
func MakeRichTextElements(thanksList *notion.ThanksList) (richTextElements []slack.RichTextElement) {
	for _, v := range thanksList.Items {
		richTextElements = append(richTextElements, slack.RichTextSection{
			Type: slack.RTESection,
			Elements: []slack.RichTextSectionElement{
				slack.NewRichTextSectionTextElement(v, nil),
			},
		})
	}
	return richTextElements
}

func makeMsgOptions(attachment *slack.Attachment) []slack.MsgOption {
	msgOption := slack.MsgOptionAttachments(*attachment)
	msgOptionIconEmoji := slack.MsgOptionIconEmoji(":gratitude-thank-you:")
	msgOptionUsername := slack.MsgOptionUsername("Thanksfulness")

	return []slack.MsgOption{
		msgOption, msgOptionIconEmoji, msgOptionUsername,
	}
}

// slack の chat.postMessage webAPI にリクエストしレスポンスを返す
func Request(c *slack.Client, msgOptions []slack.MsgOption) error {
	channelID, timestamp, err := c.PostMessage("C03EE1N8VCL", msgOptions...)
	if err != nil {
		log.Printf("Succeeded to post Slack message to channel (%v) at %v", channelID, timestamp)
		return err
	}
	return nil
}
