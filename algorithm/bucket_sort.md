# 桶排序
设定多个范围，
将数据遍历，放入对应范围的桶内
对不是空桶的桶进行排序
将各个非空桶，拼接起来

function(arr){
    var divide = 5
    var offset = arr/5
    var result
    var tempArray = [][]
    for(var i=0;i<arr.length;i++){
        var startIndex = arr[i]/offset
        tempArray[arr[i]/offset].add(arr[i])
    }
    for(var i=0;i<tempArray.length;i++){
        //对每个桶进行排序 todo
        if(tempArray[i].length>0){
            for(var j=0;j<tempArray[i].length;j++){
                result.add(tempArray[i][j])
            }
        }
    }
    return result
}
