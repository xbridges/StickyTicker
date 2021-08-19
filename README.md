# StickyTicker
  StickyTicker to start on time.
  A StickyTicker that runs at "0:00,  0:30... minutes per hour" required for resident programs.

## usage
  ```
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("start", time.Now())
	
	s := NewStickyTicker(5, 0, func(t time.Time){
		fmt.Println(t)
	})
	defer s.Stop()

	time.Sleep(15 * time.Second)
	fmt.Println("change interval", time.Now())
	s.Reset(3,0)
	fmt.Println("changed interval. 5->3")
	time.Sleep(15 * time.Second)
	fmt.Println(s)

  ```  
