# 希尔排序
希尔排序是简单插入排序的改进版，区别是：它会优先比较距离较远的元素。
希尔排序的核心在于间隔序列的设定。既可以提前设定好间隔序列，也可以动态的定义间隔序列。

func shellSort(){
    var arr=["23","43","54","6546","46"]
    for(var gap=Math.floor(arr.length/2); gap>0; gap=Math.floor(gap/2)){
        for(var i=gap; i<arr.length; i++){
            var j=i
            var current = arr[i]
            while(j-gap>=0 && current<arr[j-gap]){
                arr[j] = arr[j-gap]
                j = j-gap
            }
            arr[j] = current
        }
    }
}