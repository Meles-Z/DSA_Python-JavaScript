def rotate_array(array, k):
  k=k%len(array)
  if len(array)==0:return []
  temp=array[-k:]
  for i in reversed(range(0,len(array)-k)):
    array[i+k]=array[i]
  for i in range(len(temp)):
    array[i]=temp[i]
  return array
print(rotate_array([1,2,3,4],2))