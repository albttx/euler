
# The prime factors of 13195 are 5, 7, 13 and 29.
#
# What is the largest prime factor of the number 600851475143 ?

@nb = 60085147514
@mult = 1

print "#{@nb} ="
def prime_factor(n)
	return if @mult == @nb
	i = 2
	i += 1 while (n % i != 0)
	print " #{i} *"
	@mult = @mult * i
	prime_factor(n / i)
end

prime_factor(@nb)
