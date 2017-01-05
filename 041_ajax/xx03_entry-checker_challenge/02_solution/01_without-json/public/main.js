var entry = document.querySelector('#entry');
var output = document.querySelector('h1');

entry.addEventListener('input', function(){
    console.log('ENTRY: ', entry.value);
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/api/check');
    xhr.send(entry.value);
    xhr.addEventListener('readystatechange', function(){
        if (xhr.readyState === 4 && xhr.status === 200) {
            var taken = xhr.responseText;
            console.log('TAKEN:', taken, '\n\n');
            if (taken == 'true') {
                output.textContent = 'Word Taken!';
            } else {
                output.textContent = '';
            }
        }
    });
});

