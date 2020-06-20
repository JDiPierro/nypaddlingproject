module.exports = {
  configureWebpack: {
    externals: {
      // global app config object
      config: JSON.stringify({
        apiUrl: process.env.API_URL || 'nypaddlingproject.herokuapp.com/api'
      })
    }
  },
  devServer: {
    disableHostCheck: true
  }
}
