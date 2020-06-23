module.exports = {
  configureWebpack: {
    externals: {
      // global app config object
      config: JSON.stringify({
        apiUrl: process.env.API_URL || '/api'
      })
    }
  },
  devServer: {
    disableHostCheck: true,
    proxy: 'https://nypaddledev.com'
  }
}
