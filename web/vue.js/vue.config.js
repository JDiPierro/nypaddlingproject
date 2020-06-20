module.exports = {
  configureWebpack: {
    externals: {
      // global app config object
      config: JSON.stringify({
        apiUrl: process.env.API_URL || 'localhost:8080/api'
      })
    }
  },
  devServer: {
    disableHostCheck: true
  }
}
