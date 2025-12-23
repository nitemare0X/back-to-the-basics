const toBeFiltered = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];


function filter(numbers, condition) {
    let result = [];
    for (let i = 0; i < numbers.length; i++) {
        if (condition(numbers[i])) {
            result.push(numbers[i]);
        }
    }

    return result;
}

function map(numbers, transform) {
    let result = [];
    for (let i = 0; i < numbers.length; i++) {
        result.push(transform(numbers[i]));
    }

    return result;
}

const isEven = (num) => num % 2 === 0;
const square = (num) => num ** 2;

console.log(map(filter(toBeFiltered, isEven), square));
