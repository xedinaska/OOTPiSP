package encoding

type Encoder interface {
	Name() string
	Encode(interface{}) string
	Decode(string) interface{}
}
