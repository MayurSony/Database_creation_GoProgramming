// //In this project we are going to create our own database , so that we can store our own data
// //This will be done using the go programming language.

// package main

// import (
// 	"database/sql/driver"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// 	"sync"
// 	"log"

// 	"github.com/blend/go-sdk/stringutil"
// 	"github.com/jcelliott/lumber"
// )

// const Version = "1.0.0"

// type(
// 	Logger interface{
// 		Fatal(string,...interface{})
// 		Error(string,...interface{})
// 		Warn(string,...interface{})
// 		Info(string,...interface{})
// 		Debug(string,...interface{})
// 		Trace(string,...interface{})
// 	}

// 	Driver struct{
// 		mutex sync.Mutex
// 		mutexes map[string]*sync.Mutex
// 		dir string
// 		log Logger
// 	}

// )

// type options struct{
//   Logger
// }

// func New(dir string, options *options)(*Driver , error){
// dir = filepath.Clean(dir)

// opts := options{}

// if options !=nil{
// 	opts = *options
// }

// if opts.Logger == nil{
// 	opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
// }

// driver := Driver{
// 	dir:dir,
// 	mutexes:make(map[string]*sync.Mutex),
// 	log: opts.Logger,
// }

// if _, os.Stat(dir); err == nil{
// 	opts.Logger.Debug("using '%s' (database already exists)\n",dir)
// 	return &driver,nil;
// }

// opts.Logger.Debug("creating the database at '%s'...\n",dir)
// return &driver,os.MkdirAll(dir,0755)

// }

// func (d* Driver) Write(collection,resource string, v interface{}) error{
// 	if collection == ""{
// 		return fmt.Errorf("missing collection -no piece to save record!")
// 	}
//     if resource == ""{
// 		return fmt.Errorf("missing resource - unable to save record (no name)!")
// 	}

// mutex := d.getOrCreateMutex(collection)
// mutex.Lock()
// defer mutex.unlock()

// fir := filepath.Join(d.dir,collection)
// fnlPath := filepath.Join(dir,resource+".json")
// tmpPath := fnlPath + ".tmp"

// if err := os.MkdirAll(dir,0755); err != nil{
// 	return err
// }
// b,err := json.MarshalIndent(v,"","\t")
// if err !=nil{
// 	return err
// }

// b=append(b,byte('\n'))

// if err := ioutil.WriteFile(tmpPath, b, 0644); err != nil {
//     return err
// }

// return os.Rename(tmpPath,fnlPath)
// }

// func (d *Driver) Read(collection ,resource string, v interface{}) error{
// 	if collection == ""{
// 		return fmt.Errorf("Missing collections - no place to save record!")
// 	}
// 	if resource ==""{
// 		return fmt.Errorf("Missing resource - unable to save the record!(no name)")
// 	}

// 	record := filepath.Join(d.dir,collection,resource)

// 	if _, err:= Stat(record); err!=nil{
// 		return err
// 	}
// 	b,e := ioutil.ReadFile(record + ".json")
// 	if err!=nil {
// 		return err
// 	}

// 	return json.Unmarshal(b,&v)

// }

// func (d *Driver) ReadAll(collection string)([]string, error){
//  if collection ==""{
// 	return nil, fmt.Errorf("Missing collection - unable to read !")
//  }
//  dir := filepath.Join(d.dir,collection)

//  if _, err := Stat(dir); err!=nil{
// 	return nil,err
//  }

//  file, _:= ioutil.ReadDir(dir)
//  var records [] string

//  for _,file :=range files{
// 	b,err := ioutil.ReadFile(filepath.Dir(dir,file.Name()))
// 	if err!=nil{
// 		return nil,err
// 	}
// 	records = append(records, string(b))
//  }
//  return records,nil

// }

// func (d *Driver) Delete(collection,resource string) error{
// 	path := filepath.Join(collection, resource)
// 	mutex := d.getOrCreateMutex(collection)
// 	mutex.Lock()

// 	defer mutex.unlock()
// 	dir := filepath.Join(d.dir,path)

// 	switch fi,err := stat(dir); {
// 	case fi==nil, err!=nil:
// 		return fmt.Errorf("unable to find the file or the directory named %v\n",path)

// 	case fi.Mode().IsDir():
// 		return os.RemoveAll(dir)
// 	}

// case fi.Mode().IsRegular():
// 	return os.RemoveAll(dir + ".json")
// }
// return nil

// func getOrCreateMutex() *sync.Mutex{

// 	d.mutex.Lock()
// 	defer d.mutex.unlock()
//     m,ok := d.mutexes[collection]

// 	if !ok{
// 		m=&sync.Mutex{}
// 		d.mutexes[collection]
// 	}
// }

// func Stat(path string)(fi os.FileInfo, err error){
// 	if fi,err = os.State(path);os.IsNotExist(err){
// 		fi,err = os.Stat(path + ".json")
// 	}
// 	return &driver,nil
// }

// //creating the struct part as we have to create the json database.

// type Address struct{
// 	City string
// 	State string
// 	Country string
// 	Pincode json.Number
// }

// type User struct{
// 	Name string
// 	Age json.Number
// 	Contact string
// 	Company string
// 	Address Address
// }

// func main(){
// 	dir := "./"

