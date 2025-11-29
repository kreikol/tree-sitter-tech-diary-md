all install uninstall clean:
	$(MAKE) -C tree-sitter-tech-diary-md-blocks $@
	$(MAKE) -C tree-sitter-tech-diary-md-inline $@

.PHONY: all install uninstall clean
