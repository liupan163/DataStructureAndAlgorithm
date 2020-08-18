# 基数排序 radix sort
基数排序是按照低位先排序，然后收集，再依次按高位排序，再收集，直到最高位。

function(arr){
    //先找到最大位数的树，来确定最大位数
    for(var i=0;i<arr.length;i++){
        arrEnd[arr[i]%10].add(arr[i])
    }
    for(var i=0;i<arrEnd.length;i++){
        for(var j=0;j<arrEnd[i].length;j++){
            result.add(arrEnd[i][j])
        }
    }
    //
    var divideTen = 1
    for(maxNumber/10/divideTen>0){
        for(var i=0;i<arr.length;i++){
            arrEnd[arr[i]/10/divideTen].add(arr[i])
        }
        for(var i=0;i<arrEnd.length;i++){
            for(var j=0;j<arrEnd[i].length;j++){
                result.add(arrEnd[i][j])
            }
        }
        divideTen = divideTen * 10
    }
}