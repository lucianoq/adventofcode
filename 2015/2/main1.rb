#!/usr/bin/env ruby

total = 0

ARGF.each do |line|
  l, w, h = line.split("x").map(&:to_i)
  smallest = [l * w, l * h, w * h].min
  total += 2 * (l * w + l * h + w * h) + smallest
end

puts total
