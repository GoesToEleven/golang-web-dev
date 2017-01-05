var entry = document.querySelector('#entry');
var output = document.querySelector('h1');

entry.addEventListener('input', function(){

    // value in form text input field
    console.log('ENTRY: ', entry.value);

    // AJAX SEND
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/api/check');
    xhr.send(JSON.stringify({'Name':entry.value}));
    console.log("JSON SENT:", JSON.stringify({'Name':entry.value}));

    // AJAX RECEIVE
    xhr.addEventListener('readystatechange', function(){
        if (xhr.readyState === 4 && xhr.status === 200) {
            console.log("JSON RCVD:", xhr.responseText);
            var taken = JSON.parse(xhr.responseText);
            console.log('TAKEN:', taken, '\n\n');
            if (taken == 'true') {
                output.textContent = 'Word Taken!';
            } else {
                output.textContent = '';
            }
        }
    });

});
