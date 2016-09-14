
# Special Pythagorean triplet

LIMIT = 10000

def factor_pairs(nb)
	ret = []
	n = 1
	while n < nb / n

		if nb % n == 0
			ret << [n , nb / n]
		end
		n += 1
	end
	return ret
end

def dickson_method()
	(2..LIMIT).each do |r|
		factor_pairs((r ** 2) / 2).each do |s, t|
			if (r ** 2) == (2 * s * t)
				x = r + s
				y = r + t
				z = r + s + t
        		if x + y + z == 1000
		            print "Solution is : "
        			puts "(x = #{x} +  y=#{y} + z=#{z}) == 1000"
					puts "Product is #{x * y * z}"
		            exit
       			 end
			end
		end
	end
end

dickson_method()
