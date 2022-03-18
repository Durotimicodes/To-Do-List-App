package functionality

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

//create a custome type Data
var Struct Data

//My Data structure
type Data struct {
	List   string `json:"list"`
	Status bool   `json:"status"`
}

var DataList []Data

//first function to be executed
func init() {
	Read()
}

//READ USING IOUTIL
func Read() {
	//what file do you want me to read ?
	data, err := ioutil.ReadFile("data.csv")
	if err != nil {
		fmt.Println(err)
	}
	//after reading the file covert data to struct - string
	json.Unmarshal(data, &DataList)
}

func Write(s []Data) {
	//coverting struct to json data || coverting string to byte
	json_Data, err := json.Marshal(&DataList)
	if err != nil {
		panic(err)
	}

	//write into file
	er := ioutil.WriteFile("data.csv", json_Data, 0666)
	if er != nil {
		fmt.Println(err)
	}

}

//FUNCTIONALITY
func (a *Data) AddFunc(s string) string {
	r := Data{
		List:   s,
		Status: false,
	}

	//Read() the string s
	DataList = append(DataList, r)
	Write(DataList)
	return s + "added successfully"
}

func ListData() string {
	//Read file Read()

	for i, v := range DataList {
		if !v.Status {
			fmt.Println(i+1, v.List)
		}
	}

	return "List successfuly printed"
}

func (a *Data) DoneTask(s string) string {
	y, _ := strconv.Atoi(s)

	//Read()
	for i := range DataList {
		if y == (i + 1) {
			DataList[i].Status = true
		}
	}
	Write(DataList)

	return "Successful"
}

func (a *Data) UndoneItem(s string) string {
	y, _ := strconv.Atoi(s)
	//Read()
	for i := range DataList {
		if y == (i+1) && DataList[i].Status {
			DataList[i].Status = false
		}
	}
	Write(DataList)
	return "Tasked removed from list Successful"
}

func Cleanup() string {
	var d []Data
	r := Data{}
	//Read()
	for _, val := range DataList {
		if !val.Status {
			r.List = val.List
			r.Status = false
			d = append(d, r)
		}
	}
	Write(d)

	return "Successful"
}
