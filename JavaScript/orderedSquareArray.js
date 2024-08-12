
function sortedSquarredArray(array){
    const newArray=new Array(array.length).fill(0)
    let left=0;
    let right=array.length-1
    
    for (let i=array.length-1; i>=0; i--){
        const leftSquared=Math.pow(array[left], 2)
        const rightSquared=Math.pow(array[right], 2)
        if(leftSquared>rightSquared){
            newArray[i]=leftSquared
            left++
        }else{
            newArray[i]=rightSquared
            right--
        }
    }
    return newArray
}
console.log(sortedSquarredArray([1,2,4]))