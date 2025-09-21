## 📄 `test.md`

# 一级标题 H1
## 二级标题 H2
### 三级标题 H3
#### 四级标题 H4

---

## 段落 / 换行

这是一个普通段落。  
这是一个段落里 

这是另一个段落。

---

## 强调

*斜体*  
_斜体_  
**粗体**  
__粗体__  
~~删除线~~

---

## 列表

无序列表：
- 项目 1
- 项目 2
    - 子项目 2.1
    - 子项目 2.2

有序列表：
1. 第一项
2. 第二项
    1. 子项 a
    2. 子项 b

任务列表：
- [ ] 待办事项
- [x] 已完成事项

---

## 链接 & 图片

这是一个 [超链接](https://www.example.com)。

这是一个图片：

![photo_2025-09-19_17-24-50.jpg](https://s2.loli.net/2025/09/19/Kf4h1qjLrGxJHYB.jpg "图片提示文字")

---

## 引用

> 这是一个引用块。  
> 可以换行，支持 **Markdown** 语法。

---

## 代码

行内代码：`print("Hello World")`

代码块（无语言）：
```

这里是代码块
 缩进保持原样

```
代码块（js）：
```js
let a = 2 + 2;

switch (a) {
  case 3:
    alert( 'Too small' );
    break;
  case 4:
    alert( 'Exactly!' );
    break;
  case 5:
    alert( 'Too big' );
    break;
  default:
    alert( "I don't know such values" );
}
```

------

## 表格

| 姓名 | 年龄 | 城市 |
| ---- | ---- | ---- |
| 张三 | 25   | 北京 |
| 李四 | 30   | 上海 |
| 王五 | 28   | 广州 |

------

## 水平分割线

------


## 脚注

这里是脚注[1](https://chatgpt.com/c/68ccac7f-5560-8324-9217-d268a4a66b5d#user-content-fn-1) 的例子。

------

## 数学公式

行内公式：$E = mc^2$

块级公式：

$$
\mathbb{E}(X) = \int x d F(x) = \left\{ \begin{aligned} \sum_x x f(x) \; & \text{ if } X \text{ is discrete}
\\ \int x f(x) dx \; & \text{ if } X \text{ is continuous }
\end{aligned} \right.
$$

------

## Footnotes

1. 脚注内容。 [↩](https://chatgpt.com/c/68ccac7f-5560-8324-9217-d268a4a66b5d#user-content-fnref-1)