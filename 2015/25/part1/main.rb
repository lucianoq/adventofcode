code, i, j = 20151125, 1, 1
begin
  i, j = i == 1 ? [j + 1, 1] : [i - 1, j + 1]
  code = code * 252533 % 33554393
end until i == 2947 && j == 3029
p code
