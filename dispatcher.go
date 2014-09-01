package events

type IEventDispatcher interface {
	DispatchEvent(event IEvent)
	RemoveEventListener(string, IEventHandler)
	AddEventListener(string, IEventHandler)
	HasEventListener(string) bool
	RemoveEventListeners(name string)
	Dispose()
}
func NewEventDispatcher() IEventDispatcher {
	ed := &eventDispatcher{}
	ed.handlers = make(map[string][]IEventHandler)
	return ed
}

type eventDispatcher struct {
	handlers map[string][]IEventHandler
}

type IEventHandler interface {
	Handle(e IEvent)
}
type HandlerFunc func(IEvent)

func (this HandlerFunc) Handle(e IEvent) {
	this(e)
}

func (this *eventDispatcher) DispatchEvent(event IEvent) {
	event.setTarget(this)
	if handlers, ok := this.handlers[event.Name()]; ok {
		l := len(handlers)
		for i := 0; i < l; i++ {
			handlers[i].Handle(event)
		}
	}
}

func (this *eventDispatcher) RemoveEventListener(name string, listener IEventHandler) {
	if handlers, ok := this.handlers[name]; ok {
		for i := 0; i < len(handlers); i++ {
			if listener == handlers[i] {
				this.handlers[name] = append(handlers[:i], handlers[i+1:]...)
				break
			}
		}
	}
}

func (this *eventDispatcher) AddEventListener(name string, listener IEventHandler) {
	if handlers, ok := this.handlers[name]; ok {
		//禁止多次添加同一个handler
		for i := 0; i < len(handlers); i++ {
			if listener == handlers[i] {
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
		_, ok := this.handlers[name]
		return ok
	}
	return false
}

func (this *eventDispatcher) Dispose() {
	for k := range this.handlers {
		delete(this.handlers, k)
	}
}
