const isProd = process.env.NODE_ENV === 'production'

module.exports = {
  basePath: isProd ? '/ctf-website' : '',
  assetPrefix: isProd ? '/ctf-website' : '',
}
