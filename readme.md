## 这里是golang第一个项目组织说明 ##

<pre>
0.golang中的import name,实际是到GOPATH中去寻找name.a,使用时是这个name.a的源码中声明的package名字
1.把项目根目录加入GOPATH. (GOPATH第一个路径是默认第三方包安装目录)
2.项目目录下有 bin,src,pkg 三个目录
3.在项目目录下 go install abc_name时,系统会到GOPATH的src目录中寻找abc_name目录,然后编译其下的go文件
4.同一个目录中所有的go文件的package声明必须相同,所以main方法要单独放一个文件
</pre>