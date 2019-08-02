def checksum(name)
  l = name.tr("-", "").chars
  h = Hash.new("")
  l.uniq.each { |x| h[l.count(x)] += x }
  h.sort.reverse.map { |_, v| v.chars.sort.join }.join[0, 5]
end

sum = 0
ARGF.read.each_line do |line|
  name, id, csum = line.match(/^([\w-]+)-(\d+)\[(\w+)\]$/).captures
  sum += id.to_i if checksum(name) == csum
end
puts sum
