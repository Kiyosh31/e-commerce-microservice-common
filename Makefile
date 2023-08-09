tag:
	@git add .
	@git commit -m "Changed"
	@git push
	@git tag "v1.0.11"
	@git push --tag