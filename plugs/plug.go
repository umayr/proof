package plugs

type Plug interface {
	Fetch()
}

var Handlers = map[string]interface{}{
	"Independent": new(Independent).Fetch,
}

type Payload struct {
	Name  string
	Value string
}