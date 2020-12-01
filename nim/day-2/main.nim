import strutils

echo "Starting day two..."

let i = readFile("input.txt")
var input: seq[int] = @[]
for s in split(i.strip(), ","):
  input.add(parseInt(s))

var i_1 = input

i_1[1] = 12
i_1[2] = 2

proc op(input: var seq[int]) =
  var i = 0
  while true:
    case input[i]:
      of 1:
        input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
        i += 4
      of 2:
        input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
        i += 4
      of 99:
        return
      else:
        echo "Invalid OP code"
        return

op(i_1)
echo "Part one: ", i_1[0]

proc op_2(input: seq[int], output: int): int =
  for i in countup(0,99):
    for j in countdown(99,0):
      var t = input
      t[1] = i
      t[2] = j
      op(t)
      if t[0] == output:
        return 100 * i + j
  return -1

echo "Part two: ", op_2(input, 19690720)
