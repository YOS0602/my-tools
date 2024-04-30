package notion

import "log"

// usecase ロジック
// Notion から感謝リストを取得する
func GetThanksList() (*ThanksList, error) {
	// NotionAPI から Block 配下のデータを取得
	data, err := GetBlockChildren()
	if err != nil {
		return nil, err
	}

	// レスポンスをunmarshalする
	decoded, err := UnmarshalResponse(data)
	if err != nil {
		log.Fatalf("Failed to unmarshal notion api response. %v", data)
		return nil, err
	}

	// 欲しいフィールドだけ取り出す
	thanksList, nil := Convert(decoded)
	if err != nil {
		return thanksList, err
	}

	return thanksList, nil
}
