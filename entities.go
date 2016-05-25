package rosette

import "encoding/json"

// EntitiesInput request output.
type EntitiesInput struct {
	Content string `json:"content"`
}

type Entity struct {
	Count        int     `json:count`
	IndocChainId int     `json:"indocChainId"`
	Mention      string  `json:mention`
	Normalized   string  `json:normalized`
	Type         string  `json:"type"`
	EntityId     string  `json:"entityId"`
	Confidence   float32 `json:"confidence"`
}

// EntitiesOutput request output.
type EntitiesOutput struct {
	Entities []Entity `json:"entities"`
}

func (c *Client) NewEntitiesInput(content string) *EntitiesInput {
	return &EntitiesInput{
		Content: content,
	}
}

// Entities returns a EntitiesOutput.
func (c *Client) Entities(in *EntitiesInput) (out *EntitiesOutput, err error) {
	body, err := c.call("entities", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
