package notion_test

import (
	"my-tools/thanksfulness/notion"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	t.Parallel()
	// given
	input := defaultInput()
	// when
	actual, _ := notion.Convert(input)
	// then
	assert.Equal(t, notion.ThanksList{
		EditedAt: "2024/01/01",
		Items:    []string{"ありがとう", "Merci", "謝謝"},
	}, *actual)
}

func defaultInput() *notion.APIResponse {
	return &notion.APIResponse{
		Object: "list",
		Results: []notion.APIResult{
			{
				Object:         "block",
				Id:             "00000000-0000-0000-0000-000011112222",
				Parent:         notion.APIParent{Type: "block_id", BlockId: "00000001-0000-0000-0000-000011112222"},
				CreatedTime:    toTime("2023-01-01T09:00:00.000Z"),
				LastEditedTime: toTime("2024-02-01T03:00:00.000Z"),
				CreatedBy:      notion.APIUser{Object: "user", Id: "10000000-1111-0000-0000-000011112222"},
				LastEditedBy:   notion.APIUser{Object: "user", Id: "10000000-1111-0000-0000-000011112222"},
				HasChildren:    false,
				Archived:       false,
				InTrash:        false,
				Type:           "paragraph",
				Paragraph: &notion.APIRichText{
					RichText: []notion.APIPlainText{
						{
							Type:      "text",
							PlainText: "2024/01/01",
						},
					},
				},
			},
			{
				Object:         "block",
				Id:             "10000000-0000-0000-0000-000011112222",
				Parent:         notion.APIParent{Type: "block_id", BlockId: "00000001-0000-0000-0000-000011112222"},
				CreatedTime:    toTime("2023-01-01T09:00:00.000Z"),
				LastEditedTime: toTime("2024-02-01T03:00:00.000Z"),
				CreatedBy:      notion.APIUser{Object: "user", Id: "10000000-1111-0000-0000-000011112222"},
				LastEditedBy:   notion.APIUser{Object: "user", Id: "10000000-1111-0000-0000-000011112222"},
				HasChildren:    false,
				Archived:       false,
				InTrash:        false,
				Type:           "bulleted_list_item",
				BulletedListItem: &notion.APIRichText{
					RichText: []notion.APIPlainText{
						{
							Type:      "text",
							PlainText: "ありがとう",
						},
						{
							Type:      "text",
							PlainText: "Merci",
						},
						{
							Type:      "text",
							PlainText: "謝謝",
						},
					},
				},
			},
		},
		NextCursor: false, // nil はboolのゼロ値であるfalseにマッピングされる
		HasMore:    false,
		Type:       "block",
		Block:      make(map[string]interface{}),
		RequestId:  "30000000-0000-0000-0000-000011112222",
	}
}
