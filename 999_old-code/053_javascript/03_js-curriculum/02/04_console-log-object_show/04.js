var social = {
    title: 'list of social media links',
    sites: {
        facebook: 'www.facebook.com',
        google: 'www.google.com'
    },
    otras: [
        'facebook', 'www.facebook.com',
        'google', 'www.google.com'
    ]
};

for (var i = 0; i < social.otras.length; i++) {
    console.log(social.otras[i]);
}

console.log(social.sites.facebook);
console.log(social.sites.google);
console.log(social);