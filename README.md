# captcha-golang-demo
易盾验证码golang演示

# demo运行步骤
* 修改web/index.html中的YOUR_CAPTCHA_ID
```javascript
initNECaptcha({
  captchaId: 'YOUR_CAPTCHA_ID',
  element: '#captcha_div',
  mode: 'float', // 如果要用触发式，这里改为float即可
  width: '320px',
  onVerify: function(err, ret){
    if(!err){
        $.ajax({
            type: "POST",
            url: "/verify" ,
            data: ret,
            success: function (result) {
               alert(result)
            },
            error : function() {
                alert("异常！");
            }
        });
    }
  }
}, function (instance) {
  // 初始化成功后得到验证实例instance，可以调用实例的方法
}, function (err) {
  // 初始化失败后触发该函数，err对象描述当前错误信息
})
```

* 修改 main.go
```go
const(
    captchaId = "YOUR_CAPTCHA_ID"
    secretId = "YOUR_SECRET_ID"
    secretKey = "YOUR_SECRET_KEY"
)
```

* ``go run main.go``
* 浏览器访问 http://localhost:8080/ 查看演示