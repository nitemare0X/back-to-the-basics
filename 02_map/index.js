//const toBeFiltered = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15];
const toBeDoubled = [1, 2, 3, 4];
const toAddOne = [5, 10, 15];
const toBeSquared = [2, 4, 6];


function map(numbers, transform) {
    let result = [];
    for (let i = 0; i < numbers.length; i++) {
        result.push(transform(numbers[i]));
    }

    return result;
}

//const isDivisibleBy3 = (num) => num % 3 === 0;
const double = (num) => num * 2;
const addOne = (num) => num + 1;
const square = (num) => num ** 2;

console.log([1, 2, 3].map(double));
console.log(map(toBeDoubled, double));
console.log(map(toAddOne, addOne));
console.log(map(toBeSquared, square));
