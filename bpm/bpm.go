package bpm
import(
	 "log"
	"github.com/go-resty/resty/v2"
)
func Get_test() []byte{
  client:=resty.New()
  resp, err := client.R().
       Get("https://jsonplaceholder.typicode.com/posts")
  if err != nil {
	 log.Fatal(err)//印出error值
  }
	   
   return resp.Body()
}

func Post_test(body []byte) []byte{
	client :=resty.New()
	resp, err :=client.R().
	  SetHeader("Content-Type","application/json").
	  SetBody(body).
	  Post("https://jsonplaceholder.typicode.com/posts")
    if err !=nil{
		log.Fatal(err)
	}
	return resp.Body()
}