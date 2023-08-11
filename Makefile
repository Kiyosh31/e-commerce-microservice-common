tag:
	@git add .
	@git commit -m "Changed"
	@git push
	@git tag "v1.0.33"
	@git push --tag

tidy:
	go mod tidy