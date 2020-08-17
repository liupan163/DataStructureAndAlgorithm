# 插入排序
从前往后依次增加结果集，新增的元素跟从后往前，依次对比，增加。
var arr=["2","5","4","6"]
var result
for(var i=0;i<arr.length;i++){
    if(i=0){
       result[i]=arr[i]
       continue
    }
    for(var j=result.length-1;j>0;j--){
        if(result[j]<arr[i]){
            result[i] = arr[i]
        }else{
            result[i+1]=result[i]
        }
    }
}