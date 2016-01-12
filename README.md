# paginator
a paginator for golang. surport bootstrap.

# useage
```
package main

import (
	"fmt"
	"github.com/orca-clock/paginator"
)

func main() {
	params := map[string]string{"page": "1", "mode": "dev"}
	p := paginator.NewPaginator("https://github.com/", params, 191, 5, 17)
	//p.SetGroupPages(10)
	//configs := map[string]string{"prev": "&lt;&lt;", "next": "&gt;&gt;", "up": "&lt;", "down": "&gt;", "first": "first", "last": "last"}
	//p.MergeConfigs(configs)
	//p.SetTheme(paginator.FULL_THEME)
	//theme: %first% %prev% %up% %link% %down% %next% %last%
	//p.SetTheme("%up% %down%")

	fmt.Println(p.Build())
}
```

# output
```
<ul class='pagination'>
<li><a href="https://github.com/?page=1&mode=dev">首页</a></li> 
<li><a href="https://github.com/?page=16&mode=dev">上一页</a></li> 
<li><a href="https://github.com/?page=16&mode=dev">16</a></li>
<li><span class="current">17</span></li>
<li><a href="https://github.com/?page=18&mode=dev">18</a></li>
<li><a href="https://github.com/?page=19&mode=dev">19</a></li>
<li><a href="https://github.com/?page=20&mode=dev">20</a></li> 
<li><a href="https://github.com/?page=18&mode=dev">下一页</a></li> 
<li><a href="https://github.com/?page=39&mode=dev">尾页</a></li>
</ul>
```
