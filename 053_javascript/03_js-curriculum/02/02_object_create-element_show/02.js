var myObj = {
    fName: 'Todd',
    lName: 'McLeod',
    happy: true,
    secret: 42,
    movies: ['Alien', 'Heat', 'Hmm', 'Other']
};

var myParagraph = myObj.fName + ' ' + myObj.lName + ' is ';

if (myObj.happy) {
    myParagraph += 'happy.';
} else {
    myParagraph += 'not happy.';
}

myParagraph += ' The secret to life is ' + myObj.secret + '.';
myParagraph += ' The movies that ' + myObj.fName + ' likes are ';

for (var i = 0; i < myObj.movies.length; i++) {
    myParagraph += myObj.movies[i];
    if (i != myObj.movies.length - 1) {
        myParagraph += ', ';
    } else {
        myParagraph += '.';
    }
}

console.log(myParagraph);
console.log(myObj);

// inject
var body = document.querySelector('body');
var newH = document.createElement('h1');
body.appendChild(newH);
newH.innerHTML = myParagraph;

