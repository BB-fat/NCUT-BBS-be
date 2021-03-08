# NCUT-BBS-be

NCUT论坛后端。

## 配置自动代码格式化

通过git hook调用dartfmt在提交时自动格式化代码。

1. 在`.git/hooks`目录下新建`pre-commit`：

```
#!/bin/bash
for file in `git diff-index --cached --name-only HEAD` ; do

  extension="${file##*.}"

  if [ $extension = "go" ] && [ -f ${file} ] ; then

    echo "format ${file}"

    go fmt ${file}
    
    git add ${file}
  fi

done

```

1. 使用命令`chmod a+x pre-commit`给`pre-commit`脚本添加可执行权限。

