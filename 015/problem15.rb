
SIZE = (ARGV[0]) ? ARGV[0].to_i : 21

@map = Array.new(SIZE) { Array.new(SIZE, 0) }

x, y = 0, 0

@map[0][0] = 1
while y < SIZE
	x = 0
	while x < SIZE
		@map[y][x] += @map[y-1][x] if @map[y-1] != nil
		@map[y][x] += @map[y][x-1] if @map[y][x-1]
		x += 1
	end
	p @map[y]
	y += 1
end



puts @map[SIZE-1][SIZE-1]
