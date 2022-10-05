package utils

import (
	"encoding/json" 
	"errors" 
	"fmt" 
	"html/template" 
	"log"
	"net/http"
	"os"
	"strings"
	
	"github.com/irononet/chatapp/data" 
)

type Configuration struct{
	Address			string 
	ReadTimeout		int64 
	WriteTimeout	int64
	Static			string 
}

var config Configuration 
var logger *log.Logger 

var ROOT_DIR string = "./"

// function for printing to stdout 
func Pstdout(a ...interface{}){
	fmt.Println(a)
}

func init(){
	loadConfig() 
	file, err := os.OpenFile("chatapp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) 
	if err != nil{
		log.Fatalln("Failed to open log file", err)
	}

	logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

func LoadConfiguration(file string) Configuration{
	var config Configuration 
	configFile, err := os.Open(file) 
	if err != nil{
		fmt.Println(err.Error())
	}
	defer configFile.Close() 
	jsonParser := json.NewDecoder(configFile) 
	jsonParser.Decode(&config) 
	return config 
}

func loadConfig(){
	file, err := os.Open(ROOT_DIR + "config.json") 
	if err!= nil{
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file) 
	config = Configuration{} 
	err = decoder.Decode(&config) 
	if err != nil{
		log.Fatalln("Cannot get configuration from file", err)
	}
}

func ErrorMessage(writer http.ResponseWriter, request *http.Request, msg string){
	url := []string{"/err?msg=", msg} 
	http.Redirect(writer, request, strings.Join(url, ""), http.StatusFound)
}

// Checks if the user is logged in and has a session, if not err is not nil 
func Session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error){
	cookie, err := request.Cookie("_cookie") 
	if err == nil{
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok{
			err = errors.New("Invalid Session")
		}
	}
	return 
}

// parse HTML templates 
// pass in a list of file names, and get a template 
func ParseTemplateFiles(filenames ...string) (t *template.Template){
	var files []string 
	t = template.New("layout") 
	for _, file := range filenames{
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return 
}

func GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string){
	var files []string 
	for _, file := range filenames{
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...)) 
	templates.ExecuteTemplate(writer, "layout", data) 
}

// for logging 
func Info(args 	...interface{}){
	logger.SetPrefix("INFO ") 
	logger.Println(args...)
}

func Danger(args ...interface{}){
	logger.SetPrefix("ERROR ") 
	logger.Println("args...")
}

func Warning(args ...interface{}){
	logger.SetPrefix("WARNING ") 
	logger.Println(args...) 
}

func Version() string{
	return "0.0.1"
}