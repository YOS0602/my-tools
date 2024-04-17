package notion_test

import (
	"my-tools/thanksfulness/notion"
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
	// given
	input := []byte(testJson)
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

var testJson = `{
  "object": "list",
  "results": [
    {
      "object": "block",
      "id": "00000000-0000-0000-0000-000011112222",
      "parent": {
        "type": "block_id",
        "block_id": "00000001-0000-0000-0000-000011112222"
      },
      "created_time": "2023-01-01T09:00:00.000Z",
      "last_edited_time": "2024-02-01T03:00:00.000Z",
      "created_by": {
        "object": "user",
        "id": "10000000-1111-0000-0000-000011112222"
      },
      "last_edited_by": {
        "object": "user",
        "id": "10000000-1111-0000-0000-000011112222"
      },
      "has_children": false,
      "archived": false,
      "in_trash": false,
      "type": "paragraph",
      "paragraph": {
        "rich_text": [
          {
            "type": "text",
            "text": { "content": "2024/02/01", "link": null },
            "annotations": {
              "bold": false,
              "italic": false,
              "strikethrough": false,
              "underline": false,
              "code": false,
              "color": "default"
            },
            "plain_text": "2024/02/01",
            "href": null
          }
        ],
        "color": "default"
      }
    }
  ],
  "next_cursor": null,
  "has_more": false,
  "type": "block",
  "block": {},
  "request_id": "30000000-0000-0000-0000-000011112222"
}`
