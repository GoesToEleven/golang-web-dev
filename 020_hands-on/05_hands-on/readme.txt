Building upon the code from the previous problem, instead of ioutil.ReadAll we are now going to use bufio.Scanner. Package bufio has the Scanner type. The Scanner type "provides a convenient interface for reading data". When you have a Scanner type, you can call the Scan method on it. Successive calls to the Scan method will step through the lines of the data. The Scanner type also has a Text method. When you call this method, you will be given the text from the current line. Here is how you could use it:

scanner := bufio.NewScanner(conn)
for scanner.Scan() {
	ln := scanner.Text()
	fmt.Println(ln)
}

Use this code to READ from an incoming connection and print the incoming text to standard out (the terminal).