require "colorize"

class Virus
  UP, RIGHT, DOWN, LEFT = 0, 1, 2, 3

  def initialize(space)
    @space = space
    @pos = [0, 0]
    @direction = UP
  end

  def burst
    @space[@pos] ? self.right : self.left
    self.toggle
    self.forward
  end

  def right
    puts "right"
    @direction = (@direction + 1) % 4
  end

  def left
    puts "left"
    @direction = (@direction + 3) % 4
  end

  def forward
    puts "forward"
    case @direction
    when UP
      @pos = [@pos[0] - 1, @pos[1]]
    when RIGHT
      @pos = [@pos[0], @pos[1] + 1]
    when DOWN
      @pos = [@pos[0] + 1, @pos[1]]
    when LEFT
      @pos = [@pos[0], @pos[1] - 1]
    end

    @space[@pos] = false unless @space.key?(@pos)
  end

  def toggle
    puts "toggling"
    @space[@pos] = !@space[@pos]
  end

  def infected
    @space.keep_if { |k, v| v }.size
  end

  def print_virus
    x_min, x_max = @space.keys.map { |x| x[0] }.minmax
    y_min, y_max = @space.keys.map { |x| x[1] }.minmax

    puts "--------------"

    (x_min..x_max).each do |i|
      (y_min..y_max).each do |j|
        # (-4..4).each do |i|
        #   (-4..4).each do |j|
        if @space[[i, j]] == nil
          print "N".green
          next
        end

        if @pos == [i, j]
          print (@space[[i, j]] ? "#" : ".").red
        else
          print @space[[i, j]] ? "#" : "."
        end
      end
      puts
    end

    puts "--------------"
  end
end

def main
  input = ARGF.readlines.map(&:strip)
  rows, cols = input.size, input[0].size
  c_x, c_y = rows / 2, cols / 2

  space = {}
  rows.times do |i|
    cols.times do |j|
      space[[i - c_x, j - c_y]] = input[i][j] == "#" ? true : false
    end
  end

  v = Virus.new(space)

  v.print_virus

  10000.times do
    v.burst
  end

  v.print_virus
  puts v.infected
end

main
