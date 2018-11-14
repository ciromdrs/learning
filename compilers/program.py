def func(x):
  print "executando func(",x,")"
  y = x
  if y%2 == 0:
    y /= 2
  else:
    y += 1
  return y

func(4)
func(9)
