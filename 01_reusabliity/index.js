const toBeFiltered = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15];

function filter(numbers, condition) {
    let result = [];
    for (let i = 0; i < numbers.length; i++) {
        if (condition(numbers[i])) {
            result.push(numbers[i]);
        }
    }

    return result;
}

const isDivisibleBy3 = (num) => num % 3 === 0;

console.log(filter(toBeFiltered, isDivisibleBy3));
