package notion

func Convert(input *APIResponse) (*ThanksList, error) {
	var editedAt string
	var items []string
	for _, result := range input.Results {
		if result.Type == "paragraph" {
			editedAt = result.Paragraph.RichText[0].PlainText
		} else if result.Type == "bulleted_list_item" {
			for _, richText := range result.BulletedListItem.RichText {
				items = append(items, richText.PlainText)
			}
		}
	}

	return &ThanksList{
		EditedAt: editedAt,
		Items:    items,
	}, nil
}
