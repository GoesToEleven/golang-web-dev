Use http.ListenAndServe on port 8080 of localhost with the default ServeMux.

Use http.HandleFunc to serve a func called dog.

Use io.WriteString to write this HTML:

<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">

Set the following header using res.Header().Set:

"Content-Type", "text/html; charset=utf-8"

What happens when you do not set the header?