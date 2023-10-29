module.exports = {
  rules: {
    "header-max-length": [2, "always", 72],
    "type-case": [2, "always", "lower-case"],
    "type-empty": [2, "never"],
    "type-enum": [
      2,
      "always",
      [
        "games",
        "guides",
        "multi",
        "tools",
        "users",
        "build"
      ],
    ],
    "subject-empty": [2, "never"]
  },
};
