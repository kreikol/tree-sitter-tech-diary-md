const root = require("path").join(__dirname, "..", "..");

module.exports = require("node-gyp-build")(root);

try {
  module.exports.nodeTypeInfo = require("../../tree-sitter-tech-diary-md-blocks/src/node-types.json");
  module.exports.inline.nodeTypeInfo = require("../../tree-sitter-tech-diary-md-inline/src/node-types.json");
} catch (_) { }
