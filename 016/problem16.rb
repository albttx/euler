
nb = 2 ** 1000
res = 0

nb.to_s.each_char {|c| res += c.to_i}

puts res
