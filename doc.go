//
//A RESTful style web-services framework for the Go language. </br>
//Creating services in Go is straight forward, Restful? takes this a step further by adding a layer that 
//makes tedious tasks much more automated and avoids regular pitfalls. <br/> 
//This gives you the opportunity to focus more on the task at hand... minor low-level http handling.<br/>
//
//
//Example usage below:
//
//	package main
//	import (
// 	   "github.com/duzy/go-restful"
//	        "http"
//	)
//	func main() {
//	    restful.RegisterService(new(HelloService)) //Register our service
//	    http.Handle("/",restful.Handle())    
// 	   http.ListenAndServe(":8787",nil)
//	}
//
//	//Service Definition
//	type HelloService struct {
//	    restful.RestService `root:"/tutorial/"`
//	    helloWorld  restful.EndPoint `method:"GET" path:"/hello-world/" output:"string"`
//	    sayHello    restful.EndPoint `method:"GET" path:"/hello/{name:string}" output:"string"`
//	}
//	func(serv HelloService) HelloWorld() string{
// 	   return "Hello World"
//	}
//	func(serv HelloService) SayHello(name string) string{
//	    return "Hello " + name
//	}
package restful
