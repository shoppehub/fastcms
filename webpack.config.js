const CopyPlugin = require("copy-webpack-plugin");
const HandlebarsPlugin = require("handlebars-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const CssMinimizerPlugin = require("css-minimizer-webpack-plugin");
const RemoveEmptyScriptsPlugin = require("webpack-remove-empty-scripts");
const TerserPlugin = require("terser-webpack-plugin");
const autoprefixer = require("autoprefixer");
const path = require("path");

const paths = {
  src: {
    favicon: "./web/assets/favicon",
    fonts: "./web/assets/fonts",
    img: "./web/assets/img",
    js: "./web/assets/js",
    scss: "./web/assets/scss",
  },
  dist: {
    css: "./assets/css",
    favicon: "./assets/favicon",
    fonts: "./assets/fonts",
    img: "./assets/img",
    js: "./assets/js",
  },
};

module.exports = {
  devServer: {
    port: 4001,
    headers: {
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "*",
      "Access-Control-Allow-Headers": "*",
    },
  },
  devtool: "source-map",
  entry: {
    libs: [paths.src.scss + "/libs.scss"],
    theme: [...[paths.src.js + "/theme.js", paths.src.scss + "/theme.scss"]],
    datasource: ["./web/pages/datasources/index.js"],
    system_rule: ["./web/pages/system/rule/index.js"],
  },
  mode: "development",
  module: {
    rules: [
      {
        test: /\.(sass|scss)$/,
        include: path.resolve(__dirname, paths.src.scss.slice(2)),
        use: [
          {
            loader: MiniCssExtractPlugin.loader,
          },
          {
            loader: "css-loader",
            options: {
              url: false,
            },
          },
          {
            loader: "postcss-loader",
            options: {
              postcssOptions: {
                plugins: [["autoprefixer"]],
              },
            },
          },
          {
            loader: "sass-loader",
          },
        ],
      },
    ],
  },
  optimization: {
    splitChunks: {
      cacheGroups: {
        vendor: {
          test: /[\\/](node_modules)[\\/].+\.js$/,
          name: "vendor",
          chunks: "all",
        },
      },
    },
    minimizer: [
      new CssMinimizerPlugin(),
      new TerserPlugin({
        extractComments: false,
        terserOptions: {
          output: {
            comments: false,
          },
        },
      }),
    ],
  },
  output: {
    filename: paths.dist.js + "/[name].bundle.js",
  },
  plugins: [
    new CopyPlugin({
      patterns: [
        {
          from: paths.src.favicon,
          to: paths.dist.favicon,
        },
        {
          from: paths.src.fonts,
          to: paths.dist.fonts,
        },
        {
          from: paths.src.img,
          to: paths.dist.img,
        },
      ],
    }),

    new RemoveEmptyScriptsPlugin(),
    new MiniCssExtractPlugin({
      filename: paths.dist.css + "/[name].bundle.css",
    }),
  ],
  target: "web",
};
