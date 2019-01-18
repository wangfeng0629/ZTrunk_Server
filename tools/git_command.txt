git 基本常用命令	
git  --version	查看当前使用git版本
git  config --global user.name  "***"	设置 用户名
git  config --global user.email  "***"	设置 邮件地址
git  config --global  color.ui  true	在git输出中开启颜色指示
git  config --global alias.st   status	给 git status 命令添加别名
	
git add -u	可以将（被版本库追踪的）本地文件的变更（修改、删除）全部记录到暂存区中。
git add 	把工作区文件添加到暂存区
git commit -a	对本地所有变更的文件执行提交操作，包括对本地修改的文件和删除的文件，但不包括未被版本库跟踪的文件。这个“偷懒”的提交命令，会丢掉Git暂存区带给用户的最大好处：对提交内容进行控制的能力。
git commit -m “***”	把暂存区文件提交到版本库
	
git log --pretty=oneline	精简输出显示日志
git log	显示日志
	
git reflog  -n	查看git reflog的输出和直接查看日志文件最大的不同在于显示顺序的不同，即最新改变放在了最前面显示，而且只显示每次改变的最终的SHA1哈希值。（显示n行）
	
git status -s	查看文件精简状态
	
git diff	默认是查看工作区和暂存区的区别
git  diff HEAD	查看工作区和版本库的区别
git diff --cached	查看暂存区和版本库的区别
	
git branch	查看当前分支
	
	
git reset  = git reset HEAD	仅用HEAD指向的目录树重置暂存区，工作区不会受到影响，相当于将之前用git add命令更新到暂存区的内容撤出暂存区。引用也未改变，因为引用重置到HEAD相当于没有重置。
git reset --hard HEAD^	彻底撤销最近的提交。引用回退到前一次，而且工作区和暂存区都会回退到上一次提交的状态。自上一次以来的提交全部丢失。
git reset --soft HEAD^	工作区和暂存区不改变，但是引用向前回退一次。当对最新提交的提交说明或提交的更改不满意时，撤销最新的提交以便重新提交。
	
	
git checkout	汇总显示工作区、暂存区与HEAD的差异。
git merge	合并
git checkout--filename	用暂存区中filename文件来覆盖工作区中的filename文件。相当于取消自上次执行git add filename以来（如果执行过）的本地修改。
git checkout .（有点）	git checkout命令后的参数为一个点（“.”）。这条命令最危险！会取消所有本地的修改（相对于暂存区）。相当于用暂存区的所有文件直接覆盖本地文件，不给用户任何确认的机会！
git checkout branch--filename	维持HEAD的指向不变。用branch所指向的提交中的filename替换暂存区和工作区中相应的文件。注意会将暂存区和工作区中的filename文件直接覆盖。
	
	
git stash 	保存当前进度，会分别对暂存区和工作区的状态进行保存。
git stash list	查看保存的进度用命令
git stash pop	从最近保存的进度进行恢复。
git stash apply[--index][<stash>]	除了不删除恢复的进度之外，其余和git stash pop命令一样。
git stash clear	删除所有存储的进度。
	
git tag -m “****描述” ###	里程碑 打标签 ###名字
git   describe	显示最近的里程碑
	
git rm 	删除文件  删除动作加入了暂存区。这时执行提交动作，就从真正意义上执行了文件删除。
	
	
git pull = git fetch + git merge	拉取云端最新版本库到本地
git push	推送本地版本库到云端
	
冲突	
<<<<<<<     =======	当前分支所更改的内容
（七个等于号）=======     >>>>>>>（七个大于号）所合并的版本更改的内容。	
