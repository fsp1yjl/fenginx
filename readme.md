> tag 0.01
仅仅实现简单的tcp server

> tag 0.02
对客户端请求处理进行封装，增加IConnection接口以及实现了该接口对Connection struct
增加demo文件夹，存放测试server/client文件。

> tag 0.03
在每个连接请求的connection进行进一步封装，将请求信息封装为一个request对象，然后抽象出一个router对象，用于对request进行实际对业务逻辑处理。
router对象可由用户在创建server的时候指定，并通过绑定到server，在每次创建connection到时候传递过去，最后在connection的每次请求处理时， 调用router的Handle方法去处理请求业务。
