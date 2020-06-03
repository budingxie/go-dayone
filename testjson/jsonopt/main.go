package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	//many more fields...
}

//忽略struct里面的空字段
func testIgnoreField() {
	user := User{
		Email:    "qq@qaq",
		Password: "",
	}
	bytes, err := json.Marshal(struct {
		*User
		Password bool `json:"password,omitempty"`
	}{
		User: &user,
	})
	data, err := json.Marshal(&user)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", bytes)
	fmt.Printf("%s\n", data)
}

//临时添加额外的字段
func testAddField() {
	user := User{
		Email:    "qq@qaq",
		Password: "",
	}
	bytes, err := json.Marshal(struct {
		*User
		Token string `json:"token"`
		Age   int    `json:"age"`
	}{
		User:  &user,
		Token: "tokenVal",
		Age:   12,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", bytes)
}

type BlogPost struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}
type Analytics struct {
	Visitors  int `json:"visitors"`
	PageViews int `json:"page_views"`
}

//临时粘合两个struct
func testBonding() {
	post := BlogPost{
		URL:   "123",
		Title: "www",
	}

	bytes, err := json.Marshal(struct {
		*BlogPost
		*Analytics
	}{
		&post,
		&Analytics{
			Visitors:  123,
			PageViews: 23,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", bytes)
}

//一个json切分成两个struct
func segmentJson() {
	post := BlogPost{}
	analytics := Analytics{}
	json.Unmarshal([]byte(`{"url":"123","title":"www","visitors":123,"page_views":23}`), &struct {
		*BlogPost
		*Analytics
	}{&post, &analytics})
	fmt.Println(post, analytics)
}

//临时改名struct的字段
func updateStr() {

}

func main() {
	//testIgnoreField()
	//testAddField()
	//testBonding()
	segmentJson()
}
