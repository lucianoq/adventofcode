$stdin.readlines.map(&:to_i).combination(3).each do |x, y, z|
  if x + y + z == 2020
    puts x * y * z
    exit
  end
end
