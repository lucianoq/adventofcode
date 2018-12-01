const readline = require('readline');
const fs = require('fs');

let curr = 0;
let inputArray = [];
let freqArray = [];

freqArray.push(curr);
let found = false;

readline.createInterface({
  input: fs.createReadStream('input.txt'),
  crlfDelay: Infinity
}).on('line', (line) => {
  curr = (parseInt(line, 10));
  inputArray.push(curr);
}).on('close', () => {
  curr = 0;
  externalWhile:while(!found){
    for (let i = 0;  i < inputArray.length; i++){
      curr += inputArray[i];

      if (freqArray.indexOf(curr) > -1) {
        console.log(curr);
        found = true;
        continue externalWhile;
      }
      freqArray.push(curr);
    }
  }
});
