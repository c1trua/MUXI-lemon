##  初始化配置
- 省略(Local)：本地配置，只对本地仓库有效
- --global：全局配置，所有仓库生效(常用)
- --system：系统配置，对所有用户生效

git config --global user.name "xxx"
git config --global user.email ""
git config --global credential.helper store //存储密码
git config --global --list

## 新建仓库
本地创建：git init 
在当前目录创建仓库

在远程服务器上克隆仓库：
git clone 仓库地址

## 工作区域和文件状态
工作区(Working Directory)：.git所在的目录
工作目录，资源管理器里能看见的文件夹就是工作区

暂存区(Staging Area/Index)：.git/index
**临时存储区域**，用于保存即将提交到Git仓库的修改内容

本地仓库(Local Repository)：.git/objects
即用git init创建的本地仓库，包含完整项目历史和元数据，是Git存储代码和版本信息的主要位置

Working Directory ->(git add) Staging Area/Index ->(git commit) Local Repository
![image1](C:\LEMON\NOTES\images\bc36947f-cf62-4711-8b07-bf7d5869c734.png)

## 添加和提交文件
**git status** 查看仓库的状态
红色：Untrack

**git add** 文件名 
添加到**暂存区**
绿色：被添加到了**暂存区**，等待被提交
use "git rm --cached <file\>..." to unstage
git add \*.txt	提交所有txt结尾的文件
git add .	把当前文件夹所有文件提交到暂存区

**git commit -m "提交的信息"**
提交**暂存区**中的文件

**git log**	查看提交记录
git log --oneline	查看简洁提交记录

## git reset回退版本
HEAD^ 上一版本
**git reset --soft 版本ID**
回退到某一版本，并且保留工作区和暂存区的所有修改内容

**git reset --hard ID**
回退到某一版本，并且丢弃工作区和暂存区的所有修改内容

**git reset --mixed ID**	(reset的默认参数)
回退到某一版本，并且保留工作区的修改内容，丢弃暂存区的修改内容

**git reflog**	查看操作历史记录
**git reset ID**	回退到某个操作前的版本

## 使用git diff查看差异
**git diff**
默认比较工作区和暂存区之间的差异内容

get diff HEAD
比较工作区和版本库之间的差异

get diff --cached
比较暂存区和版本库之间的差异

get diff ID1 ID2
(或使用HEAD表示当前分支的最新提交)
HEAD~ 或 HEAD^ 表示上一个版本
HEAD~2 表示HEAD的之前两个版本
比较两个版本之间的差异内容

get diff HEAD~ HEAD file1.txt
只查看此文件的差异内容

get diff branch_name branch_name2
比较**分支**之间的差异

## 使用git rm删除文件
**git ls-files**
查看**暂存区**中的内容

**git rm 文件名**
从工作区和暂存区同时删除
记得git commit -m "" 提交

## gitignore忽略文件
应该忽略哪些文件：
- 系统或者软件自动生成的文件
- 编译产生的中间文件和结果文件
- 运行时生成日志文件、缓存文件、临时文件
- 设计身份、密码、口令、秘钥等敏感信息文件

创建.gitignore文件：
使用命令行 echo > .gitignore
之后使用文本编辑器编辑内容

.gitignore文件的匹配规则：
#开头表示注释
\*作为通配符，通配任意个字符，如\*.txt表示所有含名字.txt的文件
?匹配单个字符
[]匹配列表中的任意一个字符
	可用短中线连接，如[1-9]表示任意一位数字，[a-z]表示任意一位小写字母
\*\*匹配任意的中间目录
!表示取反，忽略指定模式以外的文件或者目录

加入文件名以忽略某一文件

之后使用git add和git commit提交忽略操作
注意：被忽略的文件不能是已经被添加到版本库中的文件

## GITHUB
生成SSH Key ssh-keygen -t rsa -b 4096
- 私钥文件：id_rsa
- 公钥文件：id_rsa.pub

克隆仓库 git clone repo-address
推送更新内容 git push <remote\> <branch\>
拉取更新内容 git pull <remote\>

## 关联本地仓库和远程仓库
添加远程仓库：
- git remote add <远程仓库别名\> <远程仓库地址\>
- git push -u <远程仓库名\> <分支名\>

查看远程仓库：
git remote -v

拉去远程仓库内容：
git pull <远程仓库名\> <远程分支名\>:<本地分支名\> 	(相同可省略)

## VSCODE
![image2](C:\LEMON\NOTES\images\45ba3ca0-09ad-4eb1-be3f-d8520d32c903.png)

## 分支简介和基本操作
- weishishengzhenshigesb
- weishishengzhenshigedsb
- weishishengzhenshigeddDsb

查看分支列表：git branch
创建分支 git branch branch-name
切换分支：git switch branch-name
合并分支：git merge branch-name
删除分支：
- 已合并	git branch -d branch-name
- 未合并	git branch -D branch-name

## 解决合并冲突
git commit -am ""	一个命令完成暂存和提交
git status	查看冲突文件列表
git diff	查看冲突具体内容

解决方法
- 手工修改冲突文件
- 添加暂存区 git add file
- 提交修改 git commit -m ""

中止合并：git merge --abort

## 回退和rebase
alias <name\>=""	为指令创建别名
使用rebase时，找到当前分支和目标分支的共同祖先，再把**当前分支**上从共同祖先到最新提交记录的所有提交都**移动到目标分支**的最新提交记录后面


merge
- 优点：不会破坏原分支的提交历史，方便回溯和查看
- 缺点：会产生额外的提交节点，分支图比较复杂

rebase
- 优点：不会新增额外的提交记录，形成线性历史，比较直观和干净
- 会改变提交历史，改变了当前分支的branch out的节点，避免在共享分支使用

## 分支管理和工作流模型
