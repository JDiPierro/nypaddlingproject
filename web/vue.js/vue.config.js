module.exports = {
  configureWebpack: {
    externals: {
      // global app config object
      config: JSON.stringify({
        apiUrl: process.env.API_URL || 'http://localhost:5000'
      })
    }
  }
}
