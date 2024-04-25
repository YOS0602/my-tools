package notion

import (
	"io"
	"log"
	"my-tools/thanksfulness/lib"
	"os"
	"strings"
)

func makeGetUrl(notionThanksfulnessBlockId string) string {
	requestUrl := "https://api.notion.com/v1/blocks/$BLOCK_ID/children?page_size=100"
	url := strings.Replace(requestUrl, "$BLOCK_ID", notionThanksfulnessBlockId, 1)
	return url
}

func makeRequestHeaders(notionApiKey string) map[string][]string {
	return map[string][]string{
		"Authorization":  {strings.Replace("Bearer $KEY", "$KEY", notionApiKey, 1)},
		"Notion-Version": {"2022-06-28"},
	}
}

func GetBlockChildren() ([]byte, error) {
	url := makeGetUrl(os.Getenv("NOTION_THANKSFULNESS_BLOCK_ID"))
	headers := makeRequestHeaders(os.Getenv("NOTION_API_KEY"))

	resp, err := lib.Get(url, headers)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Failed to get data from notion api. %v", err)
		return nil, err
	}

	// TODO status が200以外の時のハンドリング処理
	// resp.StatusCode

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Response body of notion api (Retrieve block children) is wrong. %v", resp.Status)
		return nil, err
	}

	return data, nil
}
