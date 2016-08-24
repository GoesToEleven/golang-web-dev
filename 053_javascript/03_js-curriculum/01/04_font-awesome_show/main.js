var socialSites = ['https://www.facebook.com/',
    'https://twitter.com/',
    'https://www.linkedin.com/',
    'https://www.pinterest.com/'];

for (var i = 0; i < socialSites.length; i++) {
    console.log(socialSites[i]);
}

console.log('--------------');

var output = '<ul>';
for (var i = 0; i < socialSites.length; i++) {
    output += '\n';
    output += ('<li>' + socialSites[i] + '</li>');
}
output += '\n</ul>';
console.log(output);

console.log('--------------');

var output = '<ul>';
for (var i = 0; i < socialSites.length; i++) {
    output += '\n';
    output += '<li>';
    output += ('<a href="' + socialSites[i] + '">');
    output += (socialSites[i]);
    output += '</a>';
    output += '</li>';
}
output += '\n</ul>';
console.log(output);


console.log('--------------');

var myIcon = ['fa fa-camera-retro fa-5x',
    'fa fa-home fa-fw fa-5x',
    'fa fa-cog fa-spin fa-5x',
    'fa fa-arrow-circle-down fa-5x'];

var output = '<ul>';
for (var i = 0; i < socialSites.length; i++) {
    output += '\n';
    output += '<li>';
    output += ('<a href="' + socialSites[i] + '">');
    output += '<i class="' + myIcon[i] +'"></i>';
    output += '</a>';
    output += '</li>';
}
output += '\n</ul>';
console.log(output);

// inject

var newDiv = document.createElement("div");
document.querySelector('p').appendChild(newDiv);
newDiv.innerHTML = output;

//document.querySelector('body').innerHTML = output;




