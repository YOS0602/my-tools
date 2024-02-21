package sample

// See https://qiita.com/tonnsama/items/af4ec31e35471220c8a2

type Cat struct {
	sound string
}

func NewCat(sound string) *Cat {
	return &Cat{sound: sound}
}

func (c *Cat) Cry() string {
	return c.sound
}
