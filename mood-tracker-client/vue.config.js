module.exports = {
  devServer: {
    proxy: {
      '^/grpc': {
        target: 'http://localhost:8090',
        ws: true,
        changeOrigin: true
      }
    }
  }
}
