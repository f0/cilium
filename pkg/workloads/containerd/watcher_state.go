package containerd

import (
	"sync"

	dTypesEvents "github.com/docker/engine-api/types/events"
)

// watcherState holds global close flag, per-container queues for events and ignore toggles
type watcherState struct {
	sync.Mutex
	closed bool
	events map[string]chan dTypesEvents.Message
}

func newWatcherState() *watcherState {
	return &watcherState{events: make(map[string]chan dTypesEvents.Message)}
}

func (ws *watcherState) startEventHandler(containerID string) (created bool) {
	ws.Lock()
	defer ws.Unlock()

	if _, found := ws.events[containerID]; !found {
		q := make(chan dTypesEvents.Message, eventQueueBufferSize)
		ws.events[containerID] = q
		go processContainerEvents(containerID, q)
		created = true
	}
	return
}

func (ws *watcherState) enqueueByContainerID(e dTypesEvents.Message) {
	ws.startEventHandler(e.Actor.ID)
	ws.events[e.Actor.ID] <- e
}

func (ws *watcherState) reapEmpty() {
	ws.Lock()
	defer ws.Unlock()

	for id, q := range ws.events {
		if len(q) == 0 {
			close(q)
			delete(ws.events, id)
		}
	}
}

func (ws *watcherState) close() {
	ws.Lock()
	defer ws.Unlock()

	ws.closed = true
	for id, q := range ws.events {
		close(q)
		delete(ws.events, id)
	}
}

func (ws *watcherState) isClosed() bool {
	ws.Lock()
	defer ws.Unlock()

	return ws.closed
}
