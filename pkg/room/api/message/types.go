package message

const (
	TypePermissionRequest = iota + 1
	TypeUpstreamRelayRequest
)

type Response interface {
	Raw() []byte
}
