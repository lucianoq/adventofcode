ARGF.read.each_line do |line|
  name, id, _ = line.match(/^([\w-]+)-(\d+)\[(\w+)\]$/).captures
  puts id if name.chars.map { |c| c == "-" ? " " : ((c.ord - 97 + id.to_i) % 26 + 97).chr }.join == "northpole object storage"
end