// 	db,err :=New(dir,nil)
// 	if err!= nil{
// 		fmt.Println("error",err)
// 	}

// 	employees := []User{
// 		{"john","23","123-456-789","Google",Address{"bangalore","karnataka","India","440011"}},
// 		{"Paul","24","123-456-789","Apple",Address{"bangalore","karnataka","India","440011"}},
// 		{"William","25","123-456-789","Microsoft",Address{"bangalore","karnataka","India","440011"}},
// 		{"George","26","123-456-789","Amazon",Address{"bangalore","karnataka","India","440011"}},
// 		{"Steve","27","123-456-789","Netflix",Address{"bangalore","karnataka","India","440011"}},
// 		{"Alexander","28","123-456-789","Meta",Address{"bangalore","karnataka","India","440011"}},

// 	}

// 	for _, value := range employees{
// 		db.Write("users",value.Name, User{
// 			Name:value.Name,
// 			Age:value.Age,
// 			Contact: value.Contact,
// 			Company: value.Company,
// 			Address:value.Address,
// 		})
// 	}

// 	records, err := db.ReadAll("users")
// 	if err != nil{
// 		fmt.Println("error",err)
// 	}
// 	fmt.Println(records)

// 	allusers :=[]User{}

// 	for _, f := range records{
// 		employeeFound := User{};
// 		if err := json.Unmarshal([]byte(f) ,&employeeFound); err != nil{
// 			fmt.Println("Error",err)
// 		}
// 		allusers = append(allusers, employeeFound)
// 	}
// 	fmt.Println((allusers))

// 	// if err := db.Delete("user","john"); err !=nil{
// 	// 	fmt.Println("Error",err)
// 	// }
// 	// if err := db.Delete("user",""); err !=nil{
// 	// 	fmt.Println("Error",err)

// 	// }
// }

// In this project, we are going to create our own database, so that we can store our own data.
// This will be done using the Go programming language.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
)

const Version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)

	opts := Options{}

	if options != nil {
		opts = *options
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger(lumber.INFO)
	}

	driver := &Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("using '%s' (database already exists)\n", dir)
		return driver, nil
	}

	opts.Logger.Debug("creating the database at '%s'...\n", dir)
	return driver, os.MkdirAll(dir, 0755)
}

func (d *Driver) Write(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("missing collection - no place to save record!")
	}
	if resource == "" {
		return fmt.Errorf("missing resource - unable to save record (no name)!")
	}

	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)
	fnlPath := filepath.Join(dir, resource+".json")
	tmpPath := fnlPath + ".tmp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}

	b = append(b, byte('\n'))

	if err := ioutil.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}

	return os.Rename(tmpPath, fnlPath)
}

func (d *Driver) Read(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("missing collection - no place to save record!")
	}
	if resource == "" {
		return fmt.Errorf("missing resource - unable to save record (no name)!")
	}

	record := filepath.Join(d.dir, collection, resource+".json")

	if _, err := os.Stat(record); err != nil {
		return err
	}
	b, err := ioutil.ReadFile(record)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &v)
}

func (d *Driver) ReadAll(collection string) ([]string, error) {
	if collection == "" {
		return nil, fmt.Errorf("missing collection - unable to read!")
	}
	dir := filepath.Join(d.dir, collection)

	if _, err := os.Stat(dir); err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var records []string

	for _, file := range files {
		b, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		records = append(records, string(b))
	}
	return records, nil
}

func (d *Driver) Delete(collection, resource string) error {
	path := filepath.Join(collection, resource)
	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, path)

	switch fi, err := os.Stat(dir); {
	case fi == nil, err != nil:
		return fmt.Errorf("unable to find the file or the directory named %v\n", path)

	case fi.Mode().IsDir():
		return os.RemoveAll(dir)

	case fi.Mode().IsRegular():
		return os.Remove(dir + ".json")
	}
	return nil
}

func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	m, ok := d.mutexes[collection]

	if !ok {
		m = &sync.Mutex{}
		d.mutexes[collection] = m
	}
	return m
}

// Creating the struct part as we have to create the JSON database.

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	employees := []User{
		{"john", "23", "123-456-789", "Google", Address{"bangalore", "karnataka", "India", "440011"}},
		{"Paul", "24", "123-456-789", "Apple", Address{"bangalore", "karnataka", "India", "440011"}},
		{"William", "25", "123-456-789", "Microsoft", Address{"bangalore", "karnataka", "India", "440011"}},
		{"George", "26", "123-456-789", "Amazon", Address{"bangalore", "karnataka", "India", "440011"}},
		{"Steve", "27", "123-456-789", "Netflix", Address{"bangalore", "karnataka", "India", "440011"}},
		{"Alexander", "28", "123-456-789", "Meta", Address{"bangalore", "karnataka", "India", "440011"}},
	}

	for _, value := range employees {
		err := db.Write("users", value.Name, value)
		if err != nil {
			fmt.Println("error", err)
		}
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(records)

	allUsers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("error", err)
		}
		allUsers = append(allUsers, employeeFound)
	}
	fmt.Println(allUsers)

	// if err := db.Delete("users", "john"); err != nil {
	// 	fmt.Println("error", err)
	// }
	// if err := db.Delete("users", ""); err != nil {
	// 	fmt.Println("error", err)
	// }
}
