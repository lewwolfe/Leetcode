// Example 1:
// Input: dividend = 10, divisor = 3
// Output: 3
// Explanation: 10/3 = 3.33333.. which is truncated to 3.

// Example 2:
// Input: dividend = 7, divisor = -3
// Output: -2
// Explanation: 7/-3 = -2.33333.. which is truncated to -2.

func divide(dividend int, divisor int) int {
	var result int = 0
	var sum int = 0
    var flipped bool = false


	//Change negative divisor into positive and flip back later
	if divisor < 0 {
		divisor *= -1
        flipped = true
	}

	if dividend < 0 && divisor > 0 {
		dividend *= -1
        flipped = !flipped
	}

	//keep adding to the sum until we reach divisor
	for (sum + divisor) <= dividend {
		sum += divisor
		result += 1
	}

	//flip to a negative integer
    if flipped {
        result *= -1
    }

    //Keep bounds
    if result > 2147483647 {
        result = 2147483647
    } else if result < -2147483648 {
        result = 2147483648
    }

    return result
}