count = 0
ARGF.read.each_line.each_slice(3) do |x|
  triple = x.map(&:strip).join(" ").split.map(&:to_i)
  3.times do |i|
    a, b, c = triple.values_at(*(i...9).step(3).to_a).sort
    count += 1 if a + b > c
  end
end
puts count
