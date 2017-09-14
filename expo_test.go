package expo

import "testing"

func TestChunkPushNotifications(t *testing.T) {
	messages := []*ExpoPushMessage{}

	for i:=0; i<100; i++ {
		messages = append(messages, NewExpoPushMessage())
	}

	chunked := ChunkPushNotifications(messages)

	if len(chunked) != 1 {
		t.Error("Expected 100 messages to be in 1 chunk")
	}

	messages = append(messages, NewExpoPushMessage())
	chunked = ChunkPushNotifications(messages)

	if len(chunked) != 2 {
		t.Errorf("Expected %d messages to be in 2 chunks. Got %d", len(messages), len(chunked))
	}
}