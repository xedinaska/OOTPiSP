package encoding

type Default struct{}

func (n *Default) Encode(v interface{}) string {
	return v.(string)
}

func (n *Default) Decode(v string) interface{} {
	return v
}

func (n *Default) Name() string {
	return "Default"
}
