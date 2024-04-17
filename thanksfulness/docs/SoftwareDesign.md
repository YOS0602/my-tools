# Software Design

## Slackへのリクエストデータ生成

### input

NotionAPIのレスポンスデータを加工したもの

```json
{}
```

<details><summary>NotionAPIレスポンスを TypeScript で型定義するとこう ※一部プロパティ省略</summary>

```ts
type NotionAPIResponse = {
  object: 'list';
  results: {
    object: string;
    id: string;
    parent: {
      type: string;
      block_id: string;
    };
    created_time: Date;
    last_edited_time: Date;
    created_by: {
      object: string;
      id: string;
    };
    last_edited_by: {
      object: string;
      id: string;
    };
    has_children: boolean;
    archived: boolean;
    in_trash: boolean;
    type: string | 'paragraph' | 'bulleted_list_item';
    paragraph?: {
      rich_text: {
        type: string;
        plain_text: string;
      }[];
    };
    bulleted_list_item?: {
      rich_text: {
        type: string;
        plain_text: string;
      }[];
    };
  }[];
  next_cursor: boolean;
  has_more: boolean;
  type: string;
  block: Record<string, any>;
  request_id: string;
};
```

</details>

### output

SlackAPIへのリクエストデータ

```json
{}
```
