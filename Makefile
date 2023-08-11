tag:
	@git add .
	@git commit -m "Changed"
	@git push
	@git tag "v1.0.30"
	@git push --tag

tidy:
	go mod tidy