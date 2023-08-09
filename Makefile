tag:
	@git add .
	@git commit -m "Changed"
	@git push
	@git tag "v1.0.12"
	@git push --tag