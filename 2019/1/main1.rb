#!/usr/bin/env ruby

puts ARGF.readlines.map(&:to_i).map { |x| x / 3 - 2 }.inject(:+)
