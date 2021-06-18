package stickyticker

import(
	"testing"
	"fmt"
	"time"
)

func TestStickyTickTimes( t *testing.T ){
	defer func(){
		fmt.Println("TestHelloSchedule OUT")		
	}()

	fmt.Println("TestHelloSchedule IN")

	times := []time.Duration{3,10,30,60,120,300,1800,3600}

	for _, t := range times{
		i := time.Duration(t*time.Second)
		o := time.Duration(0)
		d, n := GetNextTrigger(i,o)
		fmt.Println( time.Now(), n, t*time.Second, d )
	}
}

func TestStickyticker( t *testing.T ){

	fmt.Println("start TestStickyticker", time.Now())
	s := NewStickyTicker(5, 0, func(t time.Time){
		fmt.Println(t)
	}) // tickerと同じく、作ったら即実行
	defer s.Stop()

	time.Sleep(15 * time.Second)
	s.Reset(3,0)
	time.Sleep(15 * time.Second)
	fmt.Println(s)
}
