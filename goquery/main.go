package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"sort"
	"strings"
	"time"
)

func GetHttpHTML(url string, selector string) string {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false), //设置成无浏览器弹出
		chromedp.Flag("blink-settings", "imageEnable=false"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"),
	}
	c, _ := chromedp.NewExecAllocator(context.Background(), options...)
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	_ = chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)
	
	timeOutCtx, cancel := context.WithTimeout(chromeCtx, 90*time.Second)
	defer cancel()
	
	var htmlContent string
	err := chromedp.Run(timeOutCtx,
		chromedp.Navigate(url),
		//需要爬取的网页的url
		//chromedp.WaitVisible(`div[class="fp-tournament-award-badge_awardContent__dUtoO"]`),
		chromedp.WaitVisible(selector),
		//等待某个特定的元素出现
		chromedp.OuterHTML(`document.querySelector("body")`, &htmlContent, chromedp.ByJSPath),
		//生成最终的html文件并保存在htmlContent文件中
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(htmlContent)
	//doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	
	//awardTmp := [][]string{}
	//doc.Find(`div[class="fp-tournament-award-badge_awardContent__dUtoO"]`).Each(func(i int, selection *goquery.Selection) {
	//	award := selection.Find(`p[class="fp-tournament-award-badge_awardName__JpsZZ"]`).Text()
	//	name := selection.Find(`h4[class=" fp-tournament-award-badge_awardWinner__P_z2d"]`).Text()
	//	country := selection.Find(`p[class="fp-tournament-award-badge_awardWinnerCountry__EmjVU"]`).Text()
	//
	//	awardTmp = append(awardTmp, []string{award, name, country})
	//})
	return htmlContent
}

//
//func main() {
//	//content := GetHttpHTML("https://www.fifa.com/tournaments/mens/worldcup/2018russia",
//	//	`#content > div > section.fp-tournament-award-badge-carousel_awardBadgeCarouselSection__w_Ys5 > div > div > div.col-12.fp-tournament-award-badge-carousel_awardCarouselColumn__fQJLf.g-0 > div > div > div > div > div > div`)
//	//fmt.Println(content)
//	//
//	//player, _ := GetMatch(content)
//	//fmt.Println(player)
//
//	content := GetHttpHTML("https://www.fifa.com/tournaments/mens/worldcup/2018russia/teams/43922",
//		`#content > div > div.fp-squad_squadContainer__PpIp8 > div > div.fp-squad_squadCardsContainer__0PeCv`)
//	//fmt.Println(content)
//
//	player, _ := GetPlayer(content)
//	fmt.Println("--------------------")
//	fmt.Println(player)
//}

func GetHttpHtmlContent(url string, selector string) (html string, err error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false), //不会弹出浏览器窗口
		chromedp.Flag("blink-settings", "imageEnable=false"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"),
	}
	c, _ := chromedp.NewExecAllocator(context.Background(), options...)
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	_ = chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)
	
	timeOutCtx, cancel := context.WithTimeout(chromeCtx, 90*time.Second)
	defer cancel()
	
	var htmlContent string
	err = chromedp.Run(timeOutCtx,
		chromedp.Navigate(url),
		//需要爬取的网页的url
		//chromedp.WaitVisible(`div[class="fp-tournament-award-badge_awardContent__dUtoO"]`),
		chromedp.WaitVisible(selector),
		//等待某个特定的元素出现
		chromedp.OuterHTML(`document.querySelector("body")`, &htmlContent, chromedp.ByJSPath),
		//生成最终的html文件并保存在htmlContent文件中
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(htmlContent)
	return htmlContent, err
}

func GetPlayer(html string) ([][]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	tmp := [][]string{}
	var xing, ming, seat, paiming, country string
	doc.Find(`div[class="fp-squad_squadCardsContainer__0PeCv"]`).Each(func(i int, selection *goquery.Selection) {
		
		selection.Find(`div[class="fp-squad-player-card_playerDetails__dMUgv"]`).Each(func(i int, player *goquery.Selection) {
			player.Find(`div[class="fp-squad-player-card_firstRow__4l1On"]`).Each(func(i int, name *goquery.Selection) {
				xing = name.Find(`div[class="fp-squad-player-card_firstName__o0cWG"]`).Text()
				ming = name.Find(`div[class="fp-squad-player-card_lastName__TNGsc"]`).Text()
				paiming = name.Find(`div[class="fp-squad-player-card_jerseyNumber__wfEsB"]`).Text()
				
			})
			
			player.Find(`div[class="fp-squad-player-card_secondRow__tWyB6"]`).Each(func(i int, s *goquery.Selection) {
				seat = s.Find(`div[class="fp-squad-player-card_position__UHe_f"]`).Text()
				s.Find(`div[class="fp-squad-player-card_flag__jMbx6"]`).Each(func(i int, img *goquery.Selection) {
					country, _ = img.Find(`img[class="image_img__jrck5"]`).Attr("alt")
				})
			})
			
			tmp = append(tmp, []string{xing + ming, seat, paiming, country})
		})
		
	})
	
	return tmp, nil
}

