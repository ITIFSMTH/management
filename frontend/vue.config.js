const { defineConfig } = require('@vue/cli-service')

// Pares env variables
if (process.env.NODE_ENV == "development") {
  require('dotenv').config({
    path: '../env/.env.frontend'
  })
}

module.exports = defineConfig({
  transpileDependencies: true,
  publicPath: "/",
})