const { createProxyMiddleware } = require("http-proxy-middleware");
// 设置代理配置
module.exports = function(app) {
    app.use(
        createProxyMiddleware("/api", {
            target: "http://120.78.134.104:8080/", //服务器接口地址 // http://120.78.134.104:8080/api/upload/image
            changeOrigin: true,
            secure: false,
        })
    );
};