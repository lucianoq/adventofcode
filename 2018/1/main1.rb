#!/usr/bin/ruby

puts $stdin.readlines.map(&:strip).delete_if(&:empty?).map(&:to_i).inject(:+)