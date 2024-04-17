package notion_test

import (
	"my-tools/thanksfulness/notion"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func toTime(dateTime string) time.Time {
	t, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		return time.Now()
	}
	return t
}

func TestUnmarshalResponse(t *testing.T) {
	t.Parallel()
	// given
	input, err := os.ReadFile("./NotionAPIResponseTestData.json")
	if err != nil {
		t.Errorf("Error reading test data: %v", err)
	}
	// when
	actual, error := notion.UnmarshalResponse(input)
	// then
	assert.Nil(t, error)
	expectedResponse := &notion.APIResponse{
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
							PlainText: "2024/02/01",
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
							PlainText: "感謝永遠に",
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
	if !reflect.DeepEqual(actual, expectedResponse) {
		t.Errorf("Unexpected response: expected %v, got %v", expectedResponse, actual)
	}
}
