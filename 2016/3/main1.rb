count = 0
ARGF.read.each_line do |line|
  a, b, c = line.split.map(&:to_i).sort
  count += 1 if a + b > c
end
puts count
