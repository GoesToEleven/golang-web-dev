Building upon the code from the previous exercise:

In that previous exercise, we WROTE to the connection.

Now I want you to READ from the connection.

You can READ and WRITE to a net.Conn (a connection implements both the reader and writer interface).

Use ioutil.ReadAll to read from the connection. This will give you a []byte. You can convert a []byte to a string like this: string([]byte). You can then print that string to standard out (the terminal).

See this code: https://play.golang.org/p/LSsrFCOJR0

Launch your TCP server and in your web browser, visit localhost:8080. Now go back and look at your terminal. What did you find?

Can you answer the question as to why "I see you connected" is never written?