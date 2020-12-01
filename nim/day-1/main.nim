import strutils

proc calcFuel(mass: int): int =
  result = int(mass/3)-2

echo("Starting day one..")
let i = open("input.txt")

var line = ""
var fuel = 0
while readLine(i, line):
  fuel += calcFuel(parseInt(line))

echo "Required fuel: ", fuel
echo "Part 2..."
i.setFilePos(0)

fuel = 0
while readLine(i, line):
  var m_fuel = calcFuel(parseInt(line))
  var f = calcFuel(m_fuel)
  while f >= 0:
    m_fuel += f
    f = calcFuel(f)
  fuel += m_fuel

echo "Actually required fuel: ", fuel
i.close()
