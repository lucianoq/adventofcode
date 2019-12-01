#!/usr/bin/env ruby

puts ARGF.readlines.map(&:to_i).map { |x| x / 3 - 2 }.map { |mass|
  fuel = mass / 3 - 2
  while fuel > 0
    mass += fuel
    fuel = fuel / 3 - 2
  end
  mass
}.inject(:+)
