# 选择排序 selection sort
直观的排序方法，选择序列中，最大或者最小的元素。

func selectionSort(){
    var result;
    var minValue;
    for(var i=0;i<arr.length-1;i++){
        for(var j=i+1;j<arr.length-1-1;j++){
            if(arr[i]>arr[j]){
                minValue = arr[j]
                arr[j]=arr[i]
            }
        }
        result.add(minValue)
    }
}