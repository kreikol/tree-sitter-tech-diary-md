{
  "targets": [
    {
      "target_name": "tree_sitter_markdown_binding",
      "dependencies": [
        "<!(node -p \"require('node-addon-api').targets\"):node_addon_api_except",
      ],
      "include_dirs": [
        "tree-sitter-tech-diary-md-blocks/src",
      ],
      "sources": [
        "bindings/node/binding.cc",
        "tree-sitter-tech-diary-md-blocks/src/parser.c",
        "tree-sitter-tech-diary-md-blocks/src/scanner.c",
        "tree-sitter-tech-diary-md-inline/src/parser.c",
        "tree-sitter-tech-diary-md-inline/src/scanner.c",
      ],
      "cflags_c": [
        "-std=c11",
      ],
    }
  ]
}
