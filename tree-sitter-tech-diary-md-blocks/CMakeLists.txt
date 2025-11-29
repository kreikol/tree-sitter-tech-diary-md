set(PROJECT_DESCRIPTION "Tech Diary Markdown block grammar for tree-sitter")
set(TS_REQUIRES tree-sitter-tech-diary-md-inline)

add_parser(tech_diary_md_blocks)

add_custom_target(test "${TREE_SITTER_CLI}" test
                  WORKING_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}"
                  DEPENDS "${CMAKE_CURRENT_SOURCE_DIR}/src/parser.c"
                  COMMENT "tree-sitter test")
