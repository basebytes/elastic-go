package constructor

type Builder interface {
	Build() map[string]interface{}
}