### gconv 

### 背景

后台开发中，我们经常会遇到需要将 `结构体` 转化为 `map[string]interface{}`; 对比了github上各个转化库，发现`gconv`这个模块能够实现自定义转化字段名，忽略某些字段或者忽略空值字段等。

代码原始地址： https://github.com/gogf/gf/tree/master/util/gconv


# 基本示例

```
package main

import (
    "github.com/k0kubun/pp"
    "github.com/zhwei820/gconv"
)

func main() {
    type User struct {
        Uid  int    `c:"uid"`
        Name string `c:"name"`
    }
    // 对象
    pp.Print(gconv.Map(User{
        Uid:  1,
        Name: "john",
    }))
    // 对象指针
    pp.Print(gconv.Map(&User{
        Uid:  1,
        Name: "john",
    }))

    // 任意map类型
    pp.Print(gconv.Map(map[int]int{
        100: 10000,
    }))
}
```

执行后，终端输出：

```
{
    "name": "john",
    "uid": 1
}

{
    "name": "john",
    "uid": 1
}

{
    "100": 10000
}
```


