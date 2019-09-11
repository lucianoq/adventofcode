puts $stdin.readlines.map { |x| x.split }.count { |x| x.uniq.size == x.size }
