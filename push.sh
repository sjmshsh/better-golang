#!/bin/bash

# 获取命令行参数
commitValue="$1"

# 打印输入的值
git add .
git commit -m "$commitValue"
git push origin master