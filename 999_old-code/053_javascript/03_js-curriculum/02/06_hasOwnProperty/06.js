var user = {
    fName: 'John',
    lName: 'Doe',
    links: {
        blog: 'www.blogger.com',
        facebook: 'www.facebook.com'
    }
};

for (var i in user.links) {
    console.log(i, ' - ', user.links[i]);
}

console.log('-------------');

// safety check
for (var i in user.links) {
    if (user.links.hasOwnProperty(i)) {
        console.log(i, ' - ', user.links[i]);
    }
}
