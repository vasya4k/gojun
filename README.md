An example of a simple app which provides RESTful API to "show bgp neighbour" data via NETCONF. More info netopscasts.com/showbgp/ .

On startup, it reads it's config and a list of devices. After that, it will be waiting for an HTTP GET request to "ip:port/bgpnei" "curl -X GET -k http://127.0.0.1:8888/bgpnei" endpoint. 

When it receives a request, it will open concurrent connections to all routers from the list to collect BGP state for each neighbour. 
Then it will transform the data and return it as a JSON object. The data received from an API can then be used by a front-end or by any other tool. 

This app is just a demo of what you can do with NECONF. I am going to be updating this each time as I write a new post at netopscasts.com

Update.

I have added a simple web interface based on VueJS and CoreUI. More details in a blogpost at vasya4k.net/gojun2. This is just an example of how easy it is nowadays to create a simple but useful web interface. Here is a screenshot: 

![Alt text](gojun.png?raw=true "Gojun screenshot") 

The Go code in this repo is using https://github.com/golang/dep dependency management tool. You can install dependencies by running dep ensure command inside gojun dir. 

The web interface should work once you have cloned the repository all the required files reside in "ui/dist" directory. If you want to rebuild the web interface, go to ui directory and run "npm install" first and then  "npm run build" commands. 

