package server

import "testing"

func Test_encodeWebsocketResponseKey_returnsHash(t *testing.T) {
	// values taken directly from RFC
	// https://datatracker.ietf.org/doc/html/rfc6455#section-1.2
	websocketKey := "dGhlIHNhbXBsZSBub25jZQ=="
	guid := "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	expected := "s3pPLMBiTxaQ9kYGzzhZRbK+xOo="

	serverWsKey := encodeWebsocketResponseKey(guid, websocketKey)

	if serverWsKey != expected {
		t.Errorf("serverWsKey=%q, want %q", serverWsKey, expected)
	}

}
