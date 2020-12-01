import strutils
import algorithm

echo "Starting day three..."

let i = open("input.txt")
# let i = open("test.txt")

var w_1 = i.readLine().split(",")
var w_2 = i.readLine().split(",")

type
  Field = array[-10000..10000, array[-10000..10000, char]]
  Point = tuple[x,y,d: int, steps: array[2,int]]
  Cross = seq[Point]

var board: Field
var crossings: Cross

board[0][0] = 'x'

proc trace_wire(wire: seq[string], board: var Field, self, other: char, crossings: var Cross)  =
  var x,y = 0
  var steps = 0
  for s in wire:
    # echo x, " ", y
    var d = s[0]
    var a = parseInt(s[1..len(s)-1])
    case d:
      of 'R':
        for i in x+1..x+a:
          if board[i][y] != self and board[i][y] != other and board[i][y] != 'x':
            board[i][y] = self
          elif board[i][y] == other:
            crossings.add((i,y,-1,[steps+(i-x),0]))
            board[i][y] = 'x'
          elif board[i][y] == 'x':
            for c in crossings.mitems:
              if c.x == i and c.y == y:
                c.steps[1] = steps+i-x
        x += a
      of 'L':
        for i in countdown(x-1,x-a):
          if board[i][y] != self and board[i][y] != other and board[i][y] != 'x':
            board[i][y] = self
          elif board[i][y] == other:
            crossings.add((i,y,-1,[steps+x-i,0]))
            board[i][y] = 'x'
          elif board[i][y] == 'x':
            for c in crossings.mitems:
              if c.x == i and c.y == y:
                c.steps[1] = steps+x-i
        x -= a
      of 'U':
        for i in y+1..y+a:
          if board[x][i] != self and board[x][i] != other and board[x][i] != 'x':
            board[x][i] = self
          elif board[x][i] == other:
            crossings.add((x,i,-1,[steps+i-y,0]))
            board[x][i] = 'x'
          elif board[x][i] == 'x':
            for c in crossings.mitems:
              if c.x == x and c.y == i:
                c.steps[1] = steps+i-y
        y += a
      of 'D':
        for i in countdown(y-1,y-a):
          if board[x][i] != self and board[x][i] != other and board[x][i] != 'x':
            board[x][i] = self
          elif board[x][i] == other:
            crossings.add((x,i,-1,[steps+y-i,0]))
            board[x][i] = 'x'
          elif board[x][i] == 'x':
            for c in crossings.mitems:
              if c.x == x and c.y == i:
                c.steps[1] = steps+y-i
        y -= a
      else:
        echo "Invalid directon: ", d
    steps += a

proc m_dist(p,q: (int,int)): int =
  result = abs(p[0]-q[0])+abs(p[1]-q[1])

proc cross_cmp(x,y: Point): int =
  if x.d < y.d: -1
  elif x.d == y.d: 0
  else: 1

proc cross_cmp_2(x,y: Point): int =
  if x.steps[0] + x.steps[1] < y.steps[0] + y.steps[1]: -1
  elif x.steps[0] + x.steps[1] == y.steps[0] + y.steps[1]: 0
  else: 1
 
trace_wire(w_1, board, '1', '2', crossings)
trace_wire(w_2, board, '2', '1', crossings)
trace_wire(w_1, board, '1', '2', crossings)

for c in crossings.mitems:
  c.d = m_dist((0,0),(c.x,c.y))

crossings.sort(cross_cmp)
echo "Part one: ", crossings[0].d
crossings.sort(cross_cmp2)
echo "Part two: ", crossings[0].steps[0] + crossings[0].steps[1]
i.close()
