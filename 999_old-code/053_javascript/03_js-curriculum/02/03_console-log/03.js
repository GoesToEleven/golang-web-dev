var info = {
    name: "Ray Villalobos",
    age: 48,
    jarjar: false,
    movies: ['Star Wars', 'Star Trek', 'Galaxy Quest'],
};

console.log(info.name);
console.log(info.age);

if (info.jarjar) {
    console.log('Love Jar Jar');
} else {
    console.log('Han should have shot Jar Jar');
}

console.log('');

for (var i = 1; i<info.movies.length; i++) {
    console.log(info.movies[i]);
}

info.company = 'Lynda.com';

console.log(info);