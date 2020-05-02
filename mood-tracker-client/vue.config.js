module.exports = {
  devServer: {
    proxy: {
      '^/grpc': {
        target: 'http://localhost:8090/grpc',
        ws: true,
        changeOrigin: true
      }
    }
  }
}
