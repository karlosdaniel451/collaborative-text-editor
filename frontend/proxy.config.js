const proxy = [
  {
    context: '/api',
    "changeOrigin": true,
    target: 'http://localhost:8080',
    pathRewrite: {'^/api' : ''}
  }
]

module.exports = proxy;
