function monotonicArray(array){
    if(array<=1) return true
    let left=array[0]
    let right=array[array.length-1]
    
    if(left>right){
      for(let i=0; i<array.length-1;i++){
        if(array[i]<array[i+1]) return false
      }
    }
    else if(left===right){
      for (let i=0; i<array.length-1; i++){
        if(array[i] !==array[i+1]) return false
      }
    }
    else{
      for (let i=0; i<array.length-1;i++){
        if(array[i]>array[i+1]) return false
      }
    }
    return true
  }
  
  console.log(monotonicArray([4,3,2,5]))