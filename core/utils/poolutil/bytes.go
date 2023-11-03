package poolutil

type Bytes struct {
	c chan []byte
	w int
}

func NewBytes(length int, width int) (bp *Bytes) {
	return &Bytes{
		c: make(chan []byte, length),
		w: width,
	}
}

var _ Pool[[]byte] = (*Bytes)(nil)

func (p *Bytes) Get() (b []byte) {
	select {
	case b = <-p.c:
	// reuse existing buffer
	default:
		// create new buffer
		b = make([]byte, p.w)
	}
	return
}

func (p *Bytes) Put(b []byte) {
	if cap(b) < p.w {
		// someone tried to put back a too small buffer, discard it
		return
	}

	select {
	case p.c <- b[:p.w]:
		// buffer went back into pool
	default:
		// buffer didn't go back into pool, just discard
	}
}
