def orderdSquaredArray(array):
    #[1,3,5,6]
    new_array=[0]*len(array)
    leftSide=0
    righSide=len(array)-1
    for i in reversed(range(len(array))):
        leftSquare=array[leftSide]**2
        rightSquare=array[righSide]**2
        if(leftSquare>rightSquare):
            new_array[i]=leftSquare
            leftSide+=1
        else:
            new_array[i]=rightSquare
            righSide-=1
    return new_array
print(orderdSquaredArray([1,2,3,4]))
print(orderdSquaredArray([-9,1,-2,3,4]))
