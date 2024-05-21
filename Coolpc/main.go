package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Category struct {
	Name          string
	SubCategories []SubCategory
}

type SubCategory struct {
	Name  string
	Items []Item
}

type Item struct {
	Name  string
	Price int
}

var (
	lists []Category
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("#tbdy tr", func(e *colly.HTMLElement) {
		title := e.ChildText("td.t")
		if len(title) == 0 {
			return
		}
		fmt.Println("主類別", title)
		titleIndex := -1
		for k, v := range lists {
			if v.Name == title {
				titleIndex = k
				break
			}
		}
		if titleIndex == -1 {
			lists = append(lists, Category{
				Name: title,
			})
			titleIndex = len(lists) - 1
		}

		e.ForEach("td:nth-child(3) > select > optgroup", func(i int, e *colly.HTMLElement) {
			subTitle := e.Attr("label")
			fmt.Println("子類別", e.Attr("label"))
			//fmt.Println(e.Text)
			subCategoryIndex := -1
			for k, v := range lists[titleIndex].SubCategories {
				if v.Name == subTitle {
					subCategoryIndex = k
					break
				}
			}
			if subCategoryIndex == -1 {
				lists[titleIndex].SubCategories = append(lists[titleIndex].SubCategories, SubCategory{
					Name: subTitle,
				})
				subCategoryIndex = len(lists[titleIndex].SubCategories) - 1
			}
			e.ForEach("option", func(i int, e *colly.HTMLElement) {
				for _, a := range e.Attributes {
					if a.Key == "disabled" {
						return
					}
				}
				splits := strings.Split(e.Text, "$")
				if len(splits) < 2 {
					return
				}
				name := splits[0][:len(splits[0])-2]
				price, ok := strconv.Atoi(strings.Split(splits[1], " ")[0])
				if ok != nil {
					return
				}
				lists[titleIndex].SubCategories[subCategoryIndex].Items = append(lists[titleIndex].SubCategories[subCategoryIndex].Items, Item{
					Name:  name,
					Price: price,
				})
				fmt.Printf("商品名稱: [%s], 價格: [%d]\n", name, price)
			})
		})
		//fmt.Println("Items", len(items))
	})

	c.SetRequestTimeout(20 * time.Second)
	c.Visit("https://coolpc.com.tw/evaluate.php")
	//for k, v := range category {
	//	fmt.Println(k)
	//	for _, i := range v {
	//		fmt.Println(i.Name)
	//	}
	//}

	http.HandleFunc("/", Serve)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func Serve(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, lists)
}
