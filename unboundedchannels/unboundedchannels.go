package unboundedchannels

import (
//	"fmt"
)

func makechannel() (chan<- interface{}, <-chan interface{}) {
	in := make(chan interface{})
	out := make(chan interface{})
	go func() {
		var  q[] interface{}
		b := func() interface{} {
		if len(q) < 0 {
			return nil
		}
		return q[0]
		}
		outC := func() chan interface{} {
			if len(q) < 0 { return nil}
			return out
		}

		for len(q)>0 || in != nil {
			select {
			case i, ok := <-in:
				if !ok {
					in = nil
				} else {
					q = append(q, i)
				}
			case outC() <- b():
				q = q[1:]

	    		 }
		}
	close(out)
	}()

	return in, out
}
/*
func main() {
	in, out := makechannel()
	insert(in)
}
*/
