package embedding

type Client interface {
	Embed(message string) ([]float64, error)
}
