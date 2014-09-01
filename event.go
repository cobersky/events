package events


type IEvent interface {
	Name()string
	Data()interface {}
	Target()IEventDispatcher
	setTarget(IEventDispatcher)
}
type event struct {
	name string
	data interface {}
	target IEventDispatcher
}
func (this *event)Name()string{
	return this.name
}
func(this *event)Data()interface {}{
	return this.data
}
func(this *event)Target()IEventDispatcher{
	return this.target
}
func(this *event)setTarget(t IEventDispatcher){
	this.target=t
}
func NewEvent(name string,data interface {})*event{
	e:=event{}
	e.name=name
	e.data=data
	return &e
}
