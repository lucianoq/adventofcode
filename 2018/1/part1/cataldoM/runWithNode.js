const readline = require('readline');
const fs = require('fs');

const rl = readline.createInterface({
  input: fs.createReadStream('input.txt'),
  crlfDelay: Infinity
});

let arr = 0;

rl.on('line', (line) => {
  arr += (parseInt(line, 10));
});

rl.on('close', () => {
  console.log('arr', arr);
});
