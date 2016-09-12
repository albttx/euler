
# 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
#
# What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

MULT = 20

def is_smallest_multiple(nb)
	(2..MULT).each do |n|
		return false if nb % n != 0
	end
	return true
end

i = 9
loop do
	break if is_smallest_multiple(i)
	i += 1
end

puts i

# Discover after, i should use PPCM !
