const proxy = [
  {
    context: '/',
    target: 'http://localhost:8080',
    secure: false,
    pathRewrite: {'^/' : ''}
  }
];
module.exports = proxy;
