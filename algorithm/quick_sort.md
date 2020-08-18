# 快速排序 quick sort

一分为二，设立基准，进行对比。采用分治法 divide and conquer

function(){
    var arr=["9238","45","56","34","76"]
    repeat(arr)
}

function repeat(arr []){
    if(arr.length ==1){
        return arr[0]
    }
    int middleIndex = arr.length/2
    var middleValue = arr[middleIndex]
    var leftarr,rightarr=[]
    for(var i=0;i<arr.length;i++){
        if(arr[i]>middleValue){
            leftarr.add(arr[i])
        }else{
            rightarr.add(arr[i])
        }
    }
    if(leftarr.length!=0){
        repeat(leftarr)
    }
    if(rightarr.length!=0){
        repeat(rightarr)
    }
}