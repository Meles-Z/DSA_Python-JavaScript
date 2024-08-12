const reverse=function(nums, start, end){
    while(start<end){
      [nums[start], nums[end]]=[nums[end], nums[start]]
    start++
    end--
    }
  }
  const rotate=function(nums, k){
    // first reverse all the part of code
    k=k%nums.length;
    reverse(nums,0, nums.length-1)
    //second reverse
    reverse(nums, 0, k-1)
    // third reverse
    
    reverse(nums, k, nums.length-1)
    return nums
  }
  
  console.log(rotate([1,3,4],1))