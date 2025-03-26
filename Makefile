
#Locales struct generation
.PHONY: locales
locales:
	go run ./tools/locales/generator.go
