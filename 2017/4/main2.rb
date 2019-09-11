puts $stdin.readlines.map { |line|
       line.split.map do |word|
         word.chars.sort.join
       end
     }.count { |word_list|
       word_list.uniq.size == word_list.size
     }
