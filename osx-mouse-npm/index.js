var mouse = require('osx-mouse')();

mouse.on('move', function(x, y) {
	console.log('mouse', x, y);
});

mouse.on('left-down', function(x, y) {
	console.log('left-down', x, y);
});

mouse.on('left-up', function(x, y) {
	console.log('left-up', x, y);
});

mouse.on('left-drag', function(x, y) {
	console.log('left-drag', x, y);
});

mouse.on('right-up', function(x, y) {
	console.log('right-up', x, y);
});

mouse.on('right-down', function(x, y) {
	console.log('right-down', x, y);
});

mouse.on('right-drag', function(x, y) {
	console.log('right-drag', x, y);
});
