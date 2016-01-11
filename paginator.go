package paginator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Paginator struct {
	url         string
	params      map[string]string
	totalRows   int //总记录数
	listRows    int //每页记录数
	totalPages  int //总页数
	currentPage int //当前页数
	groupPages  int //当前分页组
	totalGroups int //总分页组数
	pageVar     string
	configs     map[string]string
	theme       string
}

const (
	FULL_THEME   = "<ul class='pagination'>\n%first% \n%prev% \n%up% \n%link% \n%down% \n%next% \n%end%\n</ul>"
	SIMPLE_THEME = "<ul class='pagination'>\n%first% \n%up% \n%link% \n%down% \n%end%\n</ul>"
)

const (
	DEFAULT_ROLL_PAGE = 5
	DEFAULT_LIST_ROWS = 20
	DEFAULT_PAGE_VAR  = "page"
	DEFAULT_THEME     = SIMPLE_THEME
)

var DEFAULT_CONFIGS = map[string]string{"up": "上一页",
	"down":  "下一页",
	"prev":  "前%d页",
	"next":  "后%d页",
	"first": "首页",
	"last":  "尾页"}

func NewPaginator(url string, params map[string]string, totalRows, listRows, currentPage int) *Paginator {
	p := &Paginator{groupPages: DEFAULT_ROLL_PAGE,
		totalRows: totalRows,
		url:       url,
		params:    params,
		listRows:  DEFAULT_LIST_ROWS,
		pageVar:   DEFAULT_PAGE_VAR,
		theme:     SIMPLE_THEME,
		configs:   DEFAULT_CONFIGS}

	p.totalRows = totalRows
	if listRows > 0 {
		p.listRows = listRows
	} else {
		p.listRows = DEFAULT_LIST_ROWS
	}

	p.totalPages = int(math.Ceil(float64(p.totalRows) / float64(p.listRows)))

	if currentPage < 1 {
		p.currentPage = 1
	} else if currentPage > p.totalPages {
		p.currentPage = p.totalPages
	} else {
		p.currentPage = currentPage
	}

	return p
}

func (p *Paginator) MergeConfigs(configs map[string]string) {
	for key, value := range configs {
		p.configs[key] = value
	}
}

func (p *Paginator) SetGroupPages(groupPages int) {
	if groupPages > 0 {
		p.groupPages = groupPages
	}
}

func (p *Paginator) Build() string {
	if p.totalRows == 0 {
		return ""
	}

	p.params[p.pageVar] = "__PAGE__"
	querys := make([]string, 0)
	for key, value := range p.params {
		querys = append(querys, key+"="+value)
	}

	url := p.url + "?" + strings.Join(querys, "&")

	replacePairs := make([]string, 0)

	upPage := "" //上一页
	upRow := p.currentPage - 1
	if upRow > 0 {
		upPage = fmt.Sprintf("<li><a href=\"%s\">%s</a></li>",
			strings.Replace(url, "__PAGE__", strconv.Itoa(upRow), 1),
			p.configs["up"])
	}
	replacePairs = append(replacePairs, "%up%", upPage)

	downPage := "" //下一页
	downRow := p.currentPage + 1
	if downRow <= p.totalPages {
		downPage = fmt.Sprintf("<li><a href=\"%s\">%s</a></li>",
			strings.Replace(url, "__PAGE__", strconv.Itoa(downRow), 1),
			p.configs["down"])
	}
	replacePairs = append(replacePairs, "%down%", downPage)

	p.totalGroups = int(math.Ceil(float64(p.totalPages) / float64(p.groupPages)))

	firstPage := "" //首页
	endPage := ""   //尾页
	prePage := ""   // 前n页
	nowPageGroup := int(math.Ceil(float64(p.currentPage) / float64(p.groupPages)))
	if nowPageGroup != 1 {
		preRow := p.currentPage - p.groupPages
		prev := p.configs["prev"]
		if strings.Index(prev, "%d") != -1 {
			prev = fmt.Sprintf(p.configs["prev"], p.groupPages)
		}
		prePage = fmt.Sprintf("<li><a href=\"%s\">%s</a></li>",
			strings.Replace(url, "__PAGE__", strconv.Itoa(preRow), 1),
			prev)
		//首页
		firstPage = fmt.Sprintf("<li><a href=\"%s\">%s</a></li>",
			strings.Replace(url, "__PAGE__", "1", 1),
			p.configs["first"])
	}
	replacePairs = append(replacePairs, "%prev%", prePage, "%first%", firstPage)

	nextPage := "" //后n页
	if nowPageGroup != p.totalGroups {
		nextRow := p.currentPage + p.groupPages
		next := p.configs["next"]
		if strings.Index(next, "%d") != -1 {
			next = fmt.Sprintf(p.configs["next"], p.groupPages)
		}
		nextPage = fmt.Sprintf("<li><a href=\"%s\">%s</a></li>",
			strings.Replace(url, "__PAGE__", strconv.Itoa(nextRow), 1),
			next)

		endPage = fmt.Sprintf("<li><a href=\"%s\">%s</a></li>",
			strings.Replace(url, "__PAGE__", strconv.Itoa(p.totalPages), 1),
			p.configs["last"])
	}
	replacePairs = append(replacePairs, "%next%", nextPage, "%end%", endPage)

	linkPages := make([]string, 0)
	for i := 1; i <= p.groupPages; i++ {
		page := (nowPageGroup-1)*p.groupPages + i
		if page != p.currentPage {
			if page < p.totalPages {
				linkPage := fmt.Sprintf("<li><a href=\"%s\">%s</a></li>",
					strings.Replace(url, "__PAGE__", strconv.Itoa(page), 1),
					strconv.Itoa(page))
				linkPages = append(linkPages, linkPage)
			} else {
				break
			}
		} else {
			if p.totalPages != 1 {
				linkPage := fmt.Sprintf("<li><span class=\"current\">%s</span></li>",
					strconv.Itoa(page))
				linkPages = append(linkPages, linkPage)
			}
		}
	}
	linkPage := strings.Join(linkPages, "\n")
	replacePairs = append(replacePairs, "%link%", linkPage)

	replacer := strings.NewReplacer(replacePairs...)
	return replacer.Replace(p.theme)
}
