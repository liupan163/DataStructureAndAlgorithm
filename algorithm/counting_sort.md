# 计数排序 counting sort
-   key-value形式
-   key是值，value是这个值出现来多少次

funtcion countingSort(arr,maxValue){
    var countResult;
    var result;
    for(var i=0;i<arr;i++){
        if(!countResult.has(arr[i])){
            countResult[arr[i]] = 1
        }else{
            countResult[arr[i]] = countResult[arr[i]] + 1
        }
    }
    var keys = countResult.keys
    for(var index=0;index<keys.length;index++){
        for(var i=0;i<countResult[keys[index]];i++){
            result.add(keys[index])
        }
    }
    return result
}