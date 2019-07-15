require "json"

def parseObj(obj)
  is_red = false

  items = []
  obj.each do |k, v|
    is_red = true if v == "red"
    items << v
  end

  return 0 if is_red

  total = 0
  items.each do |x|
    case x
    when Hash
      total += parseObj x
    when Array
      total += parseArray x
    when Numeric
      total += x
    end
  end

  return total
end

def parseArray(arr)
  total = 0
  arr.each do |x|
    case x
    when Hash
      total += parseObj x
    when Array
      total += parseArray x
    when Numeric
      total += x
    end
  end

  return total
end

json = File.read("input")
hash = JSON.parse(json)
puts(parseObj(hash))
