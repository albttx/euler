
# The sum of the squares of the first ten natural numbers is,
#
# 1^2 + 2^2 + ... + 10^2 = 385
# The square of the sum of the first ten natural numbers is,
#
# (1 + 2 + ... + 10)^2 = 552 = 3025
# Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
#
# Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

MAX = 100

def sum_squares()
	sum = 0
	(0..MAX).each do |n|
		sum += n ** 2
	end
	return sum
end

def squares_sum()
	ret = 0
	n = (MAX * (MAX + 1)) / 2
	ret = n ** 2
	return ret
end

puts squares_sum - sum_squares
