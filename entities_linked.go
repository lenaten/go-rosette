package rosette

import "encoding/json"

// EntitiesInput request output.
type EntitiesLinkedInput struct {
	Content string `json:"content"`
}

// EntitiesOutput request output.
type EntitiesLinkedOutput struct {
	Entities []Entity `json:"entities"`
}

func (c *Client) NewEntitiesLinkedInput(content string) *EntitiesInput {
	return &EntitiesInput{
		Content: content,
	}
}

// Entities returns a EntitiesOutput.
func (c *Client) EntitiesLinked(in *EntitiesInput) (out *EntitiesOutput, err error) {
	body, err := c.call("entities/linked", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
