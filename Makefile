tag:
	@git add .
	@git commit -m "Changed"
	@git push
	@git tag "v1.0.15"
	@git push --tag