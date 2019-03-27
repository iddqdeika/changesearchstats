package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var(
	why = []string{"почему","из-за чего","пачему","Почему","зачем","Зачем","Пачему","с какого"}
	goods = []string{"goods","Goods","GOODS","goods.ru","гудс.ру","гудс","гудс ру"}
	so = []string{"так","настолько","настолька","очень","ужасно","ужасна","слишком","все ещё","все-ещё","всетаки"}
	bad = []string{"медленно","медленна","долго","долга","плохо","плоха"}
	loads = []string{"грузится","загружается","открывает страницу","аткрывается","открывается"}
)
func main() {
	template := NewTemplate([][]string{why[:3],goods[:3],so[:3],bad[:3],loads[:3]})
	generator := PhraseGenerator{}
	generator.addTemplate("loading",template)

	phraseStat := make(map[string]int)
	wg := sync.WaitGroup{}
	for i:=0; i<20; i++{
		wg.Add(1)
		go func(){
			for{
				time.Sleep(time.Millisecond*200)
				phrase, err := generator.generate("loading")
				if err != nil{
					panic(err)
				}
				phraseStat[*phrase]++
				fmt.Printf("sending phrase (%v try): \"%v\"\r\n", phraseStat[*phrase], *phrase)

				_, err = http.Get("https://goods.ru/catalog/?q=" + *phrase)
				if err != nil{
					panic(err)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

