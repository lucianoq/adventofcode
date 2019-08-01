#!/usr/bin/env ruby

pin, x = "", 5
ARGF.read.each_line do |line|
  line.chars.each do |c|
    x -= x - 3 < 1 ? 0 : 3 if c == "U"
    x += x + 3 > 9 ? 0 : 3 if c == "D"
    x -= 1 unless [1, 4, 7].include? x if c == "L"
    x += 1 unless [3, 6, 9].include? x if c == "R"
  end
  pin += x.to_s
end
puts pin
