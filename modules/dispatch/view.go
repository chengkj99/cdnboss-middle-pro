package dispatch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

// View 获取所有view,   doType: glb
func View(rg *echo.Group) {
	rg.GET("/dispatch/views", view)
}

type dataT struct {
	ID         int    `json:"id"`
	View       string `json:"view"`
	ParentID   int    `json:"parent_id"`
	ViewLevel  int    `json:"view_level"`
	QansViewID int    `json:"qans_view_id"`
	Children   []dataT
}
type RespT struct {
	Code int     `json:"code"`
	Data []dataT `json:"data"`
}

//
func view(c echo.Context) error {
	url := host + c.Request().RequestURI
	req, err := http.NewRequest(c.Request().Method, url, c.Request().Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var re RespT
	err = json.Unmarshal(data, &re)
	if err != nil {
		fmt.Println(err)
		return err
	}

	optView := []dataT{}
	bigView := []dataT{}
	smallView := []dataT{}
	fmt.Println("...", re)
	for _, val := range re.Data {
		// fmt.Println("!!!!", key)
		// fmt.Println("....", val)
		if val.ViewLevel == 0 || val.ViewLevel == 1 {
			val.Children = []dataT{}
			optView = append(optView, val)
		}
		if val.ViewLevel == 2 {
			val.Children = []dataT{}
			bigView = append(bigView, val)
		}
		if val.ViewLevel == 3 {
			val.Children = []dataT{}
			smallView = append(smallView, val)
		}
	}
	fmt.Println("smallView...", smallView)

	return nil
}

// for (let i = 0; i < bigView.length; i++) {
//   for (let j = 0; j < smallView.length; j++) {
//     if (bigView[i].id === smallView[j].parent_id) {
//       bigView[i].children.push((smallView[j]))
//     }
//   }
// }

// for (let i = 0; i < optView.length; i++) {
//   for (let j = 0; j < bigView.length; j++) {
//     if (optView[i].id === bigView[j].parent_id) {
//       optView[i].children.push((bigView[j]))
//     }
//   }
// }
