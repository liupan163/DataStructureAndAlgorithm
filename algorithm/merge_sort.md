# 归并排序 merge sort

在归并操作基础上，采用分治法 divide and conquer

function(){
    var arr=["9238","45","56","34","76"]
    repeat(arr)
}

function repeat(arr []){
    if(arr.length==0){
        return
    }
    if(arr.length==1){
        return arr[0]
    }
    middleIndex = arr.length/2
    leftarr = arr[0:middleIndex]
    rightarr = arr[middleIndex:]
    return mergeAotmic(repeat(leftarr),repeat(rightarr))
}

function mergeAotmic(leftarr [],rightarr []){
    var resultarr
    if(leftarr.length>0 && rightarr.length>0){
        if(leftarr[0]<=rightarr[0]){
            resultarr.push(leftarr[0])
            leftarr.remove(0)
        }else{
            resultarr.push(rightarr[0])
            rightarr.remove(0)
        }
    }
    while(leftarr.length>0){
        resultarr.push(leftarr[0])
        leftarr.remove(0)
    }
    while(rightarr.length>0){
        resultarr.push(rightarr[0])
        rightarr.remove(0)
    }
    return resultarr
}