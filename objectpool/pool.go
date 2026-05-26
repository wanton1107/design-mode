package objectpool

import "sync"

type Pool struct {
	sync.Mutex
	Inuse []interface{}
	Idle []interface{}
	new func() interface{}
}

func NewPool(new func() interface{}) *Pool {
	return &Pool{
		new: new,
	}
}

func (p *Pool) Get() interface{} {
	p.Lock()
	defer p.Unlock()
	var obj interface{}
	if len(p.Idle) != 0 {
		obj = p.Idle[0]
		p.Idle=p.Idle[1:]
		p.Inuse=append(p.Inuse, obj)
	}else{
		obj = p.new()
		p.Inuse=append(p.Inuse, obj)
	}
	return obj
}

func (p *Pool) Release(obj interface{}) {
	p.Lock()
	defer p.Unlock()
	p.Idle = append(p.Idle, obj)
	for i, o := range p.Inuse {
		if o == obj {
			p.Inuse = append(p.Inuse[:i], p.Inuse[i+1:]...)
			break
		}
	}
}
