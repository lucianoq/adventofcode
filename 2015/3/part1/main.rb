input = ARGF.read.strip
space, santa, c = { 0 => true }, 0, 1
input.chars.each do |ch|
  santa += 1 if ch == "^"
  santa -= 1 if ch == "v"
  santa -= input.size if ch == "<"
  santa += input.size if ch == ">"
  c += 1 if space[santa].nil?
  space[santa] = true
end
p c
