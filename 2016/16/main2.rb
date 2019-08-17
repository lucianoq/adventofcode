DISK_SIZE = 35651584
INPUT = "10001001100000001"

def dragon(a)
  a + "0" + a.chars.reverse.join.tr("01", "10")
end

x = INPUT
x = dragon(x) while (x.length < DISK_SIZE)
x = x[0...DISK_SIZE].chars
x = x.each_slice(2).map { |s| s[0] == s[1] ? "1" : "0" } while x.length.even?
puts x.join
