package stickyticker

import(
	"time"
	"context"
	"sync"
)

// StickyTicker struct
type StickyTicker struct {
	*time.Timer
	interval time.Duration
	delay time.Duration
	c chan []uint
	mtx sync.Mutex
}

/*
 function name: NewStickyTicker
 NewStickyTicker: This function create new StickyTicker object.
 Requires:
 	interval: set interval time to execute a callback.  
 	delay: set delay to execute a callback if need. 
 	ctx: it needs context object to create StickyTicker.
 	callback: your customize function. 
*/

func NewStickyTicker( interval uint, delay uint, ctx context.Context, callback func(time.Time) ) *StickyTicker {

	di := time.Duration(interval) * time.Second
	do := time.Duration(delay) * time.Second
	adjt, _ := GetNextTrigger(di, do)
	s := &StickyTicker{
		Timer: time.NewTimer(adjt),
		interval: di,
		delay: do,
		c: make(chan []uint),
		mtx: sync.Mutex{},
	}
	
	stop := func() {
		if !s.Timer.Stop() {
			select {
	        case <-s.Timer.C:
    	    default:
        	}
		}
	}

	go func() {
		for {
			select {
			case t := <-s.Timer.C:
				stop()
				if callback != nil {
					callback(t)
				}
				adjt,  _ := GetNextTrigger(s.interval, s.delay)
				s.Timer.Reset(adjt)
			case <-ctx.Done():
				s.mtx.Lock()
				close(s.c)
				s.mtx.Unlock()
				stop()
				return
			case args := <-s.c:
				s.interval = time.Duration(args[0]) * time.Second
				s.delay = time.Duration(args[1]) * time.Second
				stop()
				adjt,  _ := GetNextTrigger(s.interval, s.delay)
				s.Timer.Reset(adjt)
			}
		}
	}()

	return s
}

/*
 function name: Stop
 Stop: This method be able to stop StickyTicker.
 info) StickyTicker has not re-start after stoped.
*/
func(s *StickyTicker) Stop( cancel func() ){
	cancel()
}

/*
 function name: Reset
 Reset: This method be able to set new interval and offset value.
*/
func(s *StickyTicker) Reset( interval uint, delay uint ){
	s.mtx.Lock()
	if s.c != nil {
		s.c <- []uint{interval, delay}
	}
	s.mtx.Unlock()
}

/*
 function name: GetNextTrigger
 GetNextTrigger: This function provides next interval from now and execution datetime.
*/
func GetNextTrigger( interval time.Duration, delay time.Duration) (time.Duration, time.Time){
	n := time.Now().Round(time.Millisecond) // round by milliseconds.
	next := n.Add(interval).Add(delay).Truncate(interval)
	adjt := next.Sub(n) + delay
	return adjt, next
}
