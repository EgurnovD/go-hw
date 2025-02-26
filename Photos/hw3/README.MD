Привет. Я в лекции говорил, что надо два примера объяснить, разобрать
### Первый: С разными версиями GO  
```
package main

import (
	"fmt"
)

func main() {
	var numbers []*int
	for _, value := range []int{10, 20, 30, 40} {
		fmt.Printf("%v %d\n", &value, value)
		numbers = append(numbers, &value)
	}
	fmt.Println("----------")
	for _, number := range numbers {
		fmt.Printf("%v %d\n", number, *number)
	}
}
```

### Второй
```
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(7)
	for i := 0; i < 9; i++ {
		i := i
		go func(i int) {
			defer wg.Done()
			fmt.Println("    Почему, КОЛЯ?", i)
		}(i)
	}
	wg.Wait()

	// time.Sleep(1 * time.Second)
	fmt.Println("Паника")
}
```

Почему, когда я выставляю `runtime.GOMAXPROCS(4)`, то порядок выполнения может быть любым:
```
                           Почему, КОЛЯ? 3
                           Почему, КОЛЯ? 0
                           Почему, КОЛЯ? 8
                           Почему, КОЛЯ? 4
                           Почему, КОЛЯ? 2
                           Почему, КОЛЯ? 5
                           Почему, КОЛЯ? 6
Паника
```
А если выставляю `runtime.GOMAXPROCS(1)` То всегда будет начинаться:
```
Почему, КОЛЯ? 8
```
