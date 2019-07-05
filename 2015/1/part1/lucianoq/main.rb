#!/usr/bin/env ruby

v = { "(" => +1, ")" => -1 }
puts ARGF.read.chars.map { |c| v[c] }.reduce(&:+)
