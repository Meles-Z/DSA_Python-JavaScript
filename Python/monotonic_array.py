def monotonic_array(array):

  if(len(array)<=1):
    return True
  left=array[0]
  right=array[len(array)-1]
  if(left>right):
    for i in range(len(array)-1):
      if(array[i+1]>array[i]):
        return False
  elif (left==right):
    for i in range(len(array)-1):
      if(array[i+1]!=array[i]):
        return False
  else:
    for i in range(len(array)-1):
      if(array[i+1]<array[i]):
        return False
  return True

print(monotonic_array([1,8,4]))