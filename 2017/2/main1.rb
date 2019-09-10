sum = 0
$stdin.each_line.reject{|x|x.strip.empty?}.each do |line|
  min, max = line.split.map(&:to_i).minmax
  sum += max-min
end
puts sum
