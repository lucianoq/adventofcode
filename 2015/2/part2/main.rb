#!/usr/bin/env ruby

total = 0

ARGF.each do |line|
  l, w, h = line.split("x").map(&:to_i)
  a, b = [l, w, h].sort.take 2
  ribbon = 2 * (a + b)
  volume = l * w * h
  total += ribbon + volume
end

puts total