//在拿到的HTML中找到获奖球员的信息
func GetMatch(html string) ([][]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	//var award, name, country string
	Tmp := [][]string{}
	var imgleft, countryleft, score, match, imgright, countryright string
	doc.Find(`div[class="fp-match-card_matchInfo__Kjs5W"]`).Each(func(i int, selection *goquery.Selection) {
		//award := selection.Find(`p[class="fp-tournament-award-badge_awardName__JpsZZ"]`).Text()
		//name := selection.Find(`h4[class=" fp-tournament-award-badge_awardWinner__P_z2d"]`).Text()
		//country := selection.Find(`p[class="fp-tournament-award-badge_awardWinnerCountry__EmjVU"]`).Text()
		
		//左边国家
		selection.Find(`div[class="fp-match-card_team__sv7NJ align-items-center justify-items-center"]`).Each(func(i int, left *goquery.Selection) {
			imgleft, _ = left.Find(`img[class="image_img__jrck5"]`).Attr("src")
			countryleft = left.Find(`span[class="fp-match-card_homeTeam__qhywg card-heading-tiny"]`).Text()
		})
		
		//中间比分
		selection.Find(`div[class="d-flex align-items-center fp-match-card_centerItems__Y5fV1 fp-match-card_matchResults__pGrAc"]`).Each(func(i int, center *goquery.Selection) {
			score = center.Find(`h4[class=" fp-match-card_matchScore__5VcOd ff-pt-8"]`).Text()
			match = center.Find(`p[class="fp-match-card_matchNumber__5npsh ff-text-blue-cinema"]`).Text()
		})
		
		//右边国家
		selection.Find(`div[class="fp-match-card_team__sv7NJ align-items-center justify-items-center"]`).Each(func(i int, right *goquery.Selection) {
			imgright, _ = right.Find(`img[class="image_img__jrck5"]`).Attr("src")
			countryright = right.Find(`span[class="fp-match-card_awayTeam__i_YqB card-heading-tiny"]`).Text()
		})
		Tmp = append(Tmp, []string{imgleft, countryleft, score, match, imgright, countryright})
	})
	//fmt.Println(awardTmp)
	return Tmp, err
}

func accountsMerge(accounts [][]string) (ans [][]string) {
	emailIndex := map[string]int{}
	emailName := map[string]string{}
	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			if _, has := emailIndex[email]; !has {
				emailIndex[email] = len(emailIndex)
				emailName[email] = name
			}
		}
	}
	//构建并查集
	parent := make([]int, len(emailIndex))
	for i := range parent {
		parent[i] = i
	}
	
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) //路径压缩,直接指向根节点
		}
		return parent[x]
	}
	
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}
	
	for _, account := range accounts {
		firstIndex := emailIndex[account[1]]
		for _, email := range account[2:] {
			union(emailIndex[email], firstIndex)
		}
	}
	
	indexToEmails := map[int][]string{}
	for email, index := range emailIndex {
		index = find(index)
		indexToEmails[index] = append(indexToEmails[index], email)
	}
	
	for _, emails := range indexToEmails {
		sort.Strings(emails)
		account := append([]string{emailName[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return
}

//func Find(i int) int {
//	if father[i] == i { //递归出口，当找到了祖先节点时，就返回
//		return i
//	} else {
//		return Find(father[i]) //不断往上查找祖先节点
//	}
//}

type UnionSet struct {
	Father []int
}

func (u UnionSet) Init(n int) []int {
	//var father [math.MaxInt]int
	for i := 1; i <= n; i++ {
		u.Father[i] = i
	}
	return u.Father
}

func (u UnionSet) Union(i int, j int) {
	iFather := u.Father[i]      //找到i的祖先节点
	jFather := u.Father[j]      //找到j的祖先节点
	u.Father[iFather] = jFather //i的祖先指向j的祖先
}

func (u UnionSet) Find(i int) int {
	if i == u.Father[i] {
		return i
	} else {
		u.Father[i] = u.Find(u.Father[i]) //进行路径压缩
		return u.Father[i]                //返回父节点
	}
}

func main() {
	tmp := [][]int{{10, 7}, {2, 4}, {5, 7}, {1, 3}, {8, 9}, {1, 2}, {5, 6}, {2, 3}, {3}, {3, 4}, {7, 10}, {8, 9}}
	unionset := UnionSet{
		Father: make([]int, tmp[0][0]+1),
	}
	unionset.Init(tmp[0][0])
	for i := 1; i <= tmp[0][1]; i++ {
		unionset.Union(tmp[i][0], tmp[i][1]) //将关系压入并查集
	}
	if len(tmp[tmp[0][1]+1]) != 1 {
		return
	} //{3}问询的数量
	for i := tmp[0][1] + 2; i < len(tmp); i++ {
		if unionset.Find(tmp[i][0]) == unionset.Find(tmp[i][1]) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
