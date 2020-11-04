package protocol

type Protocol interface {
	Build() []byte
	Resolve([]byte)
}
