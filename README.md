# message-board
关于表格设计：
user表：id主键 username用户名 password密码 protectionQ密保问题 protectionA密保答案

message表：id主键用于标识message 
           pid用于标识母message，在comment时输入母评论id就会在该message的pid上赋值母message的id，在查询母message评论时，就会query pid为母message的id的message（commnent）
           touesr给谁留言 fromuser留言者 messagecontent留言内容 truename用于标识匿名评论者的username likes点赞数 messagetime发布留言的时间


一些小功能：入参限制：密码必须大于三位
                     评论必须大于五位
            /message/nonamemsg&&/message/nonamecom 匿名留言和匿名评论：留言时fromuser会被赋值为noname，但truename仍会记录留言者信息，用于修改和删除评论
            /message/like 点赞：输入评论id以增加点赞数
            /user/clock 显示时间
