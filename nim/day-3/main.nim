import strutils
import algorithm

echo "Starting day three..."

let i = open("input.txt")

var w_1 = i.readLine().split(",")
var w_2 = i.readLine().split(",")

type
  Field = array[-10000..10000, array[-10000..10000, char]]
  Point = tuple[x,y,d: int]
  Cross = seq[Point]

var board: Field
var crossings: Cross

board[0][0] = 'x'

proc trace_wire(wire: seq[string], board: var Field, self, other: char, crossings: var Cross)  =
  var x,y = 0
  for s in wire:
    # echo x, " ", y
    var d = s[0]
    var a = parseInt(s[1..len(s)-1])
    case d:
      of 'R':
        for i in x+1..x+a:
          if board[i][y] != self and board[i][y] != other:
            board[i][y] = self
          elif board[i][y] == other:
            crossings.add((i,y,-1))
            board[i][y] = 'x'
        x += a
      of 'L':
        for i in x-a..x-1:
          if board[i][y] != self and board[i][y] != other:
            board[i][y] = self
          elif board[i][y] == other:
            crossings.add((i,y,-1))
            board[i][y] = 'x'
        x -= a
      of 'U':
        for i in y+1..y+a:
          if board[x][i] != self and board[x][i] != other:
            board[x][i] = self
          elif board[x][i] == other:
            crossings.add((x,i,-1))
            board[x][i] = 'x'
        y += a
      of 'D':
        for i in y-a..y-1:
          if board[x][i] != self and board[x][i] != other:
            board[x][i] = self
          elif board[x][i] == other:
            crossings.add((x,i,-1))
            board[x][i] = 'x'
        y -= a
      else:
        echo "Invalid directon: ", d

proc m_dist(p,q: (int,int)): int =
  result = abs(p[0]-q[0])+abs(p[1]-q[1])

proc cross_cmp(x,y: Point): int =
  if x.d < y.d: -1
  elif x.d == y.d: 0
  else: 1
 
trace_wire(w_1, board, '1', '2', crossings)
trace_wire(w_2, board, '2', '1', crossings)

for c in crossings.mitems:
  c.d = m_dist((0,0),(c.x,c.y))

crossings.sort(cross_cmp)

echo "Part one: ", crossings[0].d
i.close()
