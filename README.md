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
	params := map[string]string{"page": "1", "mode": "dev&ok=true"}
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
<li><a href="https://github.com/?mode=dev%26ok%3Dtrue&page=1">首页</a></li> 
<li><a href="https://github.com/?mode=dev%26ok%3Dtrue&page=16">上一页</a></li> 
<li><a href="https://github.com/?mode=dev%26ok%3Dtrue&page=16">16</a></li>
<li><span class="current">17</span></li>
<li><a href="https://github.com/?mode=dev%26ok%3Dtrue&page=18">18</a></li>
<li><a href="https://github.com/?mode=dev%26ok%3Dtrue&page=19">19</a></li>
<li><a href="https://github.com/?mode=dev%26ok%3Dtrue&page=20">20</a></li> 
<li><a href="https://github.com/?mode=dev%26ok%3Dtrue&page=18">下一页</a></li> 
<li><a href="https://github.com/?mode=dev%26ok%3Dtrue&page=39">尾页</a></li>
</ul>
```
