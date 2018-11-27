#!/usr/bin/ruby

$stdin.readlines.map(&:strip).delete_if(&:empty?).map{|x|x.split('')}.each do |a|
	puts a.push(a[0]).each_cons(2).map{|p| p[0]==p[1] ? p[0].to_i : 0}.inject(:+)
end