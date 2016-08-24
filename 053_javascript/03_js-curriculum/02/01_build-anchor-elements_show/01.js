var links = ['www.google.com', 'google',
    'www.reddit.com', 'reddit',
    'www.news.google.com', 'news'];

var nav = '<ul>';
nav += '\n';
for (var i = 0; i < links.length; i++) {
    nav += '<li>';
    // <a href="www.google.com">google</a>
    nav += '<a href="' + links[i] + '">';
    i++;
    nav += links[i];
    nav += '</a></li>\n'
}
nav += '</ul>\n';

console.log(nav);



