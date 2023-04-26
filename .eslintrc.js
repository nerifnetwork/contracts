module.exports = {
  env: {
    browser: false,
    es2021: true,
    mocha: true,
    node: true,
  },
  plugins: ["@typescript-eslint"],
  extends: [
    "eslint:recommended",
    "plugin:prettier/recommended",
    "plugin:node/recommended",
  ],
  parser: "@typescript-eslint/parser",
  parserOptions: {
    ecmaVersion: 12,
    project: ["./tsconfig.json"],
  },
  rules: {
    "max-len": [
      "error",
      { code: 120 },
    ],
    "object-curly-spacing": [
      "error",
      "always",
    ],
    "node/no-unsupported-features/es-syntax": [
      "error",
      { ignores: ["modules"] },
    ],
    "node/no-missing-import": [
      "error", {
        "tryExtensions": [".js", ".json", ".node", ".ts", ".d.ts"]
      }
    ]
  },
};
