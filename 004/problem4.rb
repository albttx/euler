
# A palindromic number reads the same both ways.
# The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
# Find the largest palindrome made from the product of two 3-digit numbers.

MIN = 100 * 100 #  10 000
MAX = 999 * 999 # 998 007

def is_palindrome(nb)
	str = nb.to_s
	if str.reverse == str
		return true
	else
		return false
	end
end

i = 100
res = []

while i < 999
	j = 100
	while j < 999
		n = i * j
		res.push(n) if is_palindrome(n)
		j += 1
	end
	i += 1
end

res.sort!
puts res.last
