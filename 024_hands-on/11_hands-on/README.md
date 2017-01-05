# Starting with the code in the "starting-files" folder:

## wire this program up so that it works

ParseGlob in an init function

Use HandleFunc for each of the routes

Combine apply & applyProcess into one func called "apply"

Inside the func "apply", use this code to create the logic to respond differently to a POST method and a GET method

``` go
if req.Method == http.MethodPost {
    		// code here
    		return
    	}
```