package events

import (
	"log"
)

type IEventDispatcher interface {
	DispatchEvent(event IEvent)
	RemoveEventListener(string, IEventHandler)
	AddEventListener(string, IEventHandler)
	HasEventListener(string) bool
	RemoveEventListeners(name string)
	Dispose()
}

type IEventHandler interface {
	Handle(e IEvent)
}
type eventDispatcher struct {
	handlers map[string][]IEventHandler
}
type HandlerFunc func(IEvent)

func (this HandlerFunc) Handle(e IEvent) {
	this(e)
}

func NewEventDispatcher() IEventDispatcher {
	ed := &eventDispatcher{}
	ed.handlers = make(map[string][]IEventHandler)
	return ed
}

func (this *eventDispatcher) DispatchEvent(event IEvent) {
	event.setTarget(this)
	if handlers, has := this.handlers[event.Name()]; has {
		l := len(handlers)
		for i := 0; i < l; i++ {
			handlers[i].Handle(event)
		}
	}
}

func (this *eventDispatcher) RemoveEventListener(name string, listener IEventHandler) {
	if handlers, has := this.handlers[name]; has {
		for i := 0; i < len(handlers); i++ {
			if listener == handlers[i] {
				log.Println("find listener and reomve it!")
				this.handlers[name] = append(handlers[:i], handlers[i+1:]...)
				break
			}
		}
	}
}

func (this *eventDispatcher) AddEventListener(name string, listener IEventHandler) {
	if handlers, has := this.handlers[name]; has {
		for i := 0; i < len(handlers); i++ {
			if listener == handlers[i] {
				log.Println("had same handler!")
				return
			}
		}
		handlers = append(handlers, listener)
	}else {
		this.handlers[name] = []IEventHandler{listener}
	}
}

func (this *eventDispatcher) RemoveEventListeners(name string) {
	if this.HasEventListener(name) {
		delete(this.handlers, name)
	}
}

func (this *eventDispatcher) HasEventListener(name string) bool {
	if this.handlers != nil {
		if _, ok := this.handlers[name]; ok {
			return true
		}
	}
	return false
}

func (this *eventDispatcher) Dispose() {
	for k := range this.handlers {
		delete(this.handlers, k)
	}
}
