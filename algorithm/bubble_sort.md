# 冒泡排序bubble sort

重复的走访过要排序的数列，一次比较两个，判断结果执行交换，直到队列完成。
重复上述步骤，
--->会让最大的元素落在最后面，最小的元素，浮起来。

实例，从小到大排序
var arr[]=["6","5","4","4"]
var result;
for(int i=0;i<arr.length-1;i++){
    for(int j=i;j<arr.length-1-1;j++){
        if(arr[j]>arr[j+1]){
            var temp = arr[j]
            arr[j]=arr[j+1]
            arr[j+1]=temp
        }
    }
    result.add(arr[i])
}
