$stdin.readlines.map(&:to_i).combination(2).each do |x, y|
  if x + y == 2020
    puts x * y
    exit
  end
end
