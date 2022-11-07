$(function(){
    app.init();
})
var app={
    init:function(){
        this.getCaptcha()
        this.captchaImgChage()
    },
    getCaptcha:function(){
        //得到图形验证码和 id
        $.get("/admin/captcha?t="+Math.random(),function(response){
            console.log(response)
            $("#captchaId").val(response.captchaId)
            $("#captchaImg").attr("src",response.captchaImage)
        })
    },
    captchaImgChage:function(){
        var that=this;
        $("#captchaImg").click(function(){
            that.getCaptcha()
        })
    }
}